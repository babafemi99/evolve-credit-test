package userService

import (
	"evc/entity/user"
	"evc/repository"
	"strconv"
	"time"
)

type UserServiceInterface interface {
	Save(user2 *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	FindByDate(limit, offset int, start, end time.Time) ([]user.User, error)
	FindAll(limit, offset int) ([]user.User, error)
}

type userService struct {
	repo repository.UserRepoInterface
}

func (u *userService) Save(user2 *user.User) (*user.User, error) {
	user2.Date = time.Now()
	return u.repo.Save(user2)
}

func (u *userService) FindByEmail(email string) (*user.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *userService) FindByDate(limit, offset int, start, end time.Time) ([]user.User, error) {
	newLimit := strconv.Itoa(limit)
	newOffset := strconv.Itoa(offset)
	return u.repo.GetByDate(newLimit, newOffset, start, end)
}

func (u *userService) FindAll(limit, offset int) ([]user.User, error) {
	newLimit := strconv.Itoa(limit)
	newOffset := strconv.Itoa(offset)
	return u.repo.GetAllUsers(newLimit, newOffset)
}

func NewUserService(repo repository.UserRepoInterface) UserServiceInterface {
	return &userService{repo: repo}
}
