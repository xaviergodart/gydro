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
		Uuid: "de0b72d7-79bf-4480-6fbb-cfe2130e423d",
		Name: "Xavier",
		ApiKey: "testkey",
	}

	err := consumer.Save()
	if err != nil {
		log.Fatal(err)
	}
}
