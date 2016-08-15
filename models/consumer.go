package models

import (
	"fmt"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/HouzuoGuo/tiedot/db"
)

type Consumer struct {
	id       int
    CustomId string
	Username string
	Keys     []string
}

// Create new consumer
func NewConsumer(customId, username string) *Consumer {
	return &Consumer{
		id:       0,
		CustomId: customId,
		Username: username,
		Keys:     nil,
	}
}

func FindConsumerByKey(key string) {
	consumers := store.Use("Consumers")
	var query interface{}
	json.Unmarshal([]byte(`[{"eq": "testkeys", "in": ["Keys"]}]`), &query)

	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, consumers, &queryResult); err != nil {
		panic(err)
	}

	// Query result are document IDs
	for id := range queryResult {
		// To get query result document, simply read it
		readBack, err := consumers.Read(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Query returned document %v\n", readBack)
	}
}

// Save consumer in database
// Insert a new document if the id == 0
func (c *Consumer) Save() (int, error) {
	consumers := store.Use("Consumers")
	if c.id == 0 {
		docID, err := consumers.Insert(structs.Map(c))
		c.id = docID
		return c.id, err
	} else {
		err := consumers.Update(c.id, structs.Map(c))
		return c.id, err
	}
}
