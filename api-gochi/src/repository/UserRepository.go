package userRepository

import (
	"net/http"

	dbConfig "github.com/brunodelucasbarbosa/learning-go/api-gochi/src/config"
	userEntity "github.com/brunodelucasbarbosa/learning-go/api-gochi/src/entity"
)

func CreateUser(user userEntity.UserInput) (bool, int, string) {
	db, err := dbConfig.Connect()
	defer db.Close()

	if err != nil {
		return false, http.StatusInternalServerError, "Error connecting to database"
	}
	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO "user" (name, age, taxid, password) VALUES ($1, $2, $3, $4)`)

	if err != nil {
		return false, http.StatusInternalServerError, "Error to insert value in database"
	}

	_, err = statement.Exec(user.Name, user.Age, user.Taxid, user.Password)

	if err != nil {
		return false, http.StatusBadRequest, "User is already registered"
	}

	return true, http.StatusCreated, "User created successfully"
}

func DeleteUser(userId int) (bool, string, int) {
	db, err := dbConfig.Connect()
	if err != nil {
		return false, "Error connecting to database", http.StatusInternalServerError
	}
	defer db.Close()

	row, err := db.Query(`SELECT id, name, age, taxid FROM "user" WHERE id = $1`, userId)
	if err != nil {
		return false, "Error to find user", http.StatusBadRequest
	}

	var user userEntity.UserResponse
	for row.Next() {
		row.Scan(&user.ID, &user.Name, &user.Age, &user.Taxid)
	}
	defer row.Close()

	if (userEntity.UserResponse{}) == user { //se tiver vazio, nao ter encontrado
		return false, "User not found", http.StatusNotFound
	}

	statement, err := db.Prepare(`DELETE FROM "user" WHERE id = $1`)

	if err != nil {
		return false, "Error to delete user", http.StatusBadRequest
	}

	if _, err := statement.Exec(userId); err != nil {
		return false, "Error to delete user", http.StatusBadRequest
	}

	return true, "User deleted", http.StatusNoContent
}

func FindUserById(userId int) (bool, userEntity.UserResponse, string, int) {
	db, err := dbConfig.Connect()
	if err != nil {
		return false, userEntity.UserResponse{}, "Error connecting to database", http.StatusBadRequest
	}
	defer db.Close()

	row, err := db.Query(`SELECT id, name, age, taxid FROM "user" WHERE id = $1`, userId)
	if err != nil {
		return false, userEntity.UserResponse{}, "Error to find user", http.StatusBadRequest
	}

	var user userEntity.UserResponse
	for row.Next() {
		row.Scan(&user.ID, &user.Name, &user.Age, &user.Taxid)
	}
	defer row.Close()

	if (userEntity.UserResponse{}) == user { //se tiver vazio, nao ter encontrado
		return false, userEntity.UserResponse{}, "User not found", http.StatusNotFound
	}

	return true, user, "User found", http.StatusOK
}

func FindAllUsers() (bool, []userEntity.UserResponse, string, int) {
	var users []userEntity.UserResponse

	db, err := dbConfig.Connect()
	if err != nil {
		return false, []userEntity.UserResponse{}, "Error connecting to database", http.StatusBadRequest
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, age, taxid FROM "user"`)
	if err != nil {
		return false, []userEntity.UserResponse{}, "Error to find all users", http.StatusInternalServerError
	}

	for rows.Next() {
		var user userEntity.UserResponse
		rows.Scan(&user.ID, &user.Name, &user.Age, &user.Taxid)
		users = append(users, user)
	}
	defer rows.Close()

	return true, users, "Users found", http.StatusOK
}

func UpdateUser(userId int, userInput userEntity.UserInput) (bool, string, int) {
	db, err := dbConfig.Connect()
	if err != nil {
		return false, "Error connecting to database", http.StatusInternalServerError
	}
	defer db.Close()

	find, _, responseString, responseStatusCode := FindUserById(userId)

	if !find {
		return false, responseString, responseStatusCode
	}

	statement, err := db.Prepare(`UPDATE "user" SET name = $1, age = $2, taxid = $3 WHERE id = $4`)

	if err != nil {
		return false, err.Error(), http.StatusInternalServerError
	}

	if _, err := statement.Exec(userInput.Name, userInput.Age, userInput.Taxid, userId); err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "User updated", http.StatusOK
}
