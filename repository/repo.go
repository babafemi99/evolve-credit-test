package repository

import (
	"evc/entity/user"
	"time"
)

type UserRepoInterface interface {
	GetByEmail(email string) (*user.User, error)
	GetByDate(limit, offset string, start, end time.Time) (*user.Users, error)
	GetAllUsers(limit, offset string) (*user.Users, error)
	GetAllUsersByLimit(limit int) (*user.Users, error)
	GetAllUsersByPage(page int) (*user.Users, error)
}
