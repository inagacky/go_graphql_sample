package user

import (
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
)

type UserRepository interface {
	Store(user *user.User) UserRepository
	FindById(userId string) (*user.User, error)
	UserList() []*user.User
}
