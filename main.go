package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type test_struct struct {
	Request_Type string
	Start_Date string
	End_Date string
	Data_Content string
}


func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		log.Println("Returning current schedule")
		w.Write([]byte(`{"Request_Type": "get_schedule", "Start_Date": "2021-January-1", "End_Date": "2021-January-1"}`))
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		if (t.Request_Type == "data_schedule") {
			log.Println("Data scheduling requested")
			res := []byte(`{"Request_Type":"` + t.Request_Type + `", "Start_Date": "` + t.Start_Date + `", "End_Date": "` + t.End_Date + `"}`)
			w.Write(res)
		} else if (t.Request_Type == "edit_schedule") {
			log.Println("Edit scheduling requested")
			res := []byte(`{"Request_Type":"` + t.Request_Type + `", "Start_Date": "` + t.Start_Date + `", "End_Date": "` + t.End_Date + `"}`)
			w.Write(res)
		} else if (t.Request_Type == "map_generation") {
			log.Println("Map generation requested")
			res := []byte(`{"Request_Type":"` + t.Request_Type + `", "Data_Content": "` + t.Data_Content + `"}`)
			w.Write(res)
		} else if (t.Request_Type == "cancel_schedule") {
			log.Println("Cancelling schedule")
			res := []byte(`{"Request_Type:"` + t.Request_Type + `"}`)
			w.Write(res)
		}
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