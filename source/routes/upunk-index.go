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
	"yam-api/source/utils/etherscan/helper"
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

func Filter(geth *ethclient.Client, contractAddress common.Address, blockStart uint64, blockEnd uint64) [][]types.Log {
	var logs [][]types.Log
	nblocks := uint64(math.Floor(10_000 / 0.25))
	var blockStarts []uint64

	for i := blockStart; i < blockEnd; i = i + nblocks {
		blockStarts = append(blockStarts, i)
	}

	for _, bs := range blockStarts {
		be := math.Min(float64(bs+nblocks-1), float64(blockEnd))

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

	for _, bs := range blockStarts {
		be := math.Min(float64(bs+nblocks-1), float64(blockEnd))

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

// --------- uPunk Index Calculation ---------

func CalculatePunkIndex(geth *ethclient.Client) map[string]interface{} {
	PUNK_FIRST_BLOCK := 3914495
	// @dev CryptoPunksMarket contract address
	contractAddress := common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB")
	eval_ts := int(time.Now().UTC().Unix())
	window := int(30 * 24 * 60 * 60)

	blockStart := helper.GetBlockNumberByTimestamp(geth, uint64(eval_ts-window), "after")
	blockEnd := helper.GetBlockNumberByTimestamp(geth, uint64(eval_ts), "after")

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
	// var punkSales = make(map[string]map[string]string)

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
					// punkSales[_punkIndex] = info
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
						}
					}

					punkSales = append(punkSales, info)
					// punkSales[_punkIndex] = info
				} else {
					log.Info("Problematic transcation Hash: ", vLog.TxHash)
				}
			}
		}
	}

	var prices []float64
	var uniqueSales = make(map[string]map[string]string)

	for _, ps := range punkSales {
		var currentValues = make(map[string]string)
		var ok bool

		// @dev Iterate over the uniqueSales mapping
		if currentValues, ok = uniqueSales[ps["punkIndex"]]; !ok {
			currentValues = map[string]string{"blockNumber": "0", "price": "0"}
		}

		psBlockNumberString := ps["blockNumber"]
		psBlockNumber, _ := strconv.Atoi(psBlockNumberString)

		currentBlockString := currentValues["blockNumber"]
		currentBlock, _ := strconv.Atoi(currentBlockString)

		if psBlockNumber > currentBlock {
			values := map[string]string{"blockNumber": ps["blockNumber"], "price": ps["price"]}
			uniqueSales[ps["punkIndex"]] = values
		}
	}

	for _, us := range uniqueSales {
		if us["price"] != "" {
			priceFloat64, _ := strconv.ParseFloat(us["price"], 64)
			log.Info(priceFloat64)
			prices = append(prices, priceFloat64)
		}
	}

	// log.Info("PunkSales", len(punkSales))
	// log.Info("UniqueSales", len(uniqueSales))
	// log.Info("Prices", len(prices))

	// @dev Index calculation
	floatMedian, _ := stats.Median(prices)
	bigMedian := big.NewFloat(floatMedian)
	upunkIndexWei := new(big.Float).Quo(bigMedian, big.NewFloat(1000000000000000000))
	upunkIndex := new(big.Float).Quo(upunkIndexWei, big.NewFloat(1000)).String()
	log.Info("Median", floatMedian)
	log.Info("Index", upunkIndex)

	values := map[string]interface{}{"price": upunkIndex, "timestamp": strconv.Itoa(eval_ts)}

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
