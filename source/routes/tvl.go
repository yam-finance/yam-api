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
	"yam-api/source/config"
	erc20 "yam-api/source/contracts/erc20"
	eth_rebaser "yam-api/source/contracts/eth_rebaser"
	masterchef "yam-api/source/contracts/masterchef"
	slp "yam-api/source/contracts/slp"
	weth "yam-api/source/contracts/weth"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

type yam struct {
	Farm float64 `json:"farm"`
}
type responseYam struct {
	Values yam     `json:"values"`
	Total  float64 `json:"total"`
}
type degenerative struct {
	UGAS map[string]float64
	UVOL map[string]float64
}
type responseDegenerative struct {
	Values degenerative `json:"values"`
	Total  float64      `json:"total"`
}

func Tvl(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		response = make(map[string]interface{})
		tvlYam := CalculateTvlYam(geth)
		values, total := CalculateTvlDegenerativeAll(geth)
		result := map[string]interface{}{}

		result["farm"] = tvlYam
		result["UGAS"] = values["UGAS"]
		result["USTONKS"] = values["USTONKS"]
		result["UPUNKS"] = values["UPUNKS"]
		response["values"] = result
		response["total"] = tvlYam + total
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func TvlYam(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		var resultTvlYam map[string]interface{}
		response = make(map[string]interface{})
		resultTvlYam = make(map[string]interface{})
		tvl := CalculateTvlYam(geth)
		resultTvlYam["farm"] = tvl
		response["values"] = resultTvlYam
		response["total"] = tvl
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func TvlDegenerative(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		response = make(map[string]interface{})
		values, total := CalculateTvlDegenerativeAll(geth)
		response["values"] = values
		response["total"] = total
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func CalculateTvlDegenerative(empcontractAddress string, geth *ethclient.Client, wethPrice *big.Float, isWeth bool) float64 {

	wethContract, err := weth.NewWeth(common.HexToAddress(contractAddress.WETH), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	usdcContract, err := erc20.NewErc20(common.HexToAddress(contractAddress.USDC), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	wethBalance, err := wethContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(empcontractAddress))
	if err != nil {
		log.Fatalf("failed to get wethBalance: %v", err)
	}
	usdcBalance, err := usdcContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(empcontractAddress))
	if err != nil {
		log.Fatalf("failed to get wethBalance: %v", err)
	}
	var wb *big.Float
	if isWeth == true {
		wb = utils.BnToDec(wethBalance, 18)
	} else {
		wb = utils.BnToDec(usdcBalance, 6)
	}
	result, _ := new(big.Float).Mul(wethPrice, wb).Float64()
	return result
}
func CalculateTvlDegenerativeAll(geth *ethclient.Client) (map[string]interface{}, float64) {
	assetsFile, err := os.Open("assets.json")

	if err != nil {
		log.Fatalf("failed to get jasonfile: %v", err)
	}

	fmt.Println("Successfully Opened assets.json")

	//defer the closing of our jsonFile so that we can parse it later on
	defer assetsFile.Close()
	byteValue, _ := ioutil.ReadAll(assetsFile)
	json.Unmarshal([]byte(byteValue), &AssetsFromFile)

	wethPrice := utils.GetWETHPrice()
	var UGAS map[string]float64
	var USTONKS map[string]float64
	var UPUNKS map[string]float64
	var result map[string]interface{}
	UGAS = make(map[string]float64)
	USTONKS = make(map[string]float64)
	UPUNKS = make(map[string]float64)
	result = make(map[string]interface{})
	var total float64
	total = 0
	//////  Calculating UGAS ///////
	for _, ugas := range AssetsFromFile.Assets.Ugas {
		if ugas.Expired != true {
			var ugasTvl float64
			if ugas.Collateral == "WETH" {
				wethPrice = utils.GetWETHPrice()
				ugasTvl = CalculateTvlDegenerative(ugas.Emp.Address, geth, wethPrice, true)
			} else {
				wethPrice = big.NewFloat(1)
				ugasTvl = CalculateTvlDegenerative(ugas.Emp.Address, geth, wethPrice, false)
			}
			UGAS[ugas.Cycle] = ugasTvl
			total = total + ugasTvl
		}
	}
	//////  Calculating USTONKS ///////
	for _, ustonks := range AssetsFromFile.Assets.Ustonks {
		if ustonks.Expired != true {
			var ustonksTvl float64
			if ustonks.Collateral == "WETH" {
				wethPrice = utils.GetWETHPrice()
				ustonksTvl = CalculateTvlDegenerative(ustonks.Emp.Address, geth, wethPrice, true)
			} else {
				wethPrice = big.NewFloat(1)
				ustonksTvl = CalculateTvlDegenerative(ustonks.Emp.Address, geth, wethPrice, false)
			}
			USTONKS[ustonks.Cycle] = ustonksTvl
			total = total + ustonksTvl
		}
	}
	//////  Calculating UPUNKS ///////
	for _, upunks := range AssetsFromFile.Assets.Upunks {
		if upunks.Expired != true {
			var upunksTvl float64
			if upunks.Collateral == "WETH" {
				wethPrice = utils.GetWETHPrice()
				upunksTvl = CalculateTvlDegenerative(upunks.Emp.Address, geth, wethPrice, true)
			} else {
				wethPrice = big.NewFloat(1)
				upunksTvl = CalculateTvlDegenerative(upunks.Emp.Address, geth, wethPrice, false)
			}
			UPUNKS[upunks.Cycle] = upunksTvl
			total = total + upunksTvl
		}
	}
	result["UGAS"] = UGAS
	result["USTONKS"] = USTONKS
	result["UPUNKS"] = UPUNKS
	return result, total
}
func CalculateTvlYam(geth *ethclient.Client) float64 {
	eth_rebaserContract, err := eth_rebaser.NewEthRebaser(common.HexToAddress(contractAddress.Eth_rebaser), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	masterChefContract, err := masterchef.NewMasterchef(common.HexToAddress(contractAddress.Masterchef), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	slpContract, err := slp.NewSlp(common.HexToAddress(contractAddress.Slp), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}

	yamPrice, err := eth_rebaserContract.GetCurrentTWAP(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("failed to get yamprice: %v", err)
	}
	userInfo, err := masterChefContract.UserInfo(&bind.CallOpts{}, big.NewInt(44), common.HexToAddress(contractAddress.ContractIncentivizer))
	if err != nil {
		log.Fatalf("failed to get amount: %v", err)
	}

	totalSLPSupply, err := slpContract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("failed to get totalSLPSupply: %v", err)
	}
	totalSLPReserves, err := slpContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("failed to get totalSLPReserves: %v", err)
	}

	yamValue := utils.BnToDec(totalSLPReserves.Reserve0, 18)
	ethValue := utils.BnToDec(totalSLPReserves.Reserve1, 18)

	yamPriceFloat := utils.BnToDec(yamPrice, 18)
	totalIncentivizerValue := userInfo.Amount
	wethPrice := utils.GetWETHPrice()

	temp1 := new(big.Float).Quo(new(big.Float).SetInt(totalIncentivizerValue), new(big.Float).SetInt(totalSLPSupply))
	temp2 := new(big.Float).Add(new(big.Float).Mul(ethValue, wethPrice), new(big.Float).Mul(yamValue, yamPriceFloat))
	temp3 := new(big.Float).Mul(temp1, temp2)
	val, _ := temp3.Float64()
	tvl := math.Round(val)
	return tvl
}
