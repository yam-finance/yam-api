package web3Helper

import (
	"context"
	"math"
	"math/big"
	"yam-api/source/utils/log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

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

	/// @dev Compute the average block time using the last 10_000 blocks
	avgBlockTime := (currentBlock.Time() - cm100Block.Time()) / nBlocks.Uint64()

	/// @dev If the timestamp is greater than the most recent block, return the most recent block
	if ts > currentBlock.Time() {
		return currentBlockNum
	}

	/// @dev Find bounds on where the block can live and check that it's inbetween the older/newer blocks
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

	/// @dev Find oldest block with timestamp greater than or equal to ts
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
