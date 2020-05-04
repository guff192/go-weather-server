package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Weather struct {
	Id      int    `json:"id"`
	Feeling string `json:"feeling"`
}

func weatherJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting content-type to JSON
	var curr_weather Weather
	curr_weather.Id = rand.Int() //generating ID
	curr_weather.Feeling = time.Now().String()
	if r.Method == "GET" {
		json_bytes, err := json.MarshalIndent(curr_weather, "", "    ") //Creating JSON in pretty-print format with 4-space indent
		if err != nil {
			log.Fatal("Convert struct to JSON: ", err)
		}
		fmt.Fprint(w, string(json_bytes))
	}
}

func main() {
	http.HandleFunc("/weather", weatherJSON)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
