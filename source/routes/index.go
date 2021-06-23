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

	// General
	Index("/", router, conf)
	Version("/version", router, conf)
	Block("/block", router, conf, geth)
	GasStats("/account-stats", router, conf, geth)

	// YAM
	Treasury("/treasury", router, conf, geth)
	Tvl("/tvl", router, conf, geth)
	TvlYam("/tvl/yam", router, conf, geth)
	TvlDegenerative("/tvl/degenerative", router, conf, geth)
	Apr("/apr", router, conf, geth)
	AprYam("/apr/yam", router, conf, geth)
	// AprDegenerative("/apr/degenerative", router, conf, geth)

	// Yam Synths (Degenerative)
	GetAssetsJson("/synths/assets", router, conf, geth)
	GetLatestPunkIndex("/synths/upunks/price", router, conf, geth)
	GetPunkIndexHistory("/synths/upunks/price-history", router, conf, geth)
	GetAssetsJson("/degenerative/assets", router, conf, geth)                               // to remove later
	GetLatestPunkIndex("/degenerative/upunks/price", router, conf, geth)                    // to remove later
	GetPunkIndexHistory("/degenerative/upunks/price-history", router, conf, geth)           // to remove later
  	GetLatestUStonksIndex("/degenerative/ustonks/index", router, conf, geth)                // to remove later
	GetLatestUStonksIndexHistory("/degenerative/ustonks/index-history", router, conf, geth) // to remove later

	// Yam Protection (Umbrella)

	return router
}

func Index(path string, router chi.Router, conf *config.Config) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		utils.ResJSON(http.StatusCreated, w,
			"api.yam.finance",
		)
	})
}
