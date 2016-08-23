package httpapi

import (
    "github.com/labstack/echo"
    "github.com/xaviergodart/gydro/models"
    "net/http"
)

// ConsumerController defines consumers routes
func ConsumerController(e *echo.Echo) {
    e.GET("/consumers/", getAllConsumers)
    e.GET("/consumers/:username/", getConsumer)
    e.POST("/consumers/", postConsumer)
}

// getAllConsumers returns all consumers
func getAllConsumers(c echo.Context) error {
    consumers := models.FindAllConsumers()
    return c.JSON(http.StatusOK, consumers)
}

// getConsumer returns a consumer for a given username
func getConsumer(c echo.Context) error {
    username := c.Param("username")
    consumer := models.FindConsumerBy("Username", username)
    if consumer == nil {
        return NewHttpError(c, 404, "Consumer not found")
    }
    return c.JSON(http.StatusOK, consumer)
}

// postConsumer create a new consumer from given values
func postConsumer(c echo.Context) error {
    username := c.FormValue("username")
    apikey := c.FormValue("apikey")
    if username == "" {
        return NewHttpError(c, 422, "Mandatory parameter is missing")
    }

    consumer, err := models.NewConsumer(username, apikey)
    if err != nil {
        return NewHttpError(c, 409, err.Error())
    }

    if _, err := consumer.Save(); err != nil {
        return NewHttpError(c, 500, "Error while creating new consumer")
    }

    return c.JSON(http.StatusCreated, consumer)
}
