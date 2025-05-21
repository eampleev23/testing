package main

import (
	"github.com/eampleev23/testing.git/internal/handlers"
	"log"
	"net/http"
)

func main() {

	users := make(map[string]handlers.User)
	u1 := handlers.User{
		ID:        "u1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	u2 := handlers.User{
		ID:        "u2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["u1"] = u1
	users["u2"] = u2

	http.HandleFunc("/status", handlers.StatusHandler)
	http.HandleFunc("/users", handlers.UserViewHandler(users))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
