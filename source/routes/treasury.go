package routes

// Yam Treasury Assets https://github.com/yam-finance/treasury/blob/master/assets.md

import (
	"math/big"
	"net/http"
	"strconv"
	"yam-api/source/config"
	IndexStakingRewards "yam-api/source/contracts/indexStakingRewards"
	Uni "yam-api/source/contracts/uni"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"
	"yam-api/source/utils/log"
	"yam-api/source/utils/mongodb"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

type TreasuryAsset struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Quantity float64 `json:"quantity"`
	Value    float64 `json:"value"`
	Price    float64 `json:"price"`
	Change   string  `json:"change"`
}

func Treasury(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		// response := GetTreasury(geth)
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
	assets := map[string]interface{}{}
	noChange := float64(0)
	wbtcBalance := utils.GetBalance(contractAddress.WBTC, contractAddress.TreasuryAddress, geth, 8)
	wbtcPrice := utils.GetValue("wrapped-bitcoin")
	wbtcValue := CalculateValue(wbtcBalance, wbtcPrice)
	wethBalanceTreasury := utils.GetBalance(contractAddress.WETH, contractAddress.TreasuryAddress, geth, 18)
	wethBalanceMultisig := utils.GetBalance(contractAddress.WETH, contractAddress.MultisigAddress, geth, 18)
	wethBalance := new(big.Float).Add(wethBalanceTreasury, wethBalanceMultisig)
	wethPrice := utils.GetValue("ethereum")
	wethValue := CalculateValue(wethBalance, wethPrice)
	yamBalanceTreasury := utils.GetBalance(contractAddress.YAM, contractAddress.TreasuryAddress, geth, 18)
	yamBalanceMultisig := utils.GetBalance(contractAddress.YAM, contractAddress.MultisigAddress, geth, 18)
	yamBalance := new(big.Float).Add(yamBalanceTreasury, yamBalanceMultisig)
	yamPrice := utils.GetValue("yam-2")
	yamValue := CalculateValue(yamBalance, yamPrice)
	dpiBalance := utils.GetBalance(contractAddress.DPI, contractAddress.TreasuryAddress, geth, 18)
	dpiPrice := utils.GetValue("defipulse-index")
	dpiValue := CalculateValue(dpiBalance, dpiPrice)
	umaBalance := utils.GetBalance(contractAddress.UMA, contractAddress.TreasuryAddress, geth, 18)
	umaPrice := utils.GetValue("uma")
	umaValue := CalculateValue(umaBalance, umaPrice)
	sushiBalanceTreasury := utils.GetBalance(contractAddress.SUSHI, contractAddress.TreasuryAddress, geth, 18)
	sushiBalanceIncentivizer := utils.GetBalance(contractAddress.SUSHI, contractAddress.ContractIncentivizer, geth, 18)
	sushiBalance := new(big.Float).Add(sushiBalanceTreasury, sushiBalanceIncentivizer)
	sushiPrice := utils.GetValue("sushi")
	sushiValue := CalculateValue(sushiBalance, sushiPrice)
	xsushiBalance := utils.GetBalance(contractAddress.XSUSHI, contractAddress.TreasuryAddress, geth, 18)
	xsushiPrice := utils.GetValue("xsushi")
	xsushiValue := CalculateValue(xsushiBalance, xsushiPrice)
	indexBalanceTreasury := utils.GetBalance(contractAddress.INDEX, contractAddress.TreasuryAddress, geth, 18)
	indexBalanceTimelock := utils.GetBalance(contractAddress.INDEX, contractAddress.ContractTimelock, geth, 18)
	indexBalance := new(big.Float).Add(indexBalanceTreasury, indexBalanceTimelock)
	indexPrice := utils.GetValue("index-cooperative")
	indexValue := CalculateValue(indexBalance, indexPrice)
	usdcBalanceTreasury := utils.GetBalance(contractAddress.USDC, contractAddress.TreasuryAddress, geth, 6)
	usdcBalanceMultisig := utils.GetBalance(contractAddress.USDC, contractAddress.MultisigAddress, geth, 6)
	usdcBalance := new(big.Float).Add(usdcBalanceTreasury, usdcBalanceMultisig)
	usdcPrice := utils.GetValue("usd-coin")
	usdcValue := CalculateValue(usdcBalance, usdcPrice)
	gtcBalance := utils.GetBalance(contractAddress.GitCoinAddress, contractAddress.TreasuryAddress, geth, 18)
	gtcPrice := utils.GetPriceByContract(contractAddress.GitCoinAddress)
	gtcValue := CalculateValue(gtcBalance, gtcPrice)
	yvusdcBalance := utils.GetBalance(contractAddress.YVUSDC, contractAddress.TreasuryAddress, geth, 6)
	yvusdcPrice := new(big.Float).SetFloat64(1) // update new
	yvusdcValue := CalculateValue(yvusdcBalance, yvusdcPrice)
	stethBalance := utils.GetBalance(contractAddress.STETH, contractAddress.TreasuryAddress, geth, 18)
	stethPrice := utils.GetValue("staked-ether")
	stethValue := CalculateValue(stethBalance, stethPrice)
	ethdpilpBalance := GetETHDPILPBalance(geth)
	ethdpilpPrice := GetETHDPILPPrice(geth)
	ethdpilpValue := CalculateValue(ethdpilpBalance, ethdpilpPrice)

	assets["wbtc"] = SetTreasuryAssetInfo("Wrapped Bitcoin", "WBTC", wbtcBalance, wbtcPrice, wbtcValue, utils.GetValueChange("weth"))
	assets["weth"] = SetTreasuryAssetInfo("Wrapped Ether", "WETH", wethBalance, wethPrice, wethValue, utils.GetValueChange("weth"))
	assets["yam"] = SetTreasuryAssetInfo("Yam Token", "YAM", yamBalance, yamPrice, yamValue, utils.GetValueChange("yam-2"))
	assets["dpi"] = SetTreasuryAssetInfo("DefiPulse Index", "DPI", dpiBalance, dpiPrice, dpiValue, utils.GetValueChange("defipulse-index"))
	assets["uma"] = SetTreasuryAssetInfo("Uma Token", "UMA", umaBalance, umaPrice, umaValue, utils.GetValueChange("uma"))
	assets["sushi"] = SetTreasuryAssetInfo("Sushi Token", "SUSHI", sushiBalance, sushiPrice, sushiValue, utils.GetValueChange("sushi"))
	assets["xsushi"] = SetTreasuryAssetInfo("XSushi Token", "XSUSHI", xsushiBalance, xsushiPrice, xsushiValue, utils.GetValueChange("xsushi"))
	assets["index"] = SetTreasuryAssetInfo("Index Token", "INDEX", indexBalance, indexPrice, indexValue, utils.GetValueChange("index-cooperative"))
	assets["usdc"] = SetTreasuryAssetInfo("USD Coin", "USDC", usdcBalance, usdcPrice, usdcValue, utils.GetValueChange("usd-coin"))
	assets["gtc"] = SetTreasuryAssetInfo("Gitcoin Token", "GTC", gtcBalance, gtcPrice, gtcValue, utils.GetValueChange("gitcoin"))
	assets["yvusdc"] = SetTreasuryAssetInfo("yvVault Token", "yvUSDC", yvusdcBalance, yvusdcPrice, yvusdcValue, utils.GetValueChange("usd-coin"))
	assets["steth"] = SetTreasuryAssetInfo("Curve stETH Token", "stETH", stethBalance, stethPrice, stethValue, utils.GetValueChange("staked-ether"))
	assets["ethdpilp"] = SetTreasuryAssetInfo("ETH/DPI LP", "ETHDPILP", ethdpilpBalance, ethdpilpPrice, ethdpilpValue, noChange)
	response["assets"] = assets
	response["total"] = utils.FixedDecimals(wbtcValue + wethValue + yamValue + dpiValue + umaValue + sushiValue + xsushiValue + indexValue + usdcValue + gtcValue + yvusdcValue + stethValue + ethdpilpValue)
	return response
}

func StoreTreasury(val map[string]interface{}) {
	mongodb.InsertTreasury(val)
}

func CalculateValue(balance *big.Float, price *big.Float) float64 {
	return utils.FixedTwoDecimals(new(big.Float).Mul(balance, price).Float64())
}

func SetTreasuryAssetInfo(name string, symbol string, balance *big.Float, price *big.Float, value float64, change float64) TreasuryAsset {
	var assetInfo TreasuryAsset
	assetInfo.Name = name
	assetInfo.Symbol = symbol
	assetInfo.Price = utils.FixedTwoDecimals(price.Float64())
	assetInfo.Quantity = utils.FixedTwoDecimals(balance.Float64())
	assetInfo.Value = value
	assetInfo.Change = strconv.FormatFloat(change, 'f', -1, 64)
	return assetInfo
}

func GetETHDPILPBalance(geth *ethclient.Client) *big.Float {
	IndexStakingRewardsContract, err := IndexStakingRewards.NewIndexStakingRewards(common.HexToAddress(contractAddress.IndexStakingRewards), geth)
	if err != nil {
		log.Error(err)
	}
	lpBalance, err := IndexStakingRewardsContract.BalanceOf(nil, common.HexToAddress(contractAddress.ContractIndexStaking))
	if err != nil {
		log.Error("failed to get lpBalance: %v", err)
	}
	result := utils.BnToDec(lpBalance, 18)
	return result
}

func GetETHDPILPPrice(geth *ethclient.Client) *big.Float {
	unilpContract, err := Uni.NewUni(common.HexToAddress(contractAddress.UnilpETHDPI), geth)
	if err != nil {
		log.Error(err)
	}
	lpTotalSupply, err := unilpContract.TotalSupply(nil)
	if err != nil {
		log.Error("failed to get lpTotalSupply: %v", err)
	}
	lpTotalSupplyValue := utils.BnToDec(lpTotalSupply, 18)
	wethBalance := utils.GetBalance(contractAddress.WETH, contractAddress.UnilpETHDPI, geth, 18)
	dpiBalance := utils.GetBalance(contractAddress.DPI, contractAddress.UnilpETHDPI, geth, 18)
	wethLPValue, _ := new(big.Float).Mul(wethBalance, utils.GetWETHPrice()).Float64()
	dpiLPValue, _ := new(big.Float).Mul(dpiBalance, utils.GetValue("defipulse-index")).Float64()
	pairPriceValueFloat := strconv.FormatFloat(utils.FixedDecimals(wethLPValue+dpiLPValue), 'f', -1, 64)
	pairPriceValue, _ := strconv.ParseFloat(pairPriceValueFloat, 64)
	lpValue, _ := new(big.Float).Quo(new(big.Float).SetFloat64(pairPriceValue), lpTotalSupplyValue).Float64()
	return new(big.Float).SetFloat64(lpValue)
}

// func GetSushiBalance(geth *ethclient.Client) *big.Float {
// 	sushiContract, err := sushi.NewSushi(common.HexToAddress(contractAddress.SushiToken), geth)
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	sushibarContract, err := sushibar.NewSushibar(common.HexToAddress(contractAddress.SushibarXSushi), geth)
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	incentivizerBalance, err := sushibarContract.BalanceOf(nil, common.HexToAddress(contractAddress.ContractIncentivizer))
// 	if err != nil {
// 		log.Error("failed to get incentivizerBalance: %v", err)
// 	}
// 	xSushiTotalSupply, err := sushibarContract.TotalSupply(nil)
// 	if err != nil {
// 		log.Error("failed to get xSushiTotalSupply: %v", err)
// 	}
// 	sushiXSushiBalance, err := sushiContract.BalanceOf(nil, common.HexToAddress(contractAddress.SushibarXSushi))
// 	if err != nil {
// 		log.Error("failed to get sushiXSushiBalance: %v", err)
// 	}
// 	SushiBalance, err := sushiContract.BalanceOf(nil, common.HexToAddress(contractAddress.ContractIncentivizer))
// 	if err != nil {
// 		log.Error("failed to get SushiBalance: %v", err)
// 	}
// 	incentivizerBalanceF := utils.BnToDec(incentivizerBalance, 18)
// 	xSushiTotalSupplyF := utils.BnToDec(xSushiTotalSupply, 18)
// 	sushiXSushiBalanceF := utils.BnToDec(sushiXSushiBalance, 18)
// 	SushiBalanceF := utils.BnToDec(SushiBalance, 18)
// 	ret := new(big.Float).Add(SushiBalanceF, new(big.Float).Quo(new(big.Float).Mul(incentivizerBalanceF, sushiXSushiBalanceF), xSushiTotalSupplyF))
// 	return ret
// }

func marketCap() {}

func treasuryValue() {}

func totalSupply() {}

func historyReserves() {}
