package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Weather struct {
	id      int
	feeling string
}

func weatherJSON(w http.ResponseWriter, r *http.Request) {
	var curr_weather Weather
	curr_weather.id = rand.Int()
	curr_weather.feeling = time.Now().String()
	if r.Method == "GET" {
		fmt.Fprintf(w, "id: %d\nfeeling: %s", curr_weather.id, curr_weather.feeling)
	}
}

func main() {
	http.HandleFunc("/weather", weatherJSON)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
