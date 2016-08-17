package main

import (
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/server"
	"log"
)

var gydroProxy *server.Proxy

func main() {
	// Open main configuration datastore
	log.Print("Initialize database connection...")
	models.InitDB("data")
	defer models.CloseDB()

	// Set example consumer
	consumer := models.NewConsumer("xavier", "", "")

	_, err := consumer.Save()
	if err != nil {
		log.Panic(err)
	}

	gydroProxy = server.NewProxy()
	gydroProxy.ListenAndServe(":8000")
}
