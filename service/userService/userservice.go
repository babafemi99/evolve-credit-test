package userService

import (
	"evc/entity/user"
	"evc/repository"
	"strconv"
	"time"
)

type UserServiceInterface interface {
	FindByEmail(email string) (*user.User, error)
	FindByDate(limit, offset int, start, end time.Time) (*user.Users, error)
	FindAll(limit, offset int) (*user.Users, error)
}

type userService struct {
	repo repository.UserRepoInterface
}

func (u *userService) FindByEmail(email string) (*user.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *userService) FindByDate(limit, offset int, start, end time.Time) (*user.Users, error) {
	newLimit := strconv.Itoa(limit)
	newOffset := strconv.Itoa(offset)
	return u.repo.GetByDate(newLimit, newOffset, start, end)
}

func (u *userService) FindAll(limit, offset int) (*user.Users, error) {
	newLimit := strconv.Itoa(limit)
	newOffset := strconv.Itoa(offset)
	return u.repo.GetAllUsers(newLimit, newOffset)
}

func NewUserService(repo repository.UserRepoInterface) UserServiceInterface {
	return &userService{repo: repo}
}
