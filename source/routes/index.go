package routes

import (
	"math/big"
	"net/http"
	"yam-api/source/config"
	"yam-api/source/utils"

	"yam-api/source/utils/mongodb"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator"
)

type Assets struct {
	Ugas    []AssetInstance
	Ustonks []AssetInstance
}
type Asset struct {
	AssetName     string
	AssetInstance AssetInstance
	AssetPrice    *big.Float
}
type AssetInstance struct {
	Name       string
	Cycle      string
	Year       string
	Collateral string
	Token      Token
	Emp        Emp
	Pool       Pool
	Apr        AprData
}
type Emp struct {
	Address string
	New     bool
}
type Pool struct {
	Address string
}
type AprData struct {
	Force int
	Extra int
}
type Token struct {
	Address  string
	Decimals int
}

var validate *validator.Validate

func Initialize(conf *config.Config, geth *ethclient.Client) chi.Router {

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	mongodb.Connect()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler)

	// /
	Index("/", router, conf)
	Version("/version", router, conf)
	Block("/block", router, conf, geth)
	Misc("/misc", router, conf)

	// YAM
	Treasury("/treasury", router, conf, geth)

	Apr("/apr", router, conf, geth)
	AprYam("/apr/yam", router, conf, geth)
	AprDegenerative("/apr/degenerative", router, conf, geth)

	return router
}

func Index(path string, router chi.Router, conf *config.Config) {
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		utils.ResJSON(http.StatusCreated, w,
			"api.yam.finance",
		)
	})
}
