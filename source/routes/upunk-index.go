package routes

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"yam-api/source/config"
	"yam-api/source/utils"
	"yam-api/source/utils/log"
	"yam-api/source/utils/mongodb"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/montanaflynn/stats"
	"github.com/robfig/cron"
	"github.com/shopspring/decimal"

	store "yam-api/source/contracts"
)

// --------- Variables ---------

var calculatePunkIndexCron *cron.Cron

// --------- Types ---------

type PunkEvent struct {
	punkIndex   string
	value       *big.Int
	fromAddress common.Address
	toAddress   common.Address
}

type PunkBidEntered struct {
	punkIndex   string
	value       *big.Int
	fromAddress common.Address
}

// --------- Helper Functions ---------

func GetBlockNumberByTimestamp(geth *ethclient.Client, ts uint64, direction string) uint64 {
	nBlocks := big.NewInt(10_000)

	header, err := geth.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error(err)
	}

	blockNumber := header.Number
	currentBlock, err := geth.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Error(err)
	}

	currentBlockNumBig := currentBlock.Number()
	currentBlockNum := currentBlock.Number().Uint64()
	cm100Block, err := geth.BlockByNumber(context.Background(), big.NewInt(0).Sub(currentBlockNumBig, nBlocks))
	if err != nil {
		log.Error(err)
	}

	avgBlockTime := (currentBlock.Time() - cm100Block.Time()) / nBlocks.Uint64()

	if ts > currentBlock.Time() {
		return currentBlockNum
	}

	tsDelta := (currentBlock.Time() - ts)
	blockNumOlder := math.Min(
		float64(currentBlockNum)-math.Floor(float64(1.2)*(float64(tsDelta)/float64(avgBlockTime))),
		float64(currentBlockNum))

	blockOlder, err := geth.BlockByNumber(context.Background(), big.NewInt(int64(blockNumOlder)))
	if err != nil {
		log.Error(err)
	}

	blockNumNewer := math.Min(
		float64(currentBlockNum)-math.Floor(float64(0.8)*(float64(tsDelta)/float64(avgBlockTime))),
		float64(currentBlockNum))

	blockNewer, err := geth.BlockByNumber(context.Background(), big.NewInt(int64(blockNumNewer)))
	if err != nil {
		log.Error(err)
	}

	if blockOlder.Time() > ts && ts > blockNewer.Time() {
		panic("Need to change the bounds.")
	}

	var blockNumProposed float64

	for true {
		blockNumProposed = math.Floor((blockNumOlder + blockNumNewer) / 2)
		blockNumProposedM1 := blockNumProposed - 1

		proposed, err := geth.BlockByNumber(context.Background(), big.NewInt(int64(blockNumProposed)))
		if err != nil {
			log.Error(err)
		}

		proposedM1, err := geth.BlockByNumber(context.Background(), big.NewInt(int64(blockNumProposedM1)))
		if err != nil {
			log.Error(err)
		}

		if proposedM1.Time() < ts && ts <= proposed.Time() {
			break
		}

		if proposed.Time() > ts {
			blockNumNewer = blockNumProposed
		} else if proposed.Time() < ts {
			blockNumOlder = blockNumProposed
		}
	}

	return uint64(blockNumProposed)
}

func Filter(geth *ethclient.Client, contractAddress common.Address, blockStart uint64, blockEnd uint64) [][]types.Log {
	var logs [][]types.Log
	nblocks := uint64(math.Floor(10_000 / 0.25))
	var blockStarts []uint64

	for i := blockStart; i < blockEnd; i = i + nblocks {
		blockStarts = append(blockStarts, i)
	}

	log.Info("Filter")

	for _, bs := range blockStarts {
		be := bs + nblocks - 1

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(bs)),
			ToBlock:   big.NewInt(int64(be)),
			Addresses: []common.Address{
				contractAddress,
			},
		}

		_logs, err := geth.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error(err)
		}

		logs = append(logs, _logs)
	}

	return logs
}

func FilterWithTopic(geth *ethclient.Client, contractAddress common.Address, blockStart uint64, blockEnd uint64, topic common.Hash) [][]types.Log {
	var logs [][]types.Log
	nblocks := uint64(math.Floor(10_000 / 0.25))
	var blockStarts []uint64

	for i := blockStart; i < blockEnd; i = i + nblocks {
		blockStarts = append(blockStarts, i)
	}

	log.Info("FilterWithTopic")

	for _, bs := range blockStarts {
		be := bs + nblocks - 1

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(bs)),
			ToBlock:   big.NewInt(int64(be)),
			Addresses: []common.Address{
				contractAddress,
			},
			Topics: [][]common.Hash{
				{topic},
			},
		}

		_logs, err := geth.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error(err)
		}

		logs = append(logs, _logs)
	}

	return logs
}

/// @dev Wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// --------- uPunk Index Calculation ---------

func CalculatePunkIndex(geth *ethclient.Client) map[string]interface{} {
	PUNK_FIRST_BLOCK := 3914495
	// @dev CryptoPunksMarket contract address
	contractAddress := common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB")
	eval_ts := int(time.Now().UTC().Unix())
	window := int(30 * 24 * 60 * 60)

	blockStart := GetBlockNumberByTimestamp(geth, uint64(eval_ts-window), "after")
	blockEnd := GetBlockNumberByTimestamp(geth, uint64(eval_ts), "after")

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Error(err)
	}

	// @dev Event Signatures
	logPunkBoughtSig := []byte("PunkBought(uint256,uint256,address,address)")
	logPunkBoughtSigHash := crypto.Keccak256Hash(logPunkBoughtSig)
	logPunkBidEnteredSig := []byte("PunkBidEntered(uint256,uint256,address)")
	logPunkBidEnteredSigHash := crypto.Keccak256Hash(logPunkBidEnteredSig)

	var boughtLogs [][]types.Log = Filter(geth, contractAddress, blockStart, blockEnd)
	var bidLogs [][]types.Log = FilterWithTopic(geth, contractAddress, uint64(PUNK_FIRST_BLOCK), blockEnd, logPunkBidEnteredSigHash)
	var punkSales []map[string]string

	for i, _ := range boughtLogs {
		for _, vLog := range boughtLogs[i] {
			switch vLog.Topics[0].Hex() {
			case logPunkBoughtSigHash.Hex():

				_punkIndex := vLog.Topics[1].Hex()
				_blockNumber := vLog.BlockNumber

				var info = make(map[string]string)
				info["punkIndex"] = _punkIndex
				info["blockNumber"] = strconv.FormatUint(_blockNumber, 10)

				/// @dev Identify transaction
				transaction, _, err := geth.TransactionByHash(context.Background(), vLog.TxHash)
				if err != nil {
					log.Error(err)
				}

				/// @dev Recover Method from tx input data
				method, err := contractAbi.MethodById(transaction.Data())
				if err != nil {
					log.Error(err)
					continue
				}

				log.Info(method)

				buyMethod := contractAbi.Methods["buyPunk"]
				bidMethod := contractAbi.Methods["acceptBidForPunk"]

				if reflect.DeepEqual(*method, buyMethod) {
					punkBoughtEventData, err := contractAbi.Unpack("PunkBought", vLog.Data)
					if err != nil {
						log.Error(err)
					}

					var price interface{} = punkBoughtEventData[0]
					priceStr := fmt.Sprintf("%v", price)
					info["price"] = priceStr

					punkSales = append(punkSales, info)
				} else if reflect.DeepEqual(*method, bidMethod) {
					for i, _ := range bidLogs {
						if len(bidLogs[i]) == 0 {
							break
						}

						vLog := bidLogs[i][0]

						if vLog.BlockNumber >= _blockNumber {
							break
						}

						var punkBidEnteredEvent PunkEvent
						punkBidEnteredEvent.punkIndex = vLog.Topics[1].Hex()

						if punkBidEnteredEvent.punkIndex == _punkIndex {
							punkBidEnteredEventData, err := contractAbi.Unpack("PunkBidEntered", vLog.Data)
							if err != nil {
								log.Error(err)
							}

							var price interface{} = punkBidEnteredEventData[0]
							priceStr := fmt.Sprintf("%v", price)

							info["price"] = priceStr

							punkSales = append(punkSales, info)
						}
					}
				} else {
					log.Info("Problematic transcation Hash: ", vLog.TxHash)
				}
			}
		}
	}

	var prices []float64

	for i := 0; i < len(punkSales); i = i + 1 {
		priceFloat64, _ := strconv.ParseFloat(punkSales[i]["price"], 64)
		prices = append(prices, priceFloat64)
	}

	median, _ := stats.Median(prices)
	fmt.Println(median / 1000)
	upunkIndexWei := median / 1000
	wei := new(big.Int)
	str := strconv.FormatFloat(upunkIndexWei, 'f', 5, 64)
	wei.SetString(str, 10)
	eth := ToDecimal(wei, 18)
	upunkIndex := eth

	values := map[string]interface{}{"price": upunkIndex.String(), "timestamp": strconv.Itoa(eval_ts)}

	return values
}

// --------- uPunk Endpoints ---------

func GetLatestPunkIndex(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		/// @dev Set up cron
		if calculatePunkIndexCron == nil {
			calculatePunkIndexCron := cron.New()
			calculatePunkIndexCron.AddFunc("@every 5m", func() {
				values := CalculatePunkIndex(geth)
				mongodb.InsertPunkIndex(values)
			})
			calculatePunkIndexCron.Start()
		}

		/// @dev Retrieve values from db
		values := mongodb.GetLatestPunkIndex()
		if values == nil {
			values = map[string]interface{}{
				"price":     "No value stored.",
				"timestamp": "No value stored.",
			}
		}

		utils.ResJSON(http.StatusCreated, w,
			map[string]interface{}{
				"price":     values["price"],
				"timestamp": values["timestamp"],
			},
		)

		if values == nil {
			values = CalculatePunkIndex(geth)
			mongodb.InsertPunkIndex(values)
		}
	})
}

func GetPunkIndexHistory(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		/// @dev Set up cron
		if calculatePunkIndexCron == nil {
			calculatePunkIndexCron := cron.New()
			calculatePunkIndexCron.AddFunc("@every 5m", func() {
				values := CalculatePunkIndex(geth)
				mongodb.InsertPunkIndex(values)
			})
			calculatePunkIndexCron.Start()
		}

		/// @dev Retrieve values from db
		history := mongodb.GetPunkIndexHistoryDaily()
		if history == nil {
			history = []interface{}{"No values stored."}
		}

		utils.ResJSON(http.StatusCreated, w,
			history,
		)

		if history == nil {
			values := CalculatePunkIndex(geth)
			mongodb.InsertPunkIndex(values)
		}
	})
}
