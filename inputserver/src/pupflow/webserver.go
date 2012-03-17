package main

import (
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

func startWebserver(addr string) {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/axis", handler)
	e := http.ListenAndServe(addr, nil)
	panic(e)
}

type DataPackage struct {
	ID uint16
	Object SceneObject
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		d := json.NewDecoder(r.Body)
		var data DataPackage
		e := d.Decode(&data)
		if e != nil {
			log.Printf("Could not parse data: %s\n", e)
			return
		}
		if data.Object.Name == "" {
			delete(config, data.ID)
			return
		}
		config[data.ID] = data.Object
	} else if r.Method == "GET" {
		strmap := make(map[string]SceneObject)
		for k, v := range config {
			strmap[strconv.Itoa(int(k))] = v
		}
		s, _ := json.Marshal(strmap)
		w.Header().Set("Content-Type", "application/json")
		w.Write(s)
	}
}
