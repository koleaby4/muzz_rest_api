package handlers

import (
	"context"
	"encoding/json"
	"github.com/koleaby4/muzz_rest_api/db"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateUserHandler handles the POST /user/create route
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var inUser db.User
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inUser.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	inUser.Password = string(hashedPassword)

	q := db.GetQueries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbUser, err := q.CreateUser(context.Background(), inUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dbUser)
}
