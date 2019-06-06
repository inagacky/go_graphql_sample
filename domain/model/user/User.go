package user

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	Id          string `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

// constructor
func NewUser(
	name string,
	email string) (*User, error) {

	// validate
	if name == "" {
		return &User{"", "", ""}, errors.New("name is empty")
	}
	if email == "" {
		return &User{"", "", ""}, errors.New("email is empty")
	}

	return &User{
		Id:   uuid.New().String(),
		Name:     name,
		Email:    email
	}, nil
}

func (user *User) Equals(other User) bool {
	if user.Id == other.Id {
		return true
	}
	return false
}
