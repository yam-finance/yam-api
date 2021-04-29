package routes

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httputil"
	"net/url"
	"time"
	"math/big"
	"strconv"

	"yam-api/source/config"
	"yam-api/source/utils"
	"yam-api/source/utils/log"
	"yam-api/source/utils/etherscan/reflect"
	"yam-api/source/utils/etherscan/helper"
	"yam-api/source/utils/etherscan/response"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

/// -------- Types --------

// @dev Ethereum network type
type Network string

// @dev API client struct
type Client struct {
	coon    *http.Client
	key     string
	baseURL string

	// @dev True to get more log information.
	Verbose bool

	// @dev Runs before every client request, usually to check the rate limit.
	BeforeRequest func(module, action string, param map[string]interface{}) error

	// @dev Runs after every client request, even when there is an error.
	AfterRequest func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

// @dev Customization for NewCustomized function.
type Customization struct {
	// @dev Timeout for API call.
	Timeout time.Duration
	// @dev API key
	Key string
	// @dev Base URL
	BaseURL string
	// @dev True to get more log information.
	Verbose bool

	// @dev Runs before every client request, usually to check the rate limit.
	BeforeRequest func(module, action string, param map[string]interface{}) error

	// @dev Runs after every client request, even when there is an error.
	AfterRequest func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

// @dev Bucket is a simple rate limiter
type Bucket struct {
	bucket     chan bool
	refillTime time.Duration
}

/// -------- Variables --------

var (
	// @dev API test client
	api *Client
	// @dev Bucket rate limiter
	bucket *Bucket
	// @dev Etherscan API key
	apiKey string
)

const (
	Mainnet Network = "api"
	Ropsten Network = "api-ropsten"
	Kovan Network = "api-kovan"
	Rinkeby Network = "api-rinkeby"
)

/// -------- Base functions --------

// @dev Returns the subdomain of etherscan API
func (n Network) SubDomain() (sub string)  {
	return string(n)
}

// @dev Initialize a new etherscan API client
func New(network Network, APIKey string) *Client {
	return NewCustomized(Customization {
		Timeout: 30 * time.Second,
		Key:     APIKey,
		BaseURL: fmt.Sprintf(`https://%s.etherscan.io/api?`, network.SubDomain()),
	})
}

// @dev Initialize a new customized API client, usually to call other API's.
func NewCustomized(config Customization) *Client {
	return &Client {
		coon: &http.Client {
			Timeout: config.Timeout,
		},
		key:           config.Key,
		baseURL:       config.BaseURL,
		Verbose:       config.Verbose,
		BeforeRequest: config.BeforeRequest,
		AfterRequest:  config.AfterRequest,
	}
}

// @dev Function that takes care of the actual API call.
func (c *Client) call(module, action string, param map[string]interface{}, outcome interface{}) (err error) {
	// @dev Call hooks if declared.
	if c.BeforeRequest != nil {
		err = c.BeforeRequest(module, action, param)
		if err != nil {
			err = log.WrapErr(err, "beforeRequest")
			return
		}
	}

	if c.AfterRequest != nil {
		defer c.AfterRequest(module, action, param, outcome, err)
	}

	// @dev Recover if panic was thrown.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v", r)
		}
	}()

	req, err := http.NewRequest(http.MethodGet, c.craftURL(module, action, param), http.NoBody)

	if err != nil {
		err = log.WrapErr(err, "http.NewRequest")
		return
	}

	req.Header.Set("User-Agent", "etherscan-api(Go)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.Verbose {
		var reqDump []byte
		reqDump, err = httputil.DumpRequestOut(req, false)

		if err != nil {
			err = log.WrapErr(err, "verbose mode req dump failed")
			return
		}

		fmt.Printf("\n%s\n", reqDump)

		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.coon.Do(req)

	if err != nil {
		err = log.WrapErr(err, "sending request")
		return
	}

	defer res.Body.Close()

	if c.Verbose {
		var resDump []byte
		resDump, err = httputil.DumpResponse(res, true)

		if err != nil {
			err = log.WrapErr(err, "verbose mode res dump failed")
			return
		}

		fmt.Printf("%s\n", resDump)
	}

	var content bytes.Buffer

	if _, err = io.Copy(&content, res.Body); err != nil {
		err = log.WrapErr(err, "reading response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", res.StatusCode, res.Status, content.String())
		return
	}

	var envelope response.Envelope
	err = json.Unmarshal(content.Bytes(), &envelope)

	if err != nil {
		err = log.WrapErr(err, "json unmarshal envelope")
		return
	}
	if envelope.Status != 1 {
		err = fmt.Errorf("etherscan server: %s", envelope.Message)
		return
	}

	// @dev Workaround for missing tokenDecimal for some token transfer calls.
	if action == "tokentx" {
		err = json.Unmarshal(bytes.Replace(envelope.Result, []byte(`"tokenDecimal":""`), []byte(`"tokenDecimal":"0"`), -1), outcome)
	} else {
		err = json.Unmarshal(envelope.Result, outcome)
	}

	if err != nil {
		err = log.WrapErr(err, "json unmarshal outcome")
		return
	}

	return
}

// @dev Returns URL built with provided params.
func (c *Client) craftURL(module, action string, param map[string]interface{}) (URL string) {
	q := url.Values{
		"module": []string{module},
		"action": []string{action},
		"apikey": []string{c.key},
	}

	for k, v := range param {
		q[k] = reflect.ExtractValue(v)
	}

	URL = c.baseURL + q.Encode()

	return
}

/// -------- Call Setup --------

// @dev Bucket factory
func NewBucket(refillTime time.Duration) (b *Bucket) {
	b = &Bucket {
		bucket:     make(chan bool),
		refillTime: refillTime,
	}

	go b.fillRoutine()

	return
}

// @dev Blocks if there is currently no action token left.
func (b *Bucket) Take() {
	<-b.bucket
}

// @dev Fill an action token into bucket.
func (b *Bucket) fill() {
	b.bucket <- true
}

func (b *Bucket) fillRoutine() {
	ticker := time.NewTicker(b.refillTime)

	for range ticker.C {
		b.fill()
	}
}

func (c *Client) setup(apiKey string) {
	var envApiKey = apiKey
	apiKey = envApiKey

	if apiKey == "" {
		panic(fmt.Sprintf("API key is empty, set env variable %q with a valid API key to proceed.", envApiKey))
	}

	bucket = NewBucket(500 * time.Millisecond)

	api = New(Mainnet, apiKey)
	api.Verbose = true
	api.BeforeRequest = func(module string, action string, param map[string]interface{}) error {
		bucket.Take()
		return nil
	}
}

/// -------- Calls --------

// @dev Get a list of "normal" transactions by address
func (c *Client) NormalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []response.Tx, err error) {
	param := helper.M {
		"address": address,
		"page":    page,
		"offset":  offset,
	}

	helper.Compose(param, "startblock", startBlock)
	helper.Compose(param, "endblock", endBlock)

	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "txlist", param, &txs)

	return
}

// @dev Get a list of "erc20 - token transfer events" by address
func (c *Client) ERC20Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []response.Tx, err error) {
	param := helper.M {
		"page":   page,
		"offset": offset,
	}

	helper.Compose(param, "contractaddress", contractAddress)
	helper.Compose(param, "address", address)
	helper.Compose(param, "startblock", startBlock)
	helper.Compose(param, "endblock", endBlock)

	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "tokentx", param, &txs)

	return
}

func (c *Client) GetBlockNumberByTimestamp(timestamp int, closest string) (blockNumber string, err error) {
	param := helper.M {
		"timestamp": timestamp,
		"closest": closest,
	}

	err = c.call("block", "getblocknobytime", param, &blockNumber)

	return
}

func GetNormalTxByAddress(accountAddress string, startBlock int, endBlock int) (txs []response.Tx, err error) {
	var a, b = startBlock, endBlock
	txs, err = api.NormalTxByAddress(accountAddress, &a, &b, 1, 500, false)

	return
}

func GetERC20Transfers(accountAddress string, startBlock int, endBlock int) (txs []response.Tx, err error) {
	var a, b = startBlock, endBlock
	var address = accountAddress
	txs, err = api.ERC20Transfers(nil, &address, &a, &b, 1, 500, false)

	return
}

func GasStats(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		api.setup(conf.EtherscanKey)

		// @todo Update with passed params.
		accountAddress := "0x4e83362442b8d1bec281594cea3050c8eb01311c"
		startTimeStamp := 1578638524
		endTimeStamp := 1619373243

		// @dev Build startBlock
		block, blockNumberErr := api.GetBlockNumberByTimestamp(startTimeStamp, "after")
		startBlock, _ := strconv.Atoi(block)

		// @dev Build endBlock
		block, blockNumberErr = api.GetBlockNumberByTimestamp(endTimeStamp, "before")
		endBlock, _ := strconv.Atoi(block)

		normalTxs, normalTxsErr := GetNormalTxByAddress(accountAddress, startBlock, endBlock)
		erc20Txs, erc20TxsErr := GetERC20Transfers(accountAddress, startBlock, endBlock)
		txs := normalTxs
		txs = append(txs, erc20Txs...)

		txsFromAddress := utils.FilterArray(txs, func(val response.Tx) bool {
			return val.From != accountAddress
		})

		failedTxsFromAddress := utils.FilterArray(txsFromAddress, func(val response.Tx) bool {
			return val.IsError != 0 // 0 = No Error, 1 = Got Error.
		})
		
		// Success Calculation
		txCount := len(txsFromAddress)
		var gasUsedArray []int
		var gasPriceArray []*helper.BigInt
		var gasFeeArray []*big.Int
		totalGasFee := new(big.Int).SetInt64(0)
		var totalGasFeeETH *big.Float

		// Failed Calculation
		failedTxCount := len(failedTxsFromAddress)
		var failedGasUsedArray []int
		var failedGasPriceArray []*helper.BigInt
		var failedGasFeeArray []*big.Int
		totalFailedGasFee := new(big.Int).SetInt64(0)
		var totalFailedGasFeeETH *big.Float

		// Average Calculation
		totalGasPrice := new(big.Int).SetInt64(0)
		var averageTxPrice *big.Int
		var averageTxPriceGWEI *big.Float

		if (txCount > 0) {
			// Success Calculation

			for i := 0; i < txCount; i++ {
				gasUsedArray = append(gasUsedArray, txsFromAddress[i].GasUsed)
			}
		
			for i := 0; i < txCount; i++ {
				gasPriceArray = append(gasPriceArray, txsFromAddress[i].GasPrice)
			}

			for i := 0; i < len(gasUsedArray); i++ {
				var gasUsed = new(big.Int).SetInt64(int64(gasUsedArray[i]))
				gasFeeArray = append(gasFeeArray, new(big.Int).Mul(gasUsed, gasPriceArray[i].Int()))
			}

			for i := 0; i < len(gasFeeArray); i++ {
				totalGasFee.Add(totalGasFee, gasFeeArray[i])  
			}

			totalGasFeeETH = utils.BnToDec(totalGasFee, 18)

			// Failed Calculation

			for i := 0; i < failedTxCount; i++ {
				failedGasUsedArray = append(failedGasUsedArray, failedTxsFromAddress[i].GasUsed)
			}
		
			for i := 0; i < failedTxCount; i++ {
				failedGasPriceArray = append(failedGasPriceArray, failedTxsFromAddress[i].GasPrice)
			}

			for i := 0; i < len(failedGasUsedArray); i++ {
				var failedGasUsed = new(big.Int).SetInt64(int64(failedGasUsedArray[i]))
				failedGasFeeArray = append(failedGasFeeArray, new(big.Int).Mul(failedGasUsed, failedGasPriceArray[i].Int()))
			}

			for i := 0; i < len(failedGasFeeArray); i++ {
				totalFailedGasFee.Add(totalFailedGasFee, failedGasFeeArray[i])  
			}

			totalFailedGasFeeETH = utils.BnToDec(totalFailedGasFee, 18)

			// Average Calculation

			for i := 0; i < len(gasPriceArray); i++ {
				totalGasPrice.Add(totalGasPrice, gasPriceArray[i].Int())  
			}

			averageTxPrice = totalGasPrice.Div(totalGasPrice, new(big.Int).SetInt64(int64(txCount)))
			averageTxPriceGWEI = utils.BnToDec(averageTxPrice, 9)
		}

		// @dev Log error messages.
		if normalTxsErr != nil {
			log.Error(normalTxsErr)
		} else if erc20TxsErr != nil {
			log.Error(erc20TxsErr)
		} else if blockNumberErr != nil {
			log.Error(blockNumberErr)
		}

		utils.ResJSON(http.StatusCreated, w,
			map[string] interface{} {
				"txCount": txCount,
				"failedTxCount": failedTxCount,
				"totalGasFeeETH": string(totalGasFeeETH.String()),
				"totalFailedGasFeeETH": string(totalFailedGasFeeETH.String()),
				"averageTxPriceGWEI": averageTxPriceGWEI.String(),
			},
		)
	})
}
