package main

import (
	"github.com/xaviergodart/gydro/models"
	"log"
)

func main() {
	// Open main configuration datastore
	log.Print("Initialize database connection...")
	models.InitDB("data")
	defer models.CloseDB()

	// Set example consumer
	consumer := models.NewConsumer("xavier", "", "")

	_, err := consumer.Save()
	if err != nil {
		log.Panic(err)
	}

	log.Println(models.FindConsumerByApiKey(consumer.ApiKey))
}
