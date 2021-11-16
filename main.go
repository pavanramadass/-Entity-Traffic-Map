package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type test_struct struct {
	Request_Type string
	Arg1 string
	Arg2 string
}


func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Println(t.Request_Type)
		log.Println(t.Arg1)
		log.Println(t.Arg2)
		a := []byte(`{"Request_Type":"` + t.Request_Type + `", "Arg1": "` + t.Arg1 + `", "Arg2": "` + t.Arg2 + `"}`)
		log.Println(a)
		w.Write(a)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/form", handleRequest)
	log.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}