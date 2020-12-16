package routes

import (
	"context"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"
	"yam-api/source/utils/log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

func Block(path string, router chi.Router, conf *config.Config, geth *ethclient.Client) {

	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		header, errHeader := geth.HeaderByNumber(context.Background(), nil)
		if errHeader != nil {
			log.Error(errHeader)
		}

		if !utils.CheckSuccess(http.StatusCreated) {
			utils.ResError(http.StatusCreated, w,
				"error",
			)
		} else {
			utils.ResJSON(http.StatusCreated, w,
				map[string]interface{}{
					"code":    http.StatusCreated,
					"message": header.Number,
				},
			)
		}
	})

}
