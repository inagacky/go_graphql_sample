package user

import (
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
	userRepo "github.com/inagacky/go_graphql_sample/app/infrastructure/user"
)

func FindUserById(userId int) (*user.User, error) {
	repository := userRepo.NewUserRepository()
	return repository.FindById(userId)
}

func Save(newUser *user.User) (*user.User, error) {

	return userRepo.NewUserRepository().Save(newUser)
}

func FindUserList() ([]*user.User, error) {

	return userRepo.NewUserRepository().FindUserList()
}
