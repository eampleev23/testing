package handlers

import (
	"encoding/json"
	"net/http"
)

// User - основной объект для теста
type User struct {
	ID        string
	FirstName string
	LastName  string
}

// UserViewHandler - хэндлер, который нужно протестировать.
func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			http.Error(rw, "user id is required", http.StatusBadRequest)
			return
		}
		user, ok := users[userId]
		if !ok {
			http.Error(rw, "user not found", http.StatusNotFound)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, "json error", http.StatusInternalServerError)
			return
		}
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonUser)
	}
}
