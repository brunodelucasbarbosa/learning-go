package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/brunodelucasbarbosa/learning-go/api-gochi/repository/userRepository"

	userService "github.com/brunodelucasbarbosa/learning-go/api-gochi/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/user/{userID}", userService.FindUserById)
	r.Get("/users", userService.FindAllUsers)
	r.Post("/user", userService.CreateUser)

	fmt.Println("Server started on port 3000......")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
