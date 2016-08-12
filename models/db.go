package models

import (
	"log"
	"github.com/tidwall/buntdb"
	"github.com/nu7hatch/gouuid"
)

var db *buntdb.DB

func InitDB(datafile string) {
	var err error
	db, err = buntdb.Open(datafile)
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
	return db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

func Get(key string) (string, error) {
	var val string
	err := db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(key)
		if err != nil{
	        return err
	    }
	    val = value
	    return nil
	})
	if err != nil {
		return "", err
	}

	return val, nil
}

func newUuid() string {
	nUuid, _ := uuid.NewV4()
	return nUuid.String()
}
