package routes

import (
	"log"
	"math"
	"math/big"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"

	//eth_rebaser "yam-api/source/contracts"
	eth_rebaser "yam-api/source/contracts/eth_rebaser"
	masterchef "yam-api/source/contracts/masterchef"
	slp "yam-api/source/contracts/slp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

func Tvl(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var ContractIncentivizer = "0xD67c05523D8ec1c60760Fd017Ef006b9F6e496D0"

		eth_rebaserContract, err := eth_rebaser.NewEthRebaser(common.HexToAddress("0xD93f403b432d39aa0f736C2021bE6051d85a1D55"), geth)
		if err != nil {
			log.Fatalf("failed to instantiate contract: %v", err)
		}
		masterChefContract, err := masterchef.NewMasterchef(common.HexToAddress("0xc2edad668740f1aa35e4d8f227fb8e17dca888cd"), geth)
		if err != nil {
			log.Fatalf("failed to instantiate contract: %v", err)
		}
		slpContract, err := slp.NewSlp(common.HexToAddress("0x0f82e57804d0b1f6fab2370a43dcfad3c7cb239c"), geth)
		if err != nil {
			log.Fatalf("failed to instantiate contract: %v", err)
		}

		yamPrice, err := eth_rebaserContract.GetCurrentTWAP(&bind.CallOpts{})
		if err != nil {
			log.Fatalf("failed to get yamprice: %v", err)
		}
		userInfo, err := masterChefContract.UserInfo(&bind.CallOpts{}, big.NewInt(44), common.HexToAddress(ContractIncentivizer))
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

		//	fmt.Println(totalSLPSupply)

		yamValue := utils.BnToDec(totalSLPReserves.Reserve0, 18)
		ethValue := utils.BnToDec(totalSLPReserves.Reserve1, 18)

		yamPriceFloat := utils.BnToDec(yamPrice, 18)
		totalIncentivizerValue := userInfo.Amount
		wethPrice := utils.GetWETHPrice()

		//	fmt.Println(yamPriceFloat)
		//	fmt.Println(wethPrice)
		//	fmt.Println(totalIncentivizerValue)
		temp1 := new(big.Float).Quo(new(big.Float).SetInt(totalIncentivizerValue), new(big.Float).SetInt(totalSLPSupply))
		temp2 := new(big.Float).Add(new(big.Float).Mul(ethValue, wethPrice), new(big.Float).Mul(yamValue, yamPriceFloat))
		temp3 := new(big.Float).Mul(temp1, temp2)
		val, _ := temp3.Float64()
		//math.Round(val)
		utils.ResJSON(http.StatusCreated, w,
			math.Round(val),
		)
	})
}
