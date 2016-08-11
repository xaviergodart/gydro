package models

import (
	"log"
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func InitDB(datafile string) {
	db, err := buntdb.Open(datafile)
    if err != nil {
        log.Fatal(err)
    }
    db.SetConfig(buntdb.Config{
		SyncPolicy:           buntdb.Always,
		AutoShrinkPercentage: 100,
		AutoShrinkMinSize:    32 * 1024 * 1024,
	})
}

func CloseDB() {
	db.Close()
}

func Set(key, value string) error {
	err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
	if err != nil {
		log.Panic(err)
	}
	return err
}

func Get(key string) error {
	err := db.View(func(tx *buntdb.Tx) error {
	    _, err := tx.Get(key)
	    return err
	})
	if err != nil {
		log.Print(err)
	}
	return err
}
