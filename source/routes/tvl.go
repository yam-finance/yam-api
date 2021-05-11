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
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"

	eth_rebaser "yam-api/source/contracts/eth_rebaser"
	masterchef "yam-api/source/contracts/masterchef"
	slp "yam-api/source/contracts/slp"
	weth "yam-api/source/contracts/weth"

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
type Assets struct {
	Ugas    []AssetInstance
	Ustonks []AssetInstance
}
type AssetInstance struct {
	Name       string
	Cycle      string
	Year       string
	Collateral string
	Token      Token
	Emp        Emp
	Pool       Pool
	Apr        AprData
}
type Emp struct {
	Address string
	New     bool
}
type Pool struct {
	Address string
}
type AprData struct {
	Force int
	Extra int
}
type Token struct {
	Address  string
	Decimals int
}

func Tvl(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(contractAddress.ContractIncentivizer)

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
		fmt.Println(userInfo)
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

		utils.ResJSON(http.StatusCreated, w,
			math.Round(val),
		)
	})
}
func TvlYam(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(contractAddress.ContractIncentivizer)

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
		fmt.Println(userInfo)
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
		//mapD := map[string]float64{"values": tvl, "total": tvl}
		response := &responseYam{
			Values: yam{Farm: tvl},
			Total:  tvl}

		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func TvlDegenerative(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
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
		wethPrice := utils.GetWETHPrice()
		fmt.Println(wethPrice)

		tvlJun21 := CalculateTvlDegenerative(assets.Ugas[3].Emp.Address, geth, wethPrice)
		tvlMar21 := CalculateTvlDegenerative(assets.Ugas[2].Emp.Address, geth, wethPrice)
		response := &responseDegenerative{
			Values: degenerative{UGAS: map[string]float64{"MAR21": tvlMar21, "JUN21": tvlJun21}},
			Total:  tvlMar21 + tvlJun21}
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
func CalculateTvlDegenerative(empcontractAddress string, geth *ethclient.Client, wethPrice *big.Float) float64 {

	wethContract, err := weth.NewWeth(common.HexToAddress(contractAddress.WETH), geth)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}
	wethBalance, err := wethContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(empcontractAddress))
	if err != nil {
		log.Fatalf("failed to get wethBalance: %v", err)
	}

	wb := utils.BnToDec(wethBalance, 18)
	result, _ := new(big.Float).Mul(wethPrice, wb).Float64()
	return result
}
