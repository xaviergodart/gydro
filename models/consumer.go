package models

import (
	"github.com/fatih/structs"
)

type Consumer struct {
	id       int
    CustomId string
	Username string
}

// Create new consumer
func NewConsumer(customId, username string) *Consumer {
	return &Consumer{
		id:       0,
		CustomId: customId,
		Username: username,
	}
}

// Convert an map[string]interface{} (from tiedot) to a Consumer struct
func GetConsumerFromInterface(id int, c map[string]interface{}) *Consumer {
	return &Consumer {
		id:       id,
		CustomId: c["CustomId"].(string),
		Username: c["Username"].(string),
	}
}

// Find consumer by api key
func FindConsumerByID(id int) *Consumer {
	c := FindByID("Consumers", id)
	return GetConsumerFromInterface(id, c)
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
