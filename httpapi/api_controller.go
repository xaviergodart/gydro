package httpapi

import (
	"github.com/labstack/echo"
	"github.com/xaviergodart/gydro/models"
	"net/http"
)

// ApiController defines apis routes
func ApiController(e *echo.Echo) {
	e.GET("/apis/", getAllApis)
	e.GET("/apis/:name/", getApi)
	e.POST("/apis/", postApi)
	e.POST("/apis/", postApi)
	e.PATCH("/apis/:name/", patchApi)
	e.DELETE("/apis/:name/", deleteApi)
}

// getAllApis returns all apis
func getAllApis(c echo.Context) error {
	apis := models.FindAllApis()
	return c.JSON(http.StatusOK, apis)
}

// getApi returns an api for a given name
func getApi(c echo.Context) error {
	name := c.Param("name")
	api := models.FindApiBy("Name", name)
	if api == nil {
		return NewHttpError(c, 404, "Api not found")
	}
	return c.JSON(http.StatusOK, api)
}

// postApi create a new api from post values
func postApi(c echo.Context) error {
	name := c.FormValue("name")
	route := c.FormValue("route")
	backends := c.FormValue("backends")
	if name == "" || route == "" || backends == "" {
		return NewHttpError(c, 422, "Mandatory parameter is missing")
	}

	api, err := models.NewApi(name, route, c.FormParams()["backends"])
	if err != nil {
		return NewHttpError(c, 409, err.Error())
	}

	if _, err := api.Save(); err != nil {
		return NewHttpError(c, 500, "Error while creating new api")
	}

	ReloadChan<-true
	return c.JSON(http.StatusCreated, api)
}

// patchApi updates an api for a given name with given post values
func patchApi(c echo.Context) error {
	name := c.Param("name")
	api := models.FindApiBy("Name", name)
	if api == nil {
		return NewHttpError(c, 404, "Api not found")
	}

	api.UpdateFromForm(c.FormParams())
	if _, err := api.Save(); err != nil {
		return NewHttpError(c, 500, "Error while updating api")
	}

	ReloadChan<-true
	return c.JSON(200, api)
}

// deleteApi removes an api for a given name
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
