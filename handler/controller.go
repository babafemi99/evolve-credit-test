package handler

import (
	"encoding/json"
	"evc/service/userService"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-type", "application/json")
	panic("implement me")
}

func (u *userController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	user, err := u.srv.FindByEmail("email")
	if err != nil {
		http.Error(w, "Error finding user by email", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func (u *userController) GetUsersByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//var date dateInput.Date
	//var err error
	//start := strings.Split(r.URL.Path/)
	panic("Implement me")

}

func (u *userController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	limit, offset := paginate(r)
	Users, err := u.srv.FindAll(limit, offset)
	if err != nil {
		http.Error(w, "Error finding clients", http.StatusInternalServerError)
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
