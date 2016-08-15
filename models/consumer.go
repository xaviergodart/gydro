package models

import (
	"log"
	"github.com/fatih/structs"
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

// Save consumer in database
// Insert a new document if the id == 0
func (c *Consumer) Save() error {
	consumers := store.Use("Consumers")
	if c.id == 0 {
		docID, err := consumers.Insert(structs.Map(c))
		c.id = docID
		return err
	} else {
		err := consumers.Update(c.id, structs.Map(c))
		return err
	}
}
