package httpapi

import (
	"github.com/labstack/echo"
	"github.com/xaviergodart/gydro/models"
	"net/http"
	"strings"
)

func ApiController(e *echo.Echo) {
	e.GET("/apis/", getAllApis)
	e.GET("/apis/:name/", getApi)
	e.POST("/apis/", postApi)
	e.POST("/apis/", postApi)
	e.DELETE("/apis/:name/", deleteApi)
}

func getAllApis(c echo.Context) error {
	apis := models.FindAllApis()
	return c.JSON(http.StatusOK, apis)
}

func getApi(c echo.Context) error {
	name := c.Param("name")
	api := models.FindApiBy("Name", name)
	if api == nil {
		return NewHttpError(c, 404, "Api not found")
	}
	return c.JSON(http.StatusOK, api)
}

func postApi(c echo.Context) error {
	name := c.FormValue("name")
	route := c.FormValue("route")
	backends := strings.Split(c.FormValue("backends"), ",")
	if name == "" || route == "" || len(backends) == 0 || backends[0] == "" {
		return NewHttpError(c, 422, "Mandatory parameter is missing")
	}

	api, err := models.NewApi(name, route, backends)
	if err != nil {
		return NewHttpError(c, 409, err.Error())
	}

	if _, err := api.Save(); err != nil {
		return NewHttpError(c, 500, "Error while creating new api")
	}

	ReloadChan<-true
	return c.JSON(http.StatusCreated, api)
}

func deleteApi(c echo.Context) error {
	name := c.Param("name")
	api := models.FindApiBy("Name", name)
	if api == nil {
		return NewHttpError(c, 404, "Api not found")
	}

	if err := api.Delete(); err != nil {
		return NewHttpError(c, 500, "Error while deleting api")
	}

	ReloadChan<-true
	return c.NoContent(204)
}
