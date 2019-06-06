package infrastructure

import (
	"github.com/inagacky/go_graphql_sample/infrastructure/user"
)

var NewUserRepository func() user.UserRepository = user.NewUserRepositoryImpl

func UseUserRepositoryImpl() {
	NewUserRepository = user.NewUserRepositoryImpl
}
