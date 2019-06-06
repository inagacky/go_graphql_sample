package service

import (
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
	userRepo "github.com/inagacky/go_graphql_sample/app/infrastructure/user"
)

func FindUserById(userId int) (*user.User, error) {
	repository := userRepo.NewUserRepository()
	return repository.FindById(userId)
}
