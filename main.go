package main

import (
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/server"
	"net/http"
	"net/http/httputil"
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

	gydroProxy = &server.Proxy{Proxies: make(map[string]*httputil.ReverseProxy)}
	http.Handle("/", gydroProxy)
	http.ListenAndServe(":8000", nil)
}
