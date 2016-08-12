package main

import (
	"log"
	"github.com/xaviergodart/gydro/models"
)

func main() {
	// Open main configuration datastore
	log.Print("Initialize database connection...")
	models.InitDB("data.db")
	defer models.CloseDB()

	// Set example consumer
	consumer := &models.Consumer{
		Name: "Xavier",
		ApiKey: "testkey",
	}

	err := consumer.Save()
	if err != nil {
		log.Fatal(err)
	}
}
