package server

import (
	"github.com/xaviergodart/gydro/middlewares"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

func RunGateway(addr string) {
	apis := models.FindAllApis()
	router := NewRouter(apis)
	http.Handle("/",
		middlewares.Logger(
		middlewares.KeyAuth(
			router,
		)),
	)
	http.ListenAndServe(addr, nil)
}
