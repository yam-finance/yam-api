package main

import (
	"fmt"

	"os"
	"yam-api/source"
	"yam-api/source/config"
	"yam-api/source/routes"
	"yam-api/source/utils/log"
	"yam-api/source/utils/mongodb"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron"
)

func run() error {
	fmt.Println("Starting yam server...")

	conf, confError := config.Load("./config/prod.yaml")
	if confError != nil {
		fmt.Print(confError)
	}

	geth, gethError := ethclient.Dial(conf.Eth.Url.Http)
	if gethError != nil {
		fmt.Print(gethError)
	}

	mongodb.Connect()

	calculatePunkIndexCron := cron.New()
	calculatePunkIndexCron.AddFunc("@every 5m", func() {
		values := routes.CalculatePunkIndex(geth)
		mongodb.InsertPunkIndex(values)
	})
	calculatePunkIndexCron.Start()

	getAprYamCron := cron.New()
	getAprYamCron.AddFunc("@every 5m", func() {
		val := routes.CalculateAprYam(geth)
		routes.StoreAprYam(val)
	})
	getAprYamCron.Start()

	/*getAprDegenerativeCron := cron.New()
	getAprDegenerativeCron.AddFunc("@every 5m", func() {
		val := routes.CalculateAprDegenerative(geth)
		routes.StoreAprDegenerative(val)
	})
	getAprDegenerativeCron.Start()*/

	routes := routes.Initialize(conf, geth)
	return source.Serve(conf, routes)
}

func main() {
	if err := run(); err != nil {
		log.Error("error :", err)
		os.Exit(1)
	}
}
