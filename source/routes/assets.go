package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

type Response struct {
	Mainnet Synth `json:"mainnet"`
	Kovan Synth `json:"kovan"`
}

type Synth struct {
	Ugas Content `json:"ugas"`
	Ustonks Content `json:"ustonks"`
	Upunks Content `json:"upunks"`
	Uvolbtc Content `json:"uvolbtc"`
	Yamdollar Content `json:"yamdollar"`
}

type Content []struct {
	Name       string `json:"name"`
	Cycle      string `json:"cycle"`
	Year       string `json:"year"`
	Collateral string `json:"collateral"`
	Token      struct {
		Address  string `json:"address"`
		Decimals int    `json:"decimals"`
	} `json:"token"`
	Emp struct {
		Address string `json:"address"`
		New     bool   `json:"new"`
	} `json:"emp"`
	Pool struct {
		Address  string `json:"address"`
		Location string `json:"location"`
	} `json:"pool"`
	Expired bool `json:"expired"`
}

func GetAssetsJson(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get("https://raw.githubusercontent.com/yam-finance/degenerative-sdk/master/src/assets.json")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		responseData, _ := ioutil.ReadAll(resp.Body)

		var responseObject Response
		json.Unmarshal(responseData, &responseObject)

		utils.ResJSON(http.StatusCreated, w,
			responseObject,
		)
	})
}
