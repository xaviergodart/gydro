package main

import (
	"encoding/json"
	"github.com/boltdb/bolt"
)

type Consumer struct {
	Id     int
	Name   string
	ApiKey string
}

var (
	AuthorizedConsumers map[string]*Consumer
)

func generateFakeConsumers() {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Consumer"))
		if err != nil {
			return err
		}
		consumer, _ := json.Marshal(Consumer{Id: 546879654321, Name: "Xavier", ApiKey: "keytest"})
		err = b.Put([]byte("546879654321"), []byte(string(consumer)))
		return err
	})
}

// func LoadConsumersFromDd() {
// 	err = db.View(func(tx *bolt.Tx) error {
// 		consumerBucket := tx.Bucket([]byte("Consumer"))
// 		consumerBucket.ForEach(func(k, v []byte) error {
// 			fmt.Printf("key=%s, value=%s\n", k, v)
// 			return nil
// 		})
// 		return nil
// 	})
// }
