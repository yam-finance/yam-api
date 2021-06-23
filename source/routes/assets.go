package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"

	"github.com/Jeffail/gabs"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

func GetAssetsJson(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get("https://raw.githubusercontent.com/yam-finance/degenerative-sdk/master/src/assets.json")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		responseData, _ := ioutil.ReadAll(resp.Body)

		responseObject, err := gabs.ParseJSON(responseData)
		if err != nil {
			log.Fatalln(err)
		}

		utils.ResJSON(http.StatusCreated, w,
			responseObject.Data(),
		)
	})
}
