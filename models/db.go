package models

import (
	"log"
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func InitDB(datafile string) {
	db, err := buntdb.Open(datafile)
	db.SetConfig(buntdb.Config{
		SyncPolicy:           buntdb.Always,
		AutoShrinkPercentage: 100,
		AutoShrinkMinSize:    32 * 1024 * 1024,
	})
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}
