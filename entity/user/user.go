package user

import "time"

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Date      time.Time `json:"date"`
}

type Users struct {
	users []*User
}

func (u *User) Vaulidate() {
	panic("Implement me ")
}

func (u *Users) Add(user *User) {
	u.users = append(u.users, user)
}
