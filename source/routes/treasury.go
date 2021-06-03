package routes

import (
	"math/big"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

func Treasury(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		var totalYUsdValue *big.Float
		response = make(map[string]interface{})
		totalUMAValue := utils.GetBalance(contractAddress.UMA, contractAddress.TreasuryAddress, geth, 18)
		totalYamHouseValue := utils.GetBalance(contractAddress.YamDaoHouse, contractAddress.TreasuryAddress, geth, 18)
		totalDPIValue := utils.GetBalance(contractAddress.DPI, contractAddress.TreasuryAddress, geth, 18)
		totalWETHTokenBalance := utils.GetBalance(contractAddress.WETH, contractAddress.TreasuryAddress, geth, 18)
		totalWETHValue := new(big.Float).Add(totalWETHTokenBalance, big.NewFloat(283))
		totalBalanceIndexCoop := utils.GetBalance(contractAddress.INDEX, contractAddress.TreasuryAddress, geth, 18)

		yamBalance := utils.GetBalance(contractAddress.Yamv3, contractAddress.TreasuryAddress, geth, 18)
		yUsdBalance := utils.GetBalance(contractAddress.YUsd, contractAddress.TreasuryAddress, geth, 18)

		yamTwap := utils.GetPriceByContract(contractAddress.Yamv3)
		yusdPrice := utils.GetYUSDPrice()
		wethPrice := utils.GetWETHPrice()
		dpiPrice := utils.GetValue("defipulse-index")
		indexPrice := utils.GetValue("index-cooperative")
		umaPrice := utils.GetValue("uma")
		sushiPrice := utils.GetValue("sushi")
		yamHousePrice := utils.GetYamHousePrice()
		rewardsIndexCoop := utils.GetIndexCoopLPRewards(geth)
		rewardsSushi := utils.GetSushiRewards(geth)

		yamYUsdValue := new(big.Float).Mul(yamTwap, yamBalance)

		if yUsdBalance != nil {
			totalYUsdValue = new(big.Float).Add(yUsdBalance, new(big.Float).Add(yamYUsdValue, big.NewFloat(718900)))
		} else {
			totalYUsdValue = new(big.Float).Add(yamYUsdValue, big.NewFloat(718900))
		}
		wethTreasury := new(big.Float).Mul(totalWETHValue, wethPrice)
		yUSDTreasury := new(big.Float).Mul(totalYUsdValue, yusdPrice)
		dpiTreasury := new(big.Float).Mul(totalDPIValue, dpiPrice)
		indexTreasury := new(big.Float).Mul(new(big.Float).Add(rewardsIndexCoop, totalBalanceIndexCoop), indexPrice)
		indexLPTreasury := new(big.Float).Add(new(big.Float).Mul(dpiPrice, big.NewFloat(2929)), new(big.Float).Mul(wethPrice, big.NewFloat(640)))
		umaTreasury := new(big.Float).Mul(totalUMAValue, umaPrice)
		yamHouseTreasury := new(big.Float).Mul(totalYamHouseValue, yamHousePrice)
		sushiTreasury := new(big.Float).Mul(rewardsSushi, sushiPrice)
		response["uma"] = umaTreasury
		response["yamHouse"] = yamHouseTreasury
		response["dpi"] = dpiTreasury
		response["weth"] = wethTreasury
		response["yyDAI+"], _ = yUSDTreasury.Float64()
		response["index"] = indexTreasury
		response["indexlp"] = indexLPTreasury
		response["sushi"] = sushiTreasury

		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}

func GetTreasury() {}

func GetTreasuryCharts() {}

func marketCap() {}

func treasuryValue() {}

func totalSupply() {}

func toRebaseYam() {}

func scalingFactor() {}

func currentPrice() {}

func historyReserves() {}

func historyScalingFactor() {}

func historySoldPerRebase() {}

func historyMintedPerRebase() {}
