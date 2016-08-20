package httpapi

import (
	"github.com/labstack/echo"
	"github.com/xaviergodart/gydro/models"
	"net/http"
	"strconv"
)

func ApiController(e *echo.Echo) {
	e.GET("/apis/", getAllApis)
	e.GET("/apis/:id/", getApi)
}

func getApi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	api := models.FindApiByID(id)
	if api == nil {
		return c.JSON(http.StatusNotFound, NotFoundError)
	}
	return c.JSON(http.StatusOK, api)
}

func getAllApis(c echo.Context) error {
	apis := models.FindAllApis()
	return c.JSON(http.StatusOK, apis)
}
