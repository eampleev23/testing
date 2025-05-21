package main

import (
	"github.com/eampleev23/testing.git/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", handlers.StatusHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
