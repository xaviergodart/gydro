package httpapi

import (
	"net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

var (
	NotFoundError = echo.NewHTTPError(http.StatusNotFound, "Resource not found")
)

func RunApiServer(addr string) {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"name": "Gydro Api Gateway", "version": "0.1.0"})
    })

    ApiController(e)

    e.Run(standard.New(addr))
}
