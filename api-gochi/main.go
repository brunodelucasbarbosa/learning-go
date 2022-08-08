package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/user/{userID}", findUser)
	r.Post("/user", createUser)

	fmt.Println("Server started on port 3000......")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	t, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(t)
}

func findUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	users := []User{}
	users = append(users, User{
		Name: "John Doe 1",
		ID:   1,
	})
	users = append(users, User{
		Name: "John Doe 2",
		ID:   2,
	})
	users = append(users, User{
		Name: "John Doe 3",
		ID:   3,
	})

	for _, u := range users {
		if userIDParsed, _ := strconv.Atoi(userID); u.ID == userIDParsed {
			user, _ := json.Marshal(u)
			w.WriteHeader(http.StatusOK)
			w.Write(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("User not found!"))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
