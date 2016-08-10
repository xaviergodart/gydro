package main

import (
	"github.com/boltdb/bolt"
	"log"
)

var (
	db bolt.DB
)

func main() {
	// Open main configuration datastore
	log.Print("Loading Gydro conguration...")
	db, err := bolt.Open("gydro.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create fake consumer
	generateFakeConsumers()

	// Iterate over all consumers

}
