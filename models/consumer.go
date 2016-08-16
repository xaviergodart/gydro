package models

import (
	"github.com/fatih/structs"
)

type Consumer struct {
	id       int
	CustomId string
	Username string
	ApiKey   string
}

// Create new consumer
func NewConsumer(username, customId, apiKey string) *Consumer {
	if apiKey == "" {
		var keyExists *Consumer
		for {
			apiKey = newUuid()
			keyExists = FindConsumerByApiKey(apiKey)
			if keyExists == nil {
				break
			}
		}
	}
	return &Consumer{
		id:       0,
		CustomId: customId,
		Username: username,
		ApiKey:   apiKey,
	}
}

// Convert an map[string]interface{} (from tiedot) to a Consumer struct
func GetConsumerFromInterface(id int, c map[string]interface{}) *Consumer {
	return &Consumer{
		id:       id,
		CustomId: c["CustomId"].(string),
		Username: c["Username"].(string),
		ApiKey:   c["ApiKey"].(string),
	}
}

func FindConsumerByID(id int) *Consumer {
	c := FindByID("Consumers", id)
	return GetConsumerFromInterface(id, c)
}

func FindConsumerByApiKey(key string) *Consumer {
	results := FindBy("Consumers", []interface{}{"ApiKey"}, key, 1)
	if len(results) == 0 {
		return nil
	}
	var consumer map[string]interface{}
	var consumerId int
	for id, c := range results {
		consumer = c
		consumerId = id
	}
	return GetConsumerFromInterface(consumerId, consumer)
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
