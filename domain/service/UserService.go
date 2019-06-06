package service

import (
	"github.com/inagacky/go_graphql_sample/domain/model/user"
	"github.com/inagacky/go_graphql_sample/infrastructure"
)

func FindUserById(userId string) (*user.User, error) {
	repository := infrastructure.NewUserRepository()
	return repository.FindById(userId)
}
