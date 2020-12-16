package source

import (
	"net/http"
	"time"
	"yam-api/source/config"
	"yam-api/source/utils/log"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type Object map[string]interface{}

func Serve(conf *config.Config, r http.Handler) error {
	connection := conf.Server.Host + ":" + conf.Server.Port

	server := http.Server{
		Addr:        connection,
		ReadTimeout: time.Second * 15,
		Handler:     r,
	}

	log.Info("Started on", "http://"+connection)
	err := server.ListenAndServe()
	if err != nil {
		log.Error(err)
	}

	return nil

}
