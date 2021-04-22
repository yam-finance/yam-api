package routes

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
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

		wethContract, err := weth.NewWeth(common.HexToAddress(contractAddress.WETH), geth)
		if err != nil {
			log.Fatalf("failed to instantiate contract: %v", err)
		}
		tvl, err := wethContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(contractAddress.UGASJUN21))
		if err != nil {
			log.Fatalf("failed to get tvl: %v", err)
		}
		result, _ := utils.BnToDec(tvl, 18).Float64()
		response := &responseDegenerative{
			Values: degenerative{UGAS: map[string]float64{"MAR21": 0, "JUN21": result}},
			Total:  result}
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}
