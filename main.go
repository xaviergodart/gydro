package main

import (
	"github.com/xaviergodart/gydro/httpapi"
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Open datastore
	log.Println("Initialize database connection...")
	models.InitDB("data")
	defer models.CloseDB()

	// reload channel is used to reload the gateway configuration
	reload := make(chan bool)
	// done channel is used to stop the gateway and exit
	done := make(chan bool)

	// Listen to SIGINT and SIGTERM to gracefully shutdown the gateway
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, os.Kill)
		<-sigchan
		done<-true
	}()

	log.Println("Api listening on localhost:8001")
	go httpapi.RunApiServer(":8001", reload)

	log.Println("Gateway listening on localhost:8000")
	server.RunGateway(":8000", reload, done)

	os.Exit(0)
}
