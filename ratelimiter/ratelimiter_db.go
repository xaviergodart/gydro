package ratelimiter

import (
	"github.com/tidwall/buntdb"
	"log"
	"strconv"
	"time"
)

var (
	store *buntdb.DB
	limiterConfiguration = map[string]map[string]interface{}{"s": {"format": "2006-01-02-15-04-05", "expires": time.Second}, "m": {"format": "2006-01-02-15-04", "expires": time.Minute}, "h": {"format": "2006-01-02-15", "expires": time.Hour}, "d": {"format": "2006-01-02", "expires": 24 * time.Hour}}
)

func InitDB(DBDir string) {
	var err error
	store, err = buntdb.Open(DBDir)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	store.Close()
}

func IsExceeded(id int, limits map[string]int) bool {
	current_values := getCurrentValues(strconv.Itoa(id))
	for k, v := range current_values {
		if limits[k] != 0 && v >= limits[k] {
			return true
		}
	}

	incr(strconv.Itoa(id), current_values)
	return false
}

func incr(id string, current_values map[string]int) {
	store.Update(func(tx *buntdb.Tx) error {
		for k, v := range limiterConfiguration {
			tx.Set(id+":"+time.Now().Format(v["format"].(string)), strconv.Itoa(current_values[k]+1), &buntdb.SetOptions{Expires: true, TTL: v["expires"].(time.Duration)})
		}
		return nil
	})
}

func getCurrentValues(id string) map[string]int {
	values := map[string]int{"s": 0, "m": 0, "h": 0, "d": 0}
	store.View(func(tx *buntdb.Tx) error {
		for k, v := range limiterConfiguration {
			counterStr, _ := tx.Get(id + ":" + time.Now().Format(v["format"].(string)))
			values[k], _ = strconv.Atoi(counterStr)
		}
		return nil
	})

	return values
}
