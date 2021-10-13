package routes

import (
	"fmt"
	"math/big"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"
	"yam-api/source/utils/mongodb"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

func Treasury(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		//response := GetTreasury(geth)
		response := mongodb.GetTreasury()
		if response == nil {
			response = map[string]interface{}{}
		}
		utils.ResJSON(http.StatusCreated, w,
			response,
		)
	})
}

func GetTreasury(geth *ethclient.Client) map[string]interface{} {
	var response map[string]interface{}

	response = make(map[string]interface{})

	totalUMATreasuryValue := utils.GetBalance(contractAddress.UMA, contractAddress.TreasuryAddress, geth, 18)
	totalUMAMultisigValue := utils.GetBalance(contractAddress.UMA, contractAddress.MultisigUMAAddress, geth, 18)
	totalUMAValue := new(big.Float).Add(totalUMATreasuryValue, totalUMAMultisigValue)

	totalYamHouseValue := utils.GetBalance(contractAddress.YamDaoHouse, contractAddress.TreasuryAddress, geth, 18)
	totalDPIValue := utils.GetBalance(contractAddress.DPI, contractAddress.TreasuryAddress, geth, 18)
	totalWETHTokenBalance := utils.GetBalance(contractAddress.WETH, contractAddress.TreasuryAddress, geth, 18)
	totalWETHValue := new(big.Float).Add(totalWETHTokenBalance, big.NewFloat(283))
	totalBalanceIndexCoop := utils.GetBalance(contractAddress.INDEX, contractAddress.ContractTimelock, geth, 18)

	//	yamBalance := utils.GetBalance(contractAddress.Yamv3, contractAddress.TreasuryAddress, geth, 18)
	yUsdBalance := utils.GetBalance(contractAddress.YUsd, contractAddress.TreasuryAddress, geth, 18)
	usdcMultisigBalance := utils.GetBalance(contractAddress.USDC, contractAddress.MultisigAddress, geth, 6)
	//  ustonks and usdc tokens of what we have in the USTONKS SEP pool
	usdcSLPBalance := utils.GetBalance(contractAddress.USDC, contractAddress.USTONKSSEPPool, geth, 6)

	//gitcoin balance
	gitcoinBalance := utils.GetBalance(contractAddress.GitCoinAddress, contractAddress.TreasuryAddress, geth, 18)

	yamTwap := utils.GetPriceByContract(contractAddress.Yamv3)
	yusdPrice := utils.GetValue("yvault-lp-ycurve")
	wethPrice := utils.GetWETHPrice()
	dpiPrice := utils.GetValue("defipulse-index")
	fmt.Println("dpi price = ", dpiPrice)
	indexPrice := utils.GetValue("index-cooperative")
	umaPrice := utils.GetValue("uma")
	sushiPrice := utils.GetValue("sushi")
	usdcPrice := utils.GetValue("usd-coin")
	yamHousePrice := utils.GetYamHousePrice()
	rewardsIndexCoop := utils.GetIndexCoopLPRewards(geth)
	rewardsSushi := utils.GetSushiRewards(geth)
	gitPrice := utils.GetPriceByContract(contractAddress.GitCoinAddress)

	//	yamYUsdValue := new(big.Float).Mul(yamTwap, yamBalance)
	fmt.Println("yamTwap=", yamTwap)
	/*	if yUsdBalance != nil {
			totalYUsdValue = new(big.Float).Add(yUsdBalance, new(big.Float).Add(yamYUsdValue, big.NewFloat(718900)))
		} else {
			totalYUsdValue = new(big.Float).Add(yamYUsdValue, big.NewFloat(718900))
		}*/
	//wethTreasury := new(big.Float).Mul(totalWETHValue, wethPrice)
	//	yUSDTreasury := new(big.Float).Mul(totalYUsdValue, yusdPrice)
	//	dpiTreasury := new(big.Float).Mul(totalDPIValue, dpiPrice)
	//	indexTreasury := new(big.Float).Mul(new(big.Float).Add(rewardsIndexCoop, totalBalanceIndexCoop), indexPrice)
	//	indexLPTreasury := new(big.Float).Add(new(big.Float).Mul(dpiPrice, big.NewFloat(2929)), new(big.Float).Mul(wethPrice, big.NewFloat(640)))
	//	umaTreasury := new(big.Float).Mul(totalUMAValue, umaPrice)
	//	yamHouseTreasury := new(big.Float).Mul(totalYamHouseValue, yamHousePrice)
	//	sushiTreasury := new(big.Float).Mul(rewardsSushi, sushiPrice)
	//	gitCoinTreasury := new(big.Float).Mul(gitcoinBalance, gitPrice)

	//	yUsdTreasury_yUSD := new(big.Float).Mul(yUsdBalance, yusdPrice)
	//	yUsdTreasury_USDC := usdcMultisigBalance
	yUsdTreasury_USTONKSLP := new(big.Float).Mul(usdcSLPBalance, big.NewFloat(2))

	change24UMA := utils.GetValueChange("uma")
	change24DPI := utils.GetValueChange("defipulse-index")
	change24WETH := utils.GetValueChange("weth")
	change24YUSD := utils.GetValueChange("yvault-lp-ycurve")
	change24IndexCoop := utils.GetValueChange("index-cooperative")
	change24Sushi := utils.GetValueChange("sushi")
	change24YAMAHOUSE := big.NewFloat(0)
	change24USDC := utils.GetValueChange("usd-coin")
	change24GTC := utils.GetValueChange("gitcoin")

	response["UMA"] = utils.SetTreasuryAssetInfo(totalUMAValue, umaPrice, change24UMA)

	response["YAMHOUSE"] = utils.SetTreasuryAssetInfo(totalYamHouseValue, yamHousePrice, change24YAMAHOUSE)

	response["DPI"] = utils.SetTreasuryAssetInfo(totalDPIValue, dpiPrice, change24DPI)

	response["WETH"] = utils.SetTreasuryAssetInfo(totalWETHValue, wethPrice, change24WETH)

	response["YUSD"] = utils.SetTreasuryAssetInfo(yUsdBalance, yusdPrice, change24YUSD)

	response["USDC"] = utils.SetTreasuryAssetInfo(usdcMultisigBalance, usdcPrice, change24USDC)

	//later need to modify

	response["USTONKSLP"] = utils.SetTreasuryAssetInfo(yUsdTreasury_USTONKSLP, usdcPrice, change24USDC)

	response["INDEX"] = utils.SetTreasuryAssetInfo(totalBalanceIndexCoop, indexPrice, change24IndexCoop)

	response["INDEXLP"] = utils.SetTreasuryAssetInfo(rewardsIndexCoop, indexPrice, change24IndexCoop)

	response["SUSHI"] = utils.SetTreasuryAssetInfo(rewardsSushi, sushiPrice, change24Sushi)

	response["GITCOIN"] = utils.SetTreasuryAssetInfo(gitcoinBalance, gitPrice, change24GTC)

	return response

}
func StoreTreasury(val map[string]interface{}) {
	mongodb.InsertTreasury(val)
}
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
