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
		Id: 8768768,
		Name: "Xavier",
		Key: "testkey",
	}

	consumer.Save()
}
