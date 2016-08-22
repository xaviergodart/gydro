package server

import (
	"github.com/xaviergodart/gydro/middlewares"
	"github.com/xaviergodart/gydro/models"
	"github.com/mailgun/manners"
	"log"
)

func RunGateway(addr string, reload chan bool, done chan bool) {
	for {
		apis := models.FindAllApis()
		router := NewRouter(apis)
		handler := middlewares.Logger(
					middlewares.KeyAuth(
						router,
					))

		go manners.ListenAndServe(addr, handler)

		select {
			case <-reload:
				log.Println("Configuration updated : reload gateway...")
				manners.Close()
			case <-done:
				manners.Close()
				log.Println("Goodbye.")
				return
		}
	}
}
