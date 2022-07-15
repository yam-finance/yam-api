package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"yam-api/source/config"
	"yam-api/source/utils"
	"yam-api/source/utils/contractAddress"
	"yam-api/source/utils/log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

type PriceDataResp struct {
	MarketData struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}

type MarketDataResp struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

type PriceAvgResp struct {
	From  string  `json:"from"`
	To    string  `json:"to"`
	Price float64 `json:"price"`
}

func Price(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		yamPrice := utils.GetPriceByContract(contractAddress.YAM)
		if !utils.CheckSuccess(http.StatusCreated) {
			utils.ResError(http.StatusCreated, w,
				"error",
			)
		} else {
			utils.ResJSON(http.StatusCreated, w,
				yamPrice,
			)
		}
	})

}

func PriceAvg(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {

	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		const base = "2006-01-02"
		const baseOut = "Monday 2, Jan 2006"
		startDate := chi.URLParam(r, "startDate")
		endDate := chi.URLParam(r, "endDate")
		startDateTS, _ := time.Parse(base, startDate)
		endDateTS, _ := time.Parse(base, endDate)
		startDateUnix := startDateTS.Unix()
		endDateUnix := endDateTS.Unix()
		marketResponse, err := http.Get("https://api.coingecko.com/api/v3/coins/yam-2/market_chart/range?vs_currency=usd&from=" + fmt.Sprint(startDateUnix) + "&to=" + fmt.Sprint(endDateUnix))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		var marketData MarketDataResp
		marketResponseBody, err := ioutil.ReadAll(marketResponse.Body)
		if err != nil {
			log.Error(err)
		}
		json.Unmarshal([]byte(marketResponseBody), &marketData)
		total := 0.0
		for _, row := range marketData.Prices {
			for _, val := range row[1:] {
				total = total + val
			}
		}
		priceAvg := total / float64(len(marketData.Prices))

		if !utils.CheckSuccess(http.StatusCreated) {
			utils.ResError(http.StatusCreated, w,
				"error",
			)
		} else {
			utils.ResJSON(http.StatusCreated, w,
				PriceAvgResp{
					From:  startDateTS.Format(baseOut),
					To:    endDateTS.Format(baseOut),
					Price: priceAvg,
				},
			)
		}
	})

}
