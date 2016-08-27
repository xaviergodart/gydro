package ratelimiter

import (
	"github.com/tidwall/buntdb"
	"log"
	"time"
	"strconv"
)

var store *buntdb.DB

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
	    tx.Set(id + ":" + time.Now().Format("2006-01-02-15-04-05"), strconv.Itoa(current_values["s"] + 1), &buntdb.SetOptions{Expires:true, TTL:time.Second})
	    tx.Set(id + ":" + time.Now().Format("2006-01-02-15-04"), strconv.Itoa(current_values["m"] + 1), &buntdb.SetOptions{Expires:true, TTL:time.Minute})
	    tx.Set(id + ":" + time.Now().Format("2006-01-02-15"), strconv.Itoa(current_values["h"] + 1), &buntdb.SetOptions{Expires:true, TTL:time.Hour})
	    tx.Set(id + ":" + time.Now().Format("2006-01-02"), strconv.Itoa(current_values["d"] + 1), &buntdb.SetOptions{Expires:true, TTL:24 * time.Hour})
	    return nil
	})
}

func getCurrentValues(id string) map[string]int {
	var values map[string]int
	store.View(func(tx *buntdb.Tx) error {
		s, _ := tx.Get(id + ":" + time.Now().Format("2006-01-02-15-04-05"))
		m, _ := tx.Get(id + ":" + time.Now().Format("2006-01-02-15-04"))
		h, _ := tx.Get(id + ":" + time.Now().Format("2006-01-02-15"))
		d, _ := tx.Get(id + ":" + time.Now().Format("2006-01-02"))
		second, _ := strconv.Atoi(s)
		minute, _ := strconv.Atoi(m)
		hour, _ := strconv.Atoi(h)
		day, _ := strconv.Atoi(d)
		values = map[string]int{"s": second, "m": minute, "h": hour, "d": day}
		return nil
	})

	return values
}
