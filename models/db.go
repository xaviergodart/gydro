package models

import (
	"log"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/nu7hatch/gouuid"
)

var store *db.DB

func InitDB(DBDir string) {
	var err error
	store, err = db.OpenDB(DBDir)
	if err != nil {
		log.Panic(err)
	}

	if colExists("Consumers") {
		return
	}

	if err = store.Create("Consumers"); err != nil {
		log.Panic(err)
	}

	consumers := store.Use("Cosumers")
	if err = consumers.Index([]string{"Keys"}); err != nil {
		log.Panic(err)
	}
}

func colExists(name string) (bool) {
	for _, v := range store.AllCols() {
		if v == name {
			return true
		}
	}
	return false
}

func CloseDB() {
	if err := store.Close(); err != nil {
		log.Panic(err)
	}
}

func newUuid() string {
	nUuid, _ := uuid.NewV4()
	return nUuid.String()
}
