package httpapi

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func RunApiServer() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Run(standard.New(":1323"))
}
