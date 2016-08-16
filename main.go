package main

import (
	"github.com/xaviergodart/gydro/models"
	"github.com/xaviergodart/gydro/server"
	"net/url"
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

	testConf := make(map[string]*httputil.ReverseProxy)
	target, _ := url.Parse("https://google.fr")
	testConf["/test"] = httputil.NewSingleHostReverseProxy(target)

	gydroProxy = &server.Proxy{Proxies: testConf}
	http.Handle("/", gydroProxy)
	http.ListenAndServe(":8000", nil)
}
