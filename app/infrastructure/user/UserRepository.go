package user

import (
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
)

type UserRepository interface {
	Save(user *user.User) (*user.User, error)
	FindById(userId int) (*user.User, error)
	FindUserList() ([]*user.User, error)
}
