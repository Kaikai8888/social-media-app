package domain

type User struct {
	Email    string
	Password string
}

func (u *User) ValidateEmail() bool {
	return true
}

func (u *User) ValidatePassword() bool {
	return true
}
