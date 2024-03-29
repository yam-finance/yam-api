package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"yam-api/source/config"
	erc20 "yam-api/source/contracts/erc20"
	masterchef "yam-api/source/contracts/masterchef"
	slp "yam-api/source/contracts/slp"
	weth "yam-api/source/contracts/weth"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"
	"yam-api/source/utils/mongodb"

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

// --------- TVL Endpoints ---------
func GetTvlIndex(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "param")
		fmt.Println(param)
		var response map[string]interface{}
		response = make(map[string]interface{})
		if param == "yam" {
			var resultTvlYam map[string]interface{}
			resultTvlYam = make(map[string]interface{})
			tvl := CalculateTvlYam(geth)
			resultTvlYam["farm"] = tvl
			response["values"] = resultTvlYam
			response["total"] = tvl
		}
		if param == "degenerative" {
			values, total := CalculateTvlDegenerativeAll(geth)
			response["values"] = values
			response["total"] = total
		}

		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}

/**
  Calculating TVL
*/
func CalculateTvl(geth *ethclient.Client) map[string]interface{} {
	var response map[string]interface{}
	response = make(map[string]interface{})
	result := map[string]interface{}{}

	tvlYam := CalculateTvlYam(geth)
	result["farm"] = tvlYam
	response["values"] = result
	response["total"] = tvlYam
	return response
}
func StoreTvl(val map[string]interface{}) {
	mongodb.InsertTvl(val)
}
func Tvl(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		response := mongodb.GetTvl()
		if response == nil {
			response = map[string]interface{}{}
		}
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

	resp, err := http.Get("https://raw.githubusercontent.com/yam-finance/synths-sdk/master/src/assets.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(responseData), &AssetsFromFile)

	wethPrice := utils.GetWETHPrice()
	var result map[string]interface{}
	result = make(map[string]interface{})
	var total float64
	total = 0
	v := reflect.ValueOf(AssetsFromFile.Assets)
	values := make([]interface{}, v.NumField())
	assetTypes := reflect.Indirect(reflect.ValueOf(AssetsFromFile.Assets))
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		var assetArray []AssetInstance
		byteData, _ := json.Marshal(values[i])
		json.Unmarshal(byteData, &assetArray)
		var assetTvls map[string]float64
		assetTvls = make(map[string]float64)
		for _, assetItem := range assetArray {

			//if assetItem.Expired != true {
			var assetTvl float64
			if assetItem.Collateral == "WETH" {
				wethPrice = utils.GetWETHPrice()
				assetTvl = CalculateTvlDegenerative(assetItem.Emp.Address, geth, wethPrice, true)
			} else {
				wethPrice = big.NewFloat(1)
				assetTvl = CalculateTvlDegenerative(assetItem.Emp.Address, geth, wethPrice, false)
			}
			assetTvls[assetItem.Cycle] = math.Floor(assetTvl*100) / 100
			total = total + assetTvl
			//	}
		}
		result[assetTypes.Type().Field(i).Name] = assetTvls
	}
	total = math.Floor(total*100) / 100
	return result, total
}
func CalculateTvlYam(geth *ethclient.Client) float64 {
	masterChefContract, err := masterchef.NewMasterchef(common.HexToAddress(contractAddress.Masterchef), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	slpContract, err := slp.NewSlp(common.HexToAddress(contractAddress.Slp), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	yamPrice := utils.GetPriceByContract(contractAddress.YAM)
	fmt.Println("yamPrice = ", yamPrice)
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

	//yamPriceFloat := utils.BnToDec(yamPrice, 18)
	totalIncentivizerValue := userInfo.Amount
	wethPrice := utils.GetWETHPrice()
	tvlBigFloat := new(big.Float).Mul(new(big.Float).Quo(new(big.Float).SetInt(totalIncentivizerValue), new(big.Float).SetInt(totalSLPSupply)), new(big.Float).Add(new(big.Float).Mul(ethValue, wethPrice), new(big.Float).Mul(yamValue, yamPrice)))
	tvlFloat64, _ := tvlBigFloat.Float64()
	tvl := math.Round(tvlFloat64)
	return tvl
}
