package utils

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	erc20 "yam-api/source/contracts/erc20"
	IndexStakingRewards "yam-api/source/contracts/indexStakingRewards"
	sushi "yam-api/source/contracts/sushi"
	sushibar "yam-api/source/contracts/sushibar"
	"yam-api/source/utils/contractAddress"
	"yam-api/source/utils/etherscan/response"
	"yam-api/source/utils/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type empList struct {
	EmpWhiteList []string
	TotalReward  float64
}

func ResError(code int, w http.ResponseWriter, msg string) {
	ResJSON(code, w, map[string]interface{}{"code": code, "message": msg})
}

func ResByte(c int, h http.ResponseWriter, payload []byte) {
	h.Header().Set("Content-Type", "application/json")
	h.WriteHeader(c)
	h.Write(payload)
}

func ResJSON(c int, h http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)
	h.Header().Set("Content-Type", "application/json")
	h.WriteHeader(c)
	h.Write(response)
}

func MakeRequest(url string) ([]byte, error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		log.Error(respErr)
		return nil, respErr
	}
	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		log.Error(bodyErr)
		return nil, bodyErr
	}
	return []byte(body), nil
}

func BuildUrl(r *http.Request, u string) string {
	url := u
	from := r.URL.Query().Get("from")
	if from != "" {
		url = url + "?from=" + from
	}
	to := r.URL.Query().Get("to")
	if to != "" {
		url = url + "?to=" + to
	}
	page := r.URL.Query().Get("page")
	if page != "" {
		url = url + "?page=" + page
	}
	log.Info("url", url)
	return url
}

func CheckSuccess(code int) bool {
	value := true
	switch code {
	case 200:
		value = true
	case 201:
		value = true
	case 429:
		value = false
	case 451:
		value = false
	}
	return value
}

func BnToDec(number *big.Int, decimal int) *big.Float {
	divineNumber := math.Pow10(decimal)

	x, y := new(big.Float).SetInt(number), big.NewFloat(divineNumber)

	z := new(big.Float).Quo(x, y)

	return z
}

func FilterArray(arr []response.Tx, cond func(response.Tx) bool) []response.Tx {
	result := []response.Tx{}

	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}

	return result
}

func GetWETHPrice() *big.Float {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/weth")
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	sb := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(sb), &result)

	market_data := result["market_data"].(map[string]interface{})
	current_price := market_data["current_price"].(map[string]interface{})
	return big.NewFloat(current_price["usd"].(float64))
}
func GetPriceByContract(address string) *big.Float {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/ethereum/contract/" + address)
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	sb := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(sb), &result)

	market_data := result["market_data"].(map[string]interface{})
	current_price := market_data["current_price"].(map[string]interface{})
	return big.NewFloat(current_price["usd"].(float64))

}

func IndexOf(params ...interface{}) int {
	v := reflect.ValueOf(params[0])
	arr := reflect.ValueOf(params[1])

	var t = reflect.TypeOf(params[1]).Kind()

	if t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}
	return -1
}

func MergeUnique(a []string, b []string) []string {
	var ret []string = a

	for _, item := range b {
		if IndexOf(item, a) == -1 {
			ret = append(ret, item)
		}
	}
	//fmt.Println(ret)
	return ret

}
func GetDevMiningEmps() empList {
	resp, err := http.Get("https://raw.githubusercontent.com/UMAprotocol/protocol/master/packages/affiliates/payouts/devmining-status.json")
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	sb := string(body)
	var result empList
	json.Unmarshal([]byte(sb), &result)
	for i, item := range result.EmpWhiteList {
		item := strings.ToLower(item)
		result.EmpWhiteList[i] = item
	}
	return result
}
func GetBalance(tokenAddress string, userAddress string, geth *ethclient.Client, decimals int) *big.Float {
	tokenContract := GetErc20Contract(tokenAddress, geth)
	balance, err := tokenContract.BalanceOf(nil, common.HexToAddress(userAddress))
	if err != nil {
		log.Error(err)
	}
	return BnToDec(balance, decimals)
}
func GetErc20Contract(address string, geth *ethclient.Client) *erc20.Erc20 {
	erc20Contract, err := erc20.NewErc20(common.HexToAddress(address), geth)
	if err != nil {
		log.Error(err)
	}
	return erc20Contract
}
func GetValue(asset_name string) *big.Float {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + asset_name)
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	sb := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(sb), &result)

	market_data := result["market_data"].(map[string]interface{})
	current_price := market_data["current_price"].(map[string]interface{})
	return big.NewFloat(current_price["usd"].(float64))
}
func GetIndexCoopLPRewards(geth *ethclient.Client) *big.Float {
	IndexStakingRewardsContract, err := IndexStakingRewards.NewIndexStakingRewards(common.HexToAddress(contractAddress.IndexStakingRewards), geth)
	if err != nil {
		log.Error(err)
	}
	lpBalanceRewards, err := IndexStakingRewardsContract.Earned(nil, common.HexToAddress(contractAddress.ContractIndexStaking))
	if err != nil {
		log.Error("failed to get lpBalanceRewards: %v", err)
	}
	return BnToDec(lpBalanceRewards, 18)

}
func GetSushiRewards(geth *ethclient.Client) *big.Float {
	sushiContract, err := sushi.NewSushi(common.HexToAddress(contractAddress.SushiToken), geth)
	if err != nil {
		log.Error(err)
	}
	sushibarContract, err := sushibar.NewSushibar(common.HexToAddress(contractAddress.SushibarXSushi), geth)
	if err != nil {
		log.Error(err)
	}
	incentivizerBalance, err := sushibarContract.BalanceOf(nil, common.HexToAddress(contractAddress.ContractIncentivizer))
	if err != nil {
		log.Error("failed to get incentivizerBalance: %v", err)
	}
	xSushiTotalSupply, err := sushibarContract.TotalSupply(nil)
	if err != nil {
		log.Error("failed to get xSushiTotalSupply: %v", err)
	}
	sushiXSushiBalance, err := sushiContract.BalanceOf(nil, common.HexToAddress(contractAddress.SushibarXSushi))
	if err != nil {
		log.Error("failed to get sushiXSushiBalance: %v", err)
	}
	SushiBalance, err := sushiContract.BalanceOf(nil, common.HexToAddress(contractAddress.ContractIncentivizer))
	if err != nil {
		log.Error("failed to get SushiBalance: %v", err)
	}
	incentivizerBalanceF := BnToDec(incentivizerBalance, 18)
	xSushiTotalSupplyF := BnToDec(xSushiTotalSupply, 18)
	sushiXSushiBalanceF := BnToDec(sushiXSushiBalance, 18)
	SushiBalanceF := BnToDec(SushiBalance, 18)
	ret := new(big.Float).Add(SushiBalanceF, new(big.Float).Quo(new(big.Float).Mul(incentivizerBalanceF, sushiXSushiBalanceF), xSushiTotalSupplyF))
	return ret
}
func GetYamHousePrice() *big.Float {
	resp, err := http.Get("https://api.tokensets.com/public/v2/portfolios/yamhouse")
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	sb := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(sb), &result)

	portfolio := result["portfolio"].(map[string]interface{})

	s, err := strconv.ParseFloat(portfolio["price_usd"].(string), 64)
	return big.NewFloat(s)
}
func FixedTwoDecimal(val float64) float64 {
	return math.Floor(val*100) / 100
}

func GetValueChange(asset_name string) *big.Float {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + asset_name)
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	sb := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(sb), &result)

	market_data := result["market_data"].(map[string]interface{})
	price_change := market_data["price_change_percentage_24h_in_currency"].(map[string]interface{})

	return big.NewFloat(price_change["usd"].(float64))
}
func CopyMap(assetInfo map[string]interface{}) map[string]interface{} {
	targetMap := map[string]interface{}{}
	for key, value := range assetInfo {
		targetMap[key] = value
	}
	return targetMap
}
func SetTreasuryAssetInfo(quantity *big.Float, price *big.Float, change *big.Float) map[string]interface{} {
	var val float64
	val, _ = quantity.Float64()
	assetInfo := map[string]interface{}{}
	assetInfo["quantity"] = FixedTwoDecimal(val)
	val, _ = price.Float64()
	assetInfo["price"] = FixedTwoDecimal(val)
	val, _ = change.Float64()
	assetInfo["change"] = FixedTwoDecimal(val)
	return assetInfo
}
