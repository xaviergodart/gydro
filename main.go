package main

import (
	"log"
	"github.com/xaviergodart/gydro/models"
)

func main() {
	// Open main configuration datastore
	log.Print("Initialize database connection...")
	models.InitDB("data")
	defer models.CloseDB()

	// Set example consumer
	consumer := models.NewConsumer("", "xavier")
	consumer.Keys = []string{"testkey"}

	_, err := consumer.Save()
	if err != nil {
		log.Panic(err)
	}

	log.Println(consumer)

	models.FindConsumerByKey("testkey")
}
