package main

import (
	"github.com/go-martini/martini"
	"encoding/json"
	"sort"
	"net/http"
	"fmt"
	"cloudapp/cloudapp/helpers"
	"reflect"
	"time"
)

// for some reason the client once work and once not , please check yours server side - http://prntscr.com/eeqx5d

func main() {
	m := martini.Classic()

	//Task 3: Show a welcome message , is it so simple???
	m.Get("/", func() string {
		current_time := time.Now().Local()
		return string("configuration and Welcome to Cloud Foundry Platform ," + "The Current date is " + current_time.Format("2006-01-02 15:04:05"))
	})

	m.Get("/buildpacks", func() string {
		b := helpers.BuildpackCustomObjects()
		jsonrsponse, _ := json.Marshal(b);
		return string(jsonrsponse)
	})

	//Task 2: Allow shuffling and sorting of the buildpacks list
	m.Get("/buildpacks/sort", func(res http.ResponseWriter, req *http.Request) {

		//http://localhost:3000/buildpacks/sort?byName
		//http://localhost:3000/buildpacks/sort?byName
		keys := reflect.ValueOf(req.URL.Query()).MapKeys()
		sortType := keys[0].String()
		fmt.Println("test", sortType)
		switch sortType {
		case "shuffle":
			b := helpers.BuildpackCustomObjects()
			//pass value by pointer
			helpers.Shuffle(&b)
			jsonrsponse, _ := json.Marshal(b);
			res.Write(jsonrsponse)
		case "byName":
			b := helpers.BuildpackCustomObjects()
			sort.Sort(helpers.SortByName(b))
			jsonrsponse, _ := json.Marshal(b);
			res.Write(jsonrsponse)
		}
	})

	m.Get("/ping", func() string {
		return "ping works"
	})

	m.Run()
}
