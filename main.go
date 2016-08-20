package main

import (
	"github.com/xaviergodart/gydro/httpapi"
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

	backends := []string{"http://localhost:8081/", "http://localhost:8082/"}
	api := models.NewApi("test", "/test", backends)
	if api != nil {
		api.Save()
	}
	log.Println(api)

	backends2 := []string{"http://localhost:8083/", "http://localhost:8084/"}
	api2 := models.NewApi("testdata", "/test/data", backends2)
	if api2 != nil {
		api2.Save()
	}
	log.Println(api2)

	log.Println("Gateway listening on localhost:8000")
	go server.RunGateway(":8000")
	log.Println("Api listening on localhost:8001")
	go httpapi.RunApiServer(":8001")

	select {}
}
