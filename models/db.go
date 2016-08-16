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

	consumers := store.Use("Consumers")
	if err = consumers.Index([]string{"CustomId"}); err != nil {
		log.Panic(err)
	}
	if err = consumers.Index([]string{"Username"}); err != nil {
		log.Panic(err)
	}
	if err = consumers.Index([]string{"ApiKey"}); err != nil {
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

func FindByID(col string, id int) map[string]interface{} {
	collection := store.Use(col)
	readBack, err := collection.Read(id)
	if err != nil {
		log.Panic(err)
	}

	return readBack
}

func FindBy(col string, field []interface{}, val interface{}, limit int) map[int]map[string]interface{} {
	collection := store.Use(col)
	query := map[string]interface{}{
	   "eq":    val,
	   "in":    field,
	   "limit": limit,
	}

	queryResult := make(map[int]struct{})
	if err := db.EvalQuery(query, collection, &queryResult); nil != err {
		log.Panic(err)
	}

	results := make(map[int]map[string]interface{})
	for id := range queryResult {
		readBack, err := collection.Read(id)
		if nil != err {
			log.Panic(err)
		}
		results[id] = readBack
	}

	return results
}

func newUuid() string {
	nUuid, _ := uuid.NewV4()
	return nUuid.String()
}
