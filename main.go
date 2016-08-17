package main

import (
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/server"
	"log"
)

var gydroProxy *server.Server

func main() {
	// Open main configuration datastore
	log.Println("Initialize database connection...")
	models.InitDB("data")
	defer models.CloseDB()

	// Set example consumer and api
	consumer := models.NewConsumer("xavier", "", "")
	consumer.Save()

	backends := []string{"http://192.168.1.43:9200/"}
	api := models.NewApi("/test", backends)
	if api != nil {
		api.Save()
	}

	log.Println("Initializing Gydro proxy...")
	gydroProxy = server.NewServer()
	log.Println("Listening on localhost:8000")
	gydroProxy.ListenAndServe(":8000")
}
