package user

import (
	"fmt"
	"github.com/inagacky/go_graphql_sample/app/application/db"
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Save(user *user.User) (*user.User, error)
	FindById(userId int) (*user.User, error)
	FindUserList() ([]*user.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{
	}
}

// ユーザー作成
func (u *userRepository) Save(user *user.User) (*user.User, error) {

	db := db.GetDB()
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("User Save Error: %v", err.Error())
		return nil, err
	}

	return user, nil
}

// IDを元にユーザー取得
func (u *userRepository) FindById(id int) (*user.User, error) {

	var user = user.User{}
	user.Id = id
	db := db.GetDB()
	if err := db.First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("FindById User errors: %v", err.Error())
			return nil, err
		} else {
			fmt.Printf("FindById User Notfound ID: %v", id)
			return nil, nil
		}
	}

	return &user, nil
}

// ユーザー一覧取得
func (u *userRepository) FindUserList() ([]*user.User, error) {

	// 構造体のインスタンス化
	userList := []*user.User{}
	db := db.GetDB()
	if err := db.Find(&userList).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("FindUserList User errors: %v", err.Error())
			return nil, err
		}
	}

	return userList, nil
}
