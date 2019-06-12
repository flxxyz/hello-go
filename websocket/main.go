package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	hub := NewHub()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	http.HandleFunc("/ws", hub.ServerHTTP)
	err := http.ListenAndServe(":30000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
