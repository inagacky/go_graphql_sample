package resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
	"github.com/inagacky/go_graphql_sample/app/domain/service"
)

// ユーザー取得
func GetUser(params graphql.ResolveParams) (interface{}, error) {
	userId, isOK := params.Args["id"].(int)
	if isOK {
		return service.FindUserById(userId)
	}

	return nil, errors.New("no userId")
}

// ユーザー一覧取得
func GetUserList(p graphql.ResolveParams) (interface{}, error) {

	return service.FindUserList()
}

// ユーザー作成
func CreateUser (params graphql.ResolveParams) (interface{}, error) {
	firstName, _ := params.Args["firstName"].(string)
	lastName, _ := params.Args["lastName"].(string)
	email, _ := params.Args["email"].(string)

	newUser, err := user.NewUser(firstName, lastName, email)
	if err != nil {
		panic(err)
	}

	service.Save(newUser)

	return newUser, nil
}
