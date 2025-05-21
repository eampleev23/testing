package handlers

import (
	"encoding/json"
	"net/http"
)

type jsonRespError struct {
	errorMsg string `json:"error"`
}

// User - основной объект для теста
type User struct {
	ID        string
	FirstName string
	LastName  string
	age       int
}

// UserViewHandler - хэндлер, который нужно протестировать.
func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("content-type", "application/json")
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			errJson := jsonRespError{errorMsg: "Bad request"}
			msg, _ := json.Marshal(errJson)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(msg)
			return
		}
		user, ok := users[userId]
		if !ok {
			errJson := jsonRespError{errorMsg: "User not found"}
			msg, _ := json.Marshal(errJson)
			rw.WriteHeader(http.StatusNotFound)
			rw.Write(msg)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			errJson := jsonRespError{errorMsg: "No valid JSON object for user"}
			msg, _ := json.Marshal(errJson)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(msg)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonUser)
	}
}
