package httpapi

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var (
	ReloadChan chan bool
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RunApiServer(addr string, reload chan bool) {
	ReloadChan = reload

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"name": "Gydro Api Gateway", "version": "0.1.0"})
	})

	ApiController(e)
	ConsumerController(e)

	e.Run(standard.New(addr))
}

func NewHttpError(c echo.Context, code int, msg string) (err error) {
	httperr := &HttpError{Code: code, Message: msg}
	return c.JSON(code, httperr)
}
