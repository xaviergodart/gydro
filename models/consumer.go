package models

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"errors"
)

type Consumer struct {
	id     int
	Username   string `json:"username"`
	ApiKey string `json:"apikey"`
}

// NewConsumer creates new consumer. It generates an apikey if none is given.
func NewConsumer(username, apiKey string) (*Consumer, error) {
	usernameExists := FindConsumerBy("Username", username)
	apikeyExists := FindConsumerBy("ApiKey", apiKey)
	if usernameExists != nil {
		return nil, errors.New("Given username already exists")
	}
	if apiKey != "" && apikeyExists != nil {
		return nil, errors.New("Given apikey already exists")
	}

	if apiKey == "" {
		var keyExists *Consumer
		for {
			apiKey = newUuid()
			keyExists = FindConsumerBy("ApiKey", apiKey)
			if keyExists == nil {
				break
			}
		}
	}
	return &Consumer{
		id:       0,
		Username: username,
		ApiKey:   apiKey,
	}, nil
}

// GetConsumerFromInterface converts an map[string]interface{} (from tiedot) to a Consumer struct
func GetConsumerFromInterface(id int, c map[string]interface{}) *Consumer {
	var consumer Consumer
	if err := mapstructure.Decode(c, &consumer); err != nil {
		return nil
	}

	consumer.id = id
	return &consumer
}

// UpdateFromForm update consumer from form values
func (c *Consumer) UpdateFromForm(form map[string][]string) {
	for k, v := range form {
		switch k {
			case "username":
				c.Username = v[0]
			case "apikey":
				c.ApiKey = v[0]
		}
	}
}

// FindAllConsumers returns all consumers
func FindAllConsumers() []*Consumer {
	consumers := store.Use("Consumers")
	var consumersList []*Consumer = make([]*Consumer, 0)
	consumers.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		doc, err := consumers.Read(id)
		if err != nil {
			return true
		}
		consumersList = append(consumersList, GetConsumerFromInterface(id, doc))
		return true // move on to the next document
	})

	return consumersList
}

func FindConsumerByID(id int) *Consumer {
	c := FindByID("Consumers", id)
	return GetConsumerFromInterface(id, c)
}

// FindConsumerBy returns consumer matching provided field->value
func FindConsumerBy(field string, value string) *Consumer {
	results := FindBy("Consumers", []interface{}{field}, value, 1)
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

func (c *Consumer) GetId() int {
	return c.id
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
