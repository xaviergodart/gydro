package main

import (
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/server"
	"log"
)

func main() {
	// Open main configuration datastore
	log.Println("Initialize database connection...")
	models.InitDB("data")
	defer models.CloseDB()

	// Set example consumer and api
	consumer := models.NewConsumer("xavier", "", "")
	log.Println(consumer)
	consumer.Save()

	backends := []string{"http://localhost:8081/", "http://localhost:8083/"}
	api := models.NewApi("/test", backends)
	if api != nil {
		api.Save()
	}

	log.Println("Listening on localhost:8000")
	server.ListenAndServe(":8000")
}
