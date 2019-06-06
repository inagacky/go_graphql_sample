package user

import (
	"errors"
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
)

type UserRepositoryImpl struct {
	users []*user.User
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{[]*user.User{}}
}

// store event to repository
func (u *UserRepositoryImpl) Store(user *user.User) UserRepository {
	u.users = append(u.users, user)
	return u
}

func (u *UserRepositoryImpl) FindById(id string) (*user.User, error) {
	for _, val := range u.users {
		if val.Id == id {
			return val, nil
		}
	}

	return nil, errors.New("user not found")
}

func (u *UserRepositoryImpl) UserList() []*user.User {
	return u.users
}
