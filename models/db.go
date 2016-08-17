package models

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/nu7hatch/gouuid"
	"log"
)

var (
	store       *db.DB
	collections = map[string][]string{"Consumers": {"CustomId", "Username", "ApiKey"}, "Apis": {"Route"}} // collection and indexes
)

func InitDB(DBDir string) {
	var err error
	store, err = db.OpenDB(DBDir)
	if err != nil {
		log.Panic(err)
	}

	for col, indexes := range collections {
		if !colExists(col) {
			if err = store.Create(col); err != nil {
				log.Panic(err)
			}
		}

		collection := store.Use(col)
		for _, index := range indexes {
			if indexExists(collection, index) {
				continue
			}
			if err = collection.Index([]string{index}); err != nil {
				log.Panic(err)
			}
		}
	}
}

func indexExists(col *db.Col, index string) bool {
	for _, v := range col.AllIndexes() {
		if v[0] == index {
			return true
		}
	}
	return false
}

func colExists(name string) bool {
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
