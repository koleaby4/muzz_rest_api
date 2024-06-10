package handlers

import (
	"context"
	"encoding/json"
	"github.com/koleaby4/muzz_rest_api/db"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req db.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := db.GetQueries()

	user, err := q.CreateUser(context.Background(), db.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Gender:   req.Gender,
		Age:      req.Age,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
