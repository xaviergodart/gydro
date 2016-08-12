package models

import (
	"encoding/json"
    "strings"
)

type Consumer struct {
	Name   string
	ApiKey string
}

func (c *Consumer) Save() error {
	cJson, _ := json.Marshal(c)
    var key string
    for {
        uuid := newUuid()
        key = strings.Join([]string{"consumer:", uuid}, "")
        existingConsumer := Get(key)
        if existingConsumer == "" {
            break
        }
    }

	return Set(key, string(cJson))
}
