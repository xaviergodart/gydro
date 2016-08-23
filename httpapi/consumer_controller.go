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
    e.PATCH("/consumers/:username/", patchConsumer)
    e.PUT("/consumers/", putConsumer)
    e.DELETE("/consumers/:username/", deleteConsumer)
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

// patchConsumer updates a consumer for a given username with given values
func patchConsumer(c echo.Context) error {
    username := c.Param("username")
    consumer := models.FindConsumerBy("Username", username)
    if consumer == nil {
        return NewHttpError(c, 404, "Consumer not found")
    }

    consumer.UpdateFromForm(c.FormParams())
    if _, err := consumer.Save(); err != nil {
        return NewHttpError(c, 500, "Error while updating consumer")
    }

    return c.JSON(200, consumer)
}

// putConsumer creates or updates a consumer with given values
func putConsumer(c echo.Context) error {
    username := c.FormValue("username")
    if username == "" {
        return NewHttpError(c, 422, "Username parameter is missing")
    }

    consumer := models.FindConsumerBy("Username", username)
    if consumer == nil {
        return postConsumer(c)
    } else {
        c.SetParamNames("username")
        c.SetParamValues(username)
        return patchConsumer(c)
    }
}

// deleteConsumer removes a consumer for a given username
func deleteConsumer(c echo.Context) error {
    username := c.Param("username")
    consumer := models.FindConsumerBy("Username", username)
    if consumer == nil {
        return NewHttpError(c, 404, "Consumer not found")
    }

    if err := consumer.Delete(); err != nil {
        return NewHttpError(c, 500, "Error while deleting consumer")
    }

    return c.NoContent(204)
}
