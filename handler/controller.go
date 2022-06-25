package handler

import (
	"encoding/json"
	user2 "evc/entity/user"
	"evc/service/userService"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserControllerInterface interface {
	SaveUser(w http.ResponseWriter, r *http.Request)
	GetUserByEmail(w http.ResponseWriter, r *http.Request)
	GetUsersByDate(w http.ResponseWriter, r *http.Request)
	GetAllUser(w http.ResponseWriter, r *http.Request)
}
type userController struct {
	srv userService.UserServiceInterface
}

func (u *userController) SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	var user user2.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("%v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	save, err := u.srv.Save(&user)
	if err != nil {
		log.Fatalf("%v", err)
		http.Error(w, "Error saving user to db", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&save)
}

func (u *userController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := strings.Split(r.URL.Path, "/")[2]
	USER, err := u.srv.FindByEmail(email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error finding user by email", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&USER)
}

func (u *userController) GetUsersByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	start := strings.Split(r.URL.Path, "/")[3]
	end := strings.Split(r.URL.Path, "/")[4]
	limit, offset := paginate(r)
	Users, err := u.srv.FindByDate(limit, offset, start, end)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Error finding clients with give dates ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Users)

}

func (u *userController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	limit, offset := paginate(r)
	Users, err := u.srv.FindAll(limit, offset)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Error finding clients", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Users)
}

func NewUserController(srv userService.UserServiceInterface) UserControllerInterface {
	return &userController{srv: srv}
}

func paginate(r *http.Request) (int, int) {
	var page, limit, offset int
	var err error
	page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		return 10, 0
	}

	limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		return 10, 0
	}
	offset = (page - 1) * limit
	return limit, offset

}
