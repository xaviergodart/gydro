package models

import (
    "github.com/fatih/structs"
)

type Api struct {
	id       int
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
        Route:    a["Route"].(string),
        Backends: backends,
    }
}

func NewApi(route string, backends []string) *Api {
    apiExists := FindApiByRoute(route)
    if apiExists != nil {
        return nil
    }
    return &Api{
        id:       0,
        Route:    route,
        Backends: backends,
    }
}

func FindAllApis() []*Api {
    apis := store.Use("Apis")
    var apisList []*Api
    apis.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
        doc, err := apis.Read(id)
        if err != nil {
            return true
        }
        apisList = append(apisList, GetApiFromInterface(id, doc))
        return true  // move on to the next document
    })

    return apisList
}

func FindApiByRoute(route string) *Api {
    results := FindBy("Apis", []interface{}{"Route"}, route, 1)
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

// Save api in database
// Insert a new document if the id == 0
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
