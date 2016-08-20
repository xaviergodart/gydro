package httpapi

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

var (
	NotFoundError = echo.NewHTTPError(http.StatusNotFound, "Resource not found")
)

func RunApiServer(addr string) {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"name": "Gydro Api Gateway", "version": "0.1.0"})
	})

	ApiController(e)

	e.Run(standard.New(addr))
}
