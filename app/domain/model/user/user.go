package user

import (
	"errors"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// constructor
func NewUser(
	firstName string,
	lastName string,
	email string) (*User, error) {

	// validate
	if firstName == "" {
		return &User{0, "", "", ""}, errors.New("first_name is empty")
	}

	if lastName == "" {
		return &User{0, "", "", ""}, errors.New("last_name is empty")
	}
	if email == "" {
		return &User{0, "", "", ""}, errors.New("email is empty")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}, nil
}

func (user *User) Equals(other User) bool {
	if user.Id == other.Id {
		return true
	}
	return false
}
