package user

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	email     string `json:"email"`
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
