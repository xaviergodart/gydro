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

	docID, err := consumer.Save()
	if err != nil {
		log.Panic(err)
	}

	log.Println(docID)
}
