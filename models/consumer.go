package models

import (
	"github.com/fatih/structs"
)

type Consumer struct {
    Uuid     string
	Username string
	Keys     []string
}

// Create new consumer
// In order to generate an uuid, just pass an empty string
func NewConsumer(uuid, username string) *Consumer {
	if uuid == "" {
		uuid = newUuid()
	}
	return &Consumer{
		Uuid:     uuid,
		Username: username,
		Keys:     nil,
	}
}

// Save consumer in database
func (c *Consumer) Save() (int, error) {
	consumers := store.Use("Consumers")
	return consumers.Insert(structs.Map(c))
}
