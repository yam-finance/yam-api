package utils

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"reflect"
	"strings"
	"yam-api/source/utils/etherscan/response"
	"yam-api/source/utils/log"
)


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
	//f, err := strconv.ParseFloat(current_price["usd"], 8)
	//fmt.Println(current_price["usd"])

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
	//f, err := strconv.ParseFloat(current_price["usd"], 8)
	//fmt.Println(current_price["usd"])

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

type empList struct {
	EmpWhiteList []string
	TotalReward  float64
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
	//fmt.Println(result.TotalReward)
	return result
}
