package userService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dbConfig "github.com/brunodelucasbarbosa/learning-go/api-gochi/config"
	userRepository "github.com/brunodelucasbarbosa/learning-go/api-gochi/repository"
	"github.com/go-chi/chi"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}

	var user userRepository.UserInput

	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error convert request body"))
		return
	}

	if user.Age <= 5 || len(user.Name) < 5 || len(user.Taxid) != 14 || len(user.Password) < 5 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Insert all values correctly"))
		return
	}

	db, err := dbConfig.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO "user" (name, age, taxid, password) VALUES ($1, $2, $3, $4)`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error to insert value in database"))
		fmt.Println(err.Error())
		return
	}
	_, err = statement.Exec(user.Name, user.Age, user.Taxid, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User is already registered"))
		return
	}

	defer statement.Close()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	fmt.Println(userID)
	// "user.id",
	// "user.name",
	// "user.age",
	// "user.taxid"
	var user userRepository.UserResponse

	db, err := dbConfig.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	row, err := db.Query(`SELECT id, name, age, taxidFROM "user" WHERE id = $1`, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User is not registered"))
		fmt.Println(err.Error())
		return
	}
	defer row.Close()

	for row.Next() {
		row.Scan(&user.ID, &user.Name, &user.Age, &user.Taxid)
	}

	response, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []userRepository.UserResponse

	db, err := dbConfig.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, age, taxid FROM "user"`)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error to find all users"))
		return
	}
	for rows.Next() {
		var user userRepository.UserResponse
		rows.Scan(&user.ID, &user.Name, &user.Age, &user.Taxid)
		users = append(users, user)
	}
	defer rows.Close()

	response, _ := json.Marshal(users)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}
