package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
	"yam-api/source/config"
	emp "yam-api/source/contracts/emp"
	erc20 "yam-api/source/contracts/erc20"
	eth_rebaser "yam-api/source/contracts/eth_rebaser"
	uni "yam-api/source/contracts/uni"
	unifact "yam-api/source/contracts/unifact"
	yamv3 "yam-api/source/contracts/yamv3"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"
	"yam-api/source/utils/mongodb"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/robfig/cron"
)

var rewards map[string]*big.Float
var eth_rebaserContract *eth_rebaser.EthRebaser
var yamv3Contract *yamv3.Yamv3
var getAprYamCron *cron.Cron
var getAprDegenerativeCron *cron.Cron

func Apr(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		response = make(map[string]interface{})
		aprYam := mongodb.GetAprYam()
		if aprYam == nil {
			aprYam = map[string]interface{}{}
		}
		aprDegenerative := mongodb.GetAprDegenerative()
		if aprDegenerative == nil {
			aprDegenerative = map[string]interface{}{}
		}

		response["yam"] = aprYam
		response["degenerative"] = aprDegenerative
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func StoreAprYam(val map[string]interface{}) {
	mongodb.InsertAprYam(val)
}
func StoreAprDegenerative(val map[string]interface{}) {
	mongodb.InsertAprDegenerative(val)
}
func AprYam(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		response := mongodb.GetAprYam()
		if response == nil {
			response = map[string]interface{}{}
		}

		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func CalculateAprYam(geth *ethclient.Client) map[string]interface{} {
	var result map[string]interface{}
	result = make(map[string]interface{})
	var err error
	var BoU = big.NewFloat(5000)
	if eth_rebaserContract == nil {
		eth_rebaserContract, err = eth_rebaser.NewEthRebaser(common.HexToAddress(contractAddress.Eth_rebaser), geth)

	}
	yamPrice, err := eth_rebaserContract.GetCurrentTWAP(nil)
	yamPriceFloat := utils.BnToDec(yamPrice, 18)
	yamScalingFactor := utils.BnToDec(getScalingFactor(geth), 18)
	yamTvl := CalculateTvlYam(geth)

	temp1 := new(big.Float).Mul(new(big.Float).Mul(BoU, yamPriceFloat), new(big.Float).Mul(big.NewFloat(365.5), yamScalingFactor))
	temp2 := new(big.Float).Quo(new(big.Float).Mul(temp1, big.NewFloat(100)), new(big.Float).Mul(big.NewFloat(7), big.NewFloat(yamTvl)))
	if err != nil {
		log.Fatalf("failed to get yamprice: %v", err)
	}
	val, _ := temp2.Float64()
	result["farm"] = val
	return result

}
func getScalingFactor(geth *ethclient.Client) *big.Int {
	var err error
	if yamv3Contract == nil {
		yamv3Contract, err = yamv3.NewYamv3(common.HexToAddress(contractAddress.Yamv3), geth)
	}
	yamScalingFactor, err := yamv3Contract.YamsScalingFactor(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("failed to get yamScalingFactor: %v", err)
	}
	return yamScalingFactor

}
func AprDegenerative(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		response := mongodb.GetAprDegenerative()
		if response == nil {
			response = map[string]interface{}{}
		}

		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func CalculateAprDegenerative(geth *ethclient.Client) map[string]interface{} {
	var payload Asset
	rewards = make(map[string]*big.Float)
	//Open assetsJson File
	assetsFile, err := os.Open("assets.json")

	if err != nil {
		log.Fatalf("failed to get jasonfile: %v", err)
	}

	//defer the closing of our jsonFile so that we can parse it later on
	defer assetsFile.Close()
	byteValue, _ := ioutil.ReadAll(assetsFile)
	//	var assets Assets
	json.Unmarshal([]byte(byteValue), &AssetsFromFile)
	emps := utils.GetDevMiningEmps()
	var assetEmpList = []string{}
	for _, ugas := range AssetsFromFile.Assets.Ugas {
		assetEmpList = append(assetEmpList, ugas.Emp.Address)
	}
	for _, ustonks := range AssetsFromFile.Assets.Ustonks {
		assetEmpList = append(assetEmpList, ustonks.Emp.Address)
	}
	for _, upunks := range AssetsFromFile.Assets.Upunks {
		assetEmpList = append(assetEmpList, upunks.Emp.Address)
	}

	emps.EmpWhiteList = utils.MergeUnique(emps.EmpWhiteList, assetEmpList)

	var allEmpInfo []EmpInfo
	for _, empAddr := range emps.EmpWhiteList {
		empInfo := GetEmpInfo(empAddr, "usd", geth)
		allEmpInfo = append(allEmpInfo, empInfo)
	}
	var totalValue = big.NewFloat(0)

	var values []big.Float

	for _, empInfo := range allEmpInfo {
		value := CalculateEmpValue(empInfo.Price, &empInfo.Size, empInfo.Decimals)
		totalValue = new(big.Float).Add(totalValue, value)
		values = append(values, *value)
	}
	for index, value := range values {
		x := new(big.Float).Mul(&value, big.NewFloat(emps.TotalReward))
		y := new(big.Float).Quo(x, totalValue)
		rewards[allEmpInfo[index].Address] = y
	}
	var UGAS map[string]float64
	var USTONKS map[string]float64
	var UPUNKS map[string]float64
	var result map[string]interface{}
	result = make(map[string]interface{})
	UGAS = make(map[string]float64)
	USTONKS = make(map[string]float64)
	UPUNKS = make(map[string]float64)

	//////  Calculating UGAS ///////
	for _, ugas := range AssetsFromFile.Assets.Ugas {
		if ugas.Expired != true {
			payload.AssetName = "UGAS"
			payload.AssetInstance = ugas
			payload.AssetPrice = GetUniPrice(payload.AssetInstance.Token.Address, payload.AssetInstance.Pool.Address, geth)
			ugasApr, _ := CalculateApr(payload, geth).Float64()
			UGAS[ugas.Cycle] = ugasApr
		}
	}
	//////  Calculating USTONKS ///////
	for _, ustonks := range AssetsFromFile.Assets.Ustonks {
		if ustonks.Expired != true {
			payload.AssetName = "USTONKS"
			payload.AssetInstance = ustonks
			payload.AssetPrice = GetUniPrice(payload.AssetInstance.Token.Address, payload.AssetInstance.Pool.Address, geth)
			ustonksApr, _ := CalculateApr(payload, geth).Float64()
			USTONKS[ustonks.Cycle] = ustonksApr
		}
	}
	//////  Calculating UPUNKS ///////
	for _, upunks := range AssetsFromFile.Assets.Upunks {
		if upunks.Expired != true {
			payload.AssetName = "UPUNKS"
			payload.AssetInstance = upunks
			payload.AssetPrice = GetUniPrice(payload.AssetInstance.Token.Address, payload.AssetInstance.Pool.Address, geth)
			upunksApr, _ := CalculateApr(payload, geth).Float64()
			UPUNKS[upunks.Cycle] = upunksApr
		}
	}

	result["UGAS"] = UGAS
	result["USTONKS"] = USTONKS
	result["UPUNKS"] = UPUNKS
	fmt.Println(result)

	return result
}
func CalculateApr(payload Asset, geth *ethclient.Client) *big.Float {
	contractEmp, err := emp.NewEmp(common.HexToAddress(payload.AssetInstance.Emp.Address), geth)
	if err != nil {
		log.Fatalf("failed to instantiate emp contract: %v", err)
	}
	contractLp, err := uni.NewUni(common.HexToAddress(payload.AssetInstance.Pool.Address), geth)
	if err != nil {
		log.Fatalf("failed to instantiate emp contract: %v", err)
	}
	contractEmpCall, err := contractEmp.RawTotalPositionCollateral(nil)
	if err != nil {
		log.Fatalf("failed to get rawtototalpositioncollateral: %v", err)
	}
	contractLpCall, err := contractLp.GetReserves(nil)
	if err != nil {
		log.Fatalf("failed to get Reserves: %v", err)
	}

	umaRewards := rewards[payload.AssetInstance.Emp.Address]
	umaPrice := utils.GetPriceByContract(contractAddress.UMA)
	yamPrice := utils.GetPriceByContract(contractAddress.Yamv3)
	ethPrice := utils.GetPriceByContract(contractAddress.WETH)
	var tokenPrice *big.Float
	var baseCollateral float64

	if payload.AssetInstance.Collateral == "USDC" {
		baseCollateral = math.Pow10(6)
		tokenPrice = payload.AssetPrice

	} else {
		baseCollateral = math.Pow10(18)
		tokenPrice = (new(big.Float).Mul(payload.AssetPrice, ethPrice))
	}
	currentTime := time.Now().Unix()
	var week1Until = big.NewFloat(1615665600)
	var week2Until = big.NewFloat(1616961600)
	var yamWeekRewards = big.NewFloat(0)
	var umaWeekRewards = big.NewFloat(0)
	if payload.AssetName == "UGAS" && payload.AssetInstance.Cycle == "JUN" && payload.AssetInstance.Year == "21" {
		ct := new(big.Float).SetInt64(currentTime)

		if ct.Cmp(week1Until) == -1 {
			yamWeekRewards = (new(big.Float).Add(yamWeekRewards, big.NewFloat(5000)))
		} else if ct.Cmp(week2Until) == -1 {
			yamWeekRewards = (new(big.Float).Add(yamWeekRewards, big.NewFloat(10000)))
		}
	} else if payload.AssetName == "USTONKS" && payload.AssetInstance.Cycle == "APR" && payload.AssetInstance.Year == "21" {
		ct := new(big.Float).SetInt64(currentTime)
		if ct.Cmp(week1Until) == -1 {
			yamWeekRewards = (new(big.Float).Add(yamWeekRewards, big.NewFloat(5000)))
			umaWeekRewards = (new(big.Float).Add(umaWeekRewards, big.NewFloat(5000)))
		} else if ct.Cmp(week2Until) == -1 {
			yamWeekRewards = (new(big.Float).Add(yamWeekRewards, big.NewFloat(10000)))
			umaWeekRewards = (new(big.Float).Add(umaWeekRewards, big.NewFloat(10000)))
		}
	}

	fmt.Println("token Price")
	fmt.Println(tokenPrice)
	yamReward := big.NewFloat(0)

	var calcAsset = big.NewFloat(0)
	var calcCollateral = big.NewFloat(0)

	normalRewards := new(big.Float).Add(new(big.Float).Mul(umaRewards, umaPrice), new(big.Float).Mul(yamReward, yamPrice))
	weekRewards := new(big.Float).Add(new(big.Float).Mul(umaWeekRewards, umaPrice), new(big.Float).Mul(yamWeekRewards, yamPrice))
	assetReserve0 := utils.BnToDec(contractLpCall.Reserve0, payload.AssetInstance.Token.Decimals)
	assetReserve1 := new(big.Float).Quo(new(big.Float).SetInt(contractLpCall.Reserve1), big.NewFloat(baseCollateral))

	if payload.AssetName == "USTONKS" {
		if payload.AssetInstance.Cycle == "APR" {
			calcAsset = new(big.Float).Mul(assetReserve1, tokenPrice)
			if payload.AssetInstance.Collateral == "WETH" {
				calcCollateral = new(big.Float).Mul(assetReserve0, ethPrice)
			} else {
				calcCollateral = assetReserve0
			}
		} else {
			calcAsset = new(big.Float).Mul(assetReserve0, tokenPrice)
			if payload.AssetInstance.Collateral == "WETH" {
				calcCollateral = new(big.Float).Mul(assetReserve1, ethPrice)
			} else {
				calcCollateral = assetReserve1
			}
		}

	} else {
		calcAsset = new(big.Float).Mul(assetReserve0, tokenPrice)
		if payload.AssetInstance.Collateral == "WETH" {
			calcCollateral = new(big.Float).Mul(assetReserve1, ethPrice)
		} else {
			calcCollateral = assetReserve1
		}
	}

	empTvl := utils.BnToDec(contractEmpCall, payload.AssetInstance.Token.Decimals)
	if payload.AssetInstance.Collateral == "WETH" {
		empTvl = new(big.Float).Mul(empTvl, ethPrice)
	}
	uniLpPair := new(big.Float).Add(calcAsset, calcCollateral)
	assetReserveValue := new(big.Float).Add(empTvl, new(big.Float).Mul(uniLpPair, big.NewFloat(0.5)))
	aprCalculate := new(big.Float).Quo(new(big.Float).Mul(normalRewards, big.NewFloat(52*0.82*100)), assetReserveValue)
	aprCalculateExtra := new(big.Float).Quo(new(big.Float).Mul(weekRewards, big.NewFloat(52*100)), assetReserveValue)
	totalAprCalculation := new(big.Float).Add(aprCalculate, aprCalculateExtra)
	return totalAprCalculation
}
func CalculateEmpValue(price big.Float, size *big.Int, decimals int) *big.Float {
	fixedPrice := price
	fixedSize := utils.BnToDec(size, decimals)
	ret := new(big.Float).Mul(&fixedPrice, fixedSize)
	return ret
}
func GetUniPrice(tokenA string, pair string, geth *ethclient.Client) *big.Float {
	uniPair := GetUni(pair, geth)
	token0, err := uniPair.Token0(nil)

	if err != nil {
		log.Fatalf("failed to get token0: %v", err)
	}
	res, err := uniPair.GetReserves(nil)
	if err != nil {
		log.Fatalf("failed to get reserves: %v", err)
	}
	reserves0 := res.Reserve0
	reserves1 := res.Reserve1
	ret := big.NewFloat(0)
	if (strings.ToLower(token0.String()) == strings.ToLower(tokenA)) || (token0.String() == contractAddress.USDC) {
		x, y := new(big.Float).SetInt(reserves0), new(big.Float).SetInt(reserves1)
		ret = new(big.Float).Quo(x, y)

	} else {
		x, y := new(big.Float).SetInt(reserves0), new(big.Float).SetInt(reserves1)
		ret = new(big.Float).Quo(y, x)
	}
	fmt.Println(ret)
	return ret

}
func GetUNIFact(geth *ethclient.Client) *unifact.Unifact {
	uniFactContract, err := unifact.NewUnifact(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}

	return uniFactContract
}
func GetUni(address string, geth *ethclient.Client) *uni.Uni {
	uniContract, err := uni.NewUni(common.HexToAddress(address), geth)
	if err != nil {
		log.Fatalf("failed to instantiate uni contract: %v", err)
	}
	return uniContract
}
func GetEmpInfo(address string, toCurrency string, geth *ethclient.Client) EmpInfo {
	var ret EmpInfo
	empContract, err := emp.NewEmp(common.HexToAddress(address), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}

	collateralAddress, err := empContract.CollateralCurrency(nil)
	if err != nil {
		log.Fatalf("failed to get collateralAddress: %v", err)
	}
	erc20Contract, err := erc20.NewErc20(collateralAddress, geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	size, err := empContract.RawTotalPositionCollateral(nil)
	if err != nil {
		log.Fatalf("failed to get size: %v", err)
	}
	decimals, err := erc20Contract.Decimals(nil)
	if err != nil {
		log.Fatalf("failed to get size: %v", err)
	}
	price := utils.GetPriceByContract(collateralAddress.Hex())
	ret.Address = address
	ret.Size = *size
	ret.Price = *price
	ret.Decimals = int(decimals)
	ret.CollateralAddress = collateralAddress.Hex()
	return ret

}
