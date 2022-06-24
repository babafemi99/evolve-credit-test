package repository

import (
	"evc/entity/user"
)

type UserRepoInterface interface {
	Save(user2 *user.User) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	GetByDate(limit, offset string, start, end string) ([]user.User, error)
	GetAllUsers(limit, offset string) ([]user.User, error)
}
