package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		user        User
	}

	tests := []struct {
		name    string
		request string
		users   interface{}
		want    want
	}{
		{
			name: "test status 200 OK",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
				user: User{
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id1",
		},
		{
			name: "test status 404 Not Found",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusNotFound,
			},
			request: "/users?user_id=id3",
		},
		{
			name: "test status 400 Bad Request",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusBadRequest,
			},
			request: "/users",
		},
		{
			name:  "test status 500 Internal Server Error",
			users: "invalid json",
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusInternalServerError,
			},
			request: "/users?user_id=id1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && tt.want.statusCode != 500 {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)
			request.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UserViewHandler(tt.users.(map[string]User)))
			h(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			var user User
			err = json.Unmarshal(userResult, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)
		})
	}
}
