package models

import (
	"github.com/fatih/structs"
	"errors"
)

type Api struct {
	id       int
	Name     string
	Route    string
	Backends []string
}

// Convert an map[string]interface{} (from tiedot) to a api struct
func GetApiFromInterface(id int, a map[string]interface{}) *Api {
	var backends []string
	for _, v := range a["Backends"].([]interface{}) {
		backends = append(backends, v.(string))
	}
	return &Api{
		id:       id,
		Name:     a["Name"].(string),
		Route:    a["Route"].(string),
		Backends: backends,
	}
}

func NewApi(name, route string, backends []string) (*Api, error) {
	routeExists := FindApiBy("Route", route)
	nameExists := FindApiBy("Name", name)
	if routeExists != nil {
		return nil, errors.New("Given route already exists")
	}
	if nameExists != nil {
		return nil, errors.New("Given name already exists")
	}
	return &Api{
		id:       0,
		Route:    route,
		Name:     name,
		Backends: backends,
	}, nil
}

func FindAllApis() []*Api {
	apis := store.Use("Apis")
	var apisList []*Api = make([]*Api, 0)
	apis.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		doc, err := apis.Read(id)
		if err != nil {
			return true
		}
		apisList = append(apisList, GetApiFromInterface(id, doc))
		return true // move on to the next document
	})

	return apisList
}

func FindApiBy(field string, value string) *Api {
	results := FindBy("Apis", []interface{}{field}, value, 1)
	if len(results) == 0 {
		return nil
	}
	var api map[string]interface{}
	var apiId int
	for id, c := range results {
		api = c
		apiId = id
	}
	return GetApiFromInterface(apiId, api)
}

func FindApiByID(id int) *Api {
	api := FindByID("Apis", id)
	if api == nil {
		return nil
	}

	return GetApiFromInterface(id, api)
}

// Save saves api in database
func (a *Api) Save() (int, error) {
	apis := store.Use("Apis")
	if a.id == 0 {
		docID, err := apis.Insert(structs.Map(a))
		a.id = docID
		return a.id, err
	} else {
		err := apis.Update(a.id, structs.Map(a))
		return a.id, err
	}
}

// Delete removes api from database
func (a *Api) Delete() error {
	apis := store.Use("Apis")
	return apis.Delete(a.id)
}
