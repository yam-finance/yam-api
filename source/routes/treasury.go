package routes

import (
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

func Treasury(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		utils.ResJSON(http.StatusCreated, w,
			"treasury v"+conf.Version,
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
