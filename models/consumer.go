package models

import (
	"encoding/json"
    "strings"
)

type Consumer struct {
    Uuid   string
	Name   string
	ApiKey string
}

func (c *Consumer) Save() error {
	cJson, _ := json.Marshal(c)
    var key string
    if c.Uuid != "" {
        key = strings.Join([]string{"consumer:", c.Uuid}, "")
    } else {
        for {
            uuid := newUuid()
            key = strings.Join([]string{"consumer:", uuid}, "")
            existingConsumer, _ := Get(key)
            if existingConsumer == "" {
                break
            }
            c.Uuid = uuid
        }
    }

	return Set(key, string(cJson))
}
