package userService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	userEntity "github.com/brunodelucasbarbosa/learning-go/api-gochi/src/entity"
	userRepository "github.com/brunodelucasbarbosa/learning-go/api-gochi/src/repository"
	utils "github.com/brunodelucasbarbosa/learning-go/api-gochi/src/utils"
	"github.com/go-chi/chi/v5"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}

	var user userEntity.UserInput

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

	_, responseStatusCode, responseString := userRepository.CreateUser(user)

	response := utils.Response{Message: responseString, StatusCode: responseStatusCode}
	res, _ := json.Marshal(response)
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(res))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	deleted, responseString, responseStatusCode := userRepository.DeleteUser(userId)

	if deleted {
		w.WriteHeader(responseStatusCode)
		w.Write([]byte(responseString))
	}

	response := utils.Response{Message: responseString, StatusCode: responseStatusCode}
	res, _ := json.Marshal(response)
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(res))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	find, userFind, responseString, responseStatusCode := userRepository.FindUserById(userID)
	user, _ := json.Marshal(userFind)

	if find {
		w.WriteHeader(responseStatusCode)
		w.Write([]byte(user))
		return
	}

	response := utils.Response{Message: responseString, StatusCode: responseStatusCode}
	res, _ := json.Marshal(response)
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(res))
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	find, usersFind, responseString, responseStatusCode := userRepository.FindAllUsers()
	users, _ := json.Marshal(usersFind)

	if find {
		w.WriteHeader(responseStatusCode)
		w.Write([]byte(users))
		return
	}
	response := utils.Response{Message: responseString, StatusCode: responseStatusCode}
	res, _ := json.Marshal(response)
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(res))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "id"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading request body"))
		return
	}

	var user userEntity.UserInput
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error convert request body"))
		return
	}

	if user.Age <= 5 || len(user.Name) < 5 || len(user.Taxid) != 14 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Insert all values correctly"))
		return
	}

	_, responseString, responseStatusCode := userRepository.UpdateUser(userId, user)

	response := utils.Response{Message: responseString, StatusCode: responseStatusCode}
	res, _ := json.Marshal(response)
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(res))
}
