package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// Open main configuration datastore
	log.Print("Loading Gydro conguration...")
	db, err := bolt.Open("gydro.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//
}

