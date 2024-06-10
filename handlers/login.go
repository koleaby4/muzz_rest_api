package handlers

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/koleaby4/muzz_rest_api/db"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var cookieStore = sessions.NewCookieStore([]byte("DUMMY_SESSION_KEY"))

func getToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, _ := token.SignedString([]byte("secret-key"))
	return tokenString
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := []byte(r.FormValue("password"))

	queries := db.GetQueries()
	user, err := queries.GetUser(context.Background(), email)
	if err != nil {
		log.Printf("error (%v) fetching user with email=%v\n", err, email)
	}

	if user.ID == nil { // user does not exist
		log.Println("User does not exist", user)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Println("User exists", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)

	if err != nil {
		fmt.Println("Passwords do not match")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["UserID"] = user.ID
	session.Values["AuthToken"] = getToken(user.Email)

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
