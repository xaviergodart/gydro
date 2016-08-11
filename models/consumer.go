package models

import (
	"encoding/json"
)

type Consumer struct {
	Id   int
	Name string
	Key  string
}

func (c *Consumer) Save() error {
	cJson, _ := json.Marshal(c)
	return Set(string(c.Id), string(cJson))
}

// func generateFakeConsumers() {
// 	db.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucket([]byte("Consumer"))
// 		if err != nil {
// 			return err
// 		}
// 		consumer, _ := json.Marshal(Consumer{Id: 546879654321, Name: "Xavier", ApiKey: "keytest"})
// 		err = b.Put([]byte("546879654321"), []byte(string(consumer)))
// 		return err
// 	})
// }
