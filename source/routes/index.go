package routes

import (
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator"
)

var validate *validator.Validate
var AssetsFromFile AssetsFile

func Initialize(conf *config.Config, geth *ethclient.Client) chi.Router {

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler)

	// /
	Index("/", router, conf)
	Version("/version", router, conf)
	Block("/block", router, conf, geth)

	// YAM
	Treasury("/treasury", router, conf, geth)
	Apr("/apr", router, conf, geth)
	AprYam("/apr/yam", router, conf, geth)
	AprDegenerative("/apr/degenerative", router, conf, geth)

	// Degenerativde
	GetAssets("/degenerative/assets", router, conf, geth)
	GetLatestPunkIndex("/degenerative/upunks/price", router, conf, geth)
	GetPunkIndexHistory("/degenerative/upunks/price-history", router, conf, geth)
	AprDegenerative("/degenerative/apr", router, conf, geth)
	GetAssetIndex("/degenerative/ustonks/index-history-daily", router, conf, geth)

	// Other
	GasStats("/account-stats", router, conf, geth)

	return router
}

func Index(path string, router chi.Router, conf *config.Config) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		utils.ResJSON(http.StatusCreated, w,
			"api.yam.finance",
		)
	})
}
