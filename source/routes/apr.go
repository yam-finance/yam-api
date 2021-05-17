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
		var payload Asset
		rewards = make(map[string]*big.Float)
		//Open assetsJson File
		assetsFile, err := os.Open("assets.json")
		if err != nil {
			log.Fatalf("failed to get jasonfile: %v", err)
		}
		fmt.Println("Successfully Opened assets.json")
		//defer the closing of our jsonFile so that we can parse it later on
		defer assetsFile.Close()

		byteValue, _ := ioutil.ReadAll(assetsFile)
		var assets Assets
		json.Unmarshal([]byte(byteValue), &assets)
		fmt.Println(assets.Ugas[0].Emp.Address)

		emps := utils.GetDevMiningEmps()
		var assetEmpList = []string{assets.Ugas[0].Emp.Address, assets.Ugas[1].Emp.Address, assets.Ugas[2].Emp.Address, assets.Ugas[3].Emp.Address, assets.Ustonks[0].Emp.Address}
		emps.EmpWhiteList = utils.MergeUnique(emps.EmpWhiteList, assetEmpList)
		fmt.Println(emps.EmpWhiteList)
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
			fmt.Println(empInfo.Address)
			fmt.Println(value)
		}
		for index, value := range values {
			x := new(big.Float).Mul(&value, big.NewFloat(emps.TotalReward))
			y := new(big.Float).Quo(x, totalValue)
			fmt.Println(y)
			rewards[allEmpInfo[index].Address] = y
		}

		payload.AssetName = "UGAS"
		payload.AssetInstance = assets.Ugas[3]
		payload.AssetPrice = GetUniPrice(payload.AssetInstance.Token.Address, contractAddress.WETH, geth)

		utils.ResJSON(http.StatusCreated, w,
			CalculateApr(payload, geth),
		)
	})
}
func StoreAprYam(val float64) {
	mongodb.InsertAprYam(val)
}
func StoreAprDegenerative(val map[string]float64) {
	mongodb.InsertAprDegenerative(val)
}
func AprYam(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		val := mongodb.GetAprYam()
		if val == 0 {
			val = CalculateAprYam(geth)
			StoreAprYam(val)
		}
		response := &responseAprYam{
			Value: val}
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func CalculateAprYam(geth *ethclient.Client) float64 {
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
	return val

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

		val := mongodb.GetAprDegenerative()
		if val == nil {
			val = CalculateAprDegenerative(geth).UGAS
			StoreAprDegenerative(val)
		}
		response := &responseAprDegenerative{
			UGAS: val,
		}

		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func CalculateAprDegenerative(geth *ethclient.Client) *responseAprDegenerative {
	var payload Asset
	rewards = make(map[string]*big.Float)
	//Open assetsJson File
	assetsFile, err := os.Open("assets.json")

	if err != nil {
		log.Fatalf("failed to get jasonfile: %v", err)
	}

	fmt.Println("Successfully Opened assets.json")

	//defer the closing of our jsonFile so that we can parse it later on
	defer assetsFile.Close()
	byteValue, _ := ioutil.ReadAll(assetsFile)
	var assets Assets
	json.Unmarshal([]byte(byteValue), &assets)
	fmt.Println(assets.Ugas[0].Emp.Address)

	emps := utils.GetDevMiningEmps()
	var assetEmpList = []string{assets.Ugas[0].Emp.Address, assets.Ugas[1].Emp.Address, assets.Ugas[2].Emp.Address, assets.Ugas[3].Emp.Address, assets.Ustonks[0].Emp.Address}
	emps.EmpWhiteList = utils.MergeUnique(emps.EmpWhiteList, assetEmpList)
	fmt.Println(emps.EmpWhiteList)
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
		fmt.Println(empInfo.Address)
		fmt.Println(value)
	}
	for index, value := range values {
		x := new(big.Float).Mul(&value, big.NewFloat(emps.TotalReward))
		y := new(big.Float).Quo(x, totalValue)
		fmt.Println(y)
		rewards[allEmpInfo[index].Address] = y
	}

	payload.AssetName = "UGAS"
	payload.AssetInstance = assets.Ugas[3]
	payload.AssetPrice = GetUniPrice(payload.AssetInstance.Token.Address, contractAddress.WETH, geth)
	ugasJunApr, _ := CalculateApr(payload, geth).Float64()
	payload.AssetName = "UGAS"
	payload.AssetInstance = assets.Ugas[2]
	payload.AssetPrice = GetUniPrice(payload.AssetInstance.Token.Address, contractAddress.WETH, geth)

	ugasMarApr, _ := CalculateApr(payload, geth).Float64()

	response := &responseAprDegenerative{
		UGAS: map[string]float64{"MAR21": ugasMarApr, "JUN21": ugasJunApr},
	}
	return response
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
func GetUniPrice(tokenA string, tokenB string, geth *ethclient.Client) *big.Float {
	uniFact := GetUNIFact(geth)
	pair, err := uniFact.GetPair(nil, common.HexToAddress(tokenA), common.HexToAddress(tokenB))
	if err != nil {
		log.Fatalf("failed to get pair: %v", err)
	}
	uniPair := GetUni(pair.Hex(), geth)
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

	if strings.ToLower(token0.String()) == strings.ToLower(tokenA) {
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
