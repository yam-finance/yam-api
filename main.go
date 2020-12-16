package main

import (
	"fmt"
	"os"
	"yam-api/source"
	"yam-api/source/config"
	"yam-api/source/routes"
	"yam-api/source/utils/log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func run() error {
	fmt.Println("Starting yam server...")

	conf, confError := config.Load("./config/local.yaml")
	if confError != nil {
		fmt.Print(confError)
	}

	geth, gethError := ethclient.Dial(conf.Eth.Url.Http)
	if gethError != nil {
		fmt.Print(gethError)
	}

	routes := routes.Initialize(conf, geth)
	return source.Serve(conf, routes)
}

func main() {
	if err := run(); err != nil {
		log.Error("error :", err)
		os.Exit(1)
	}
}
