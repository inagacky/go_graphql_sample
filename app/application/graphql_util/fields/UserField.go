package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/application/graphql_util/types"
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
	"github.com/inagacky/go_graphql_sample/app/domain/service"
	userRepo "github.com/inagacky/go_graphql_sample/app/infrastructure/user"
)

// fetch single user
var UserField = &graphql.Field {
	Type:        types.UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument {
		"id": &graphql.ArgumentConfig {
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userId, isOK := params.Args["id"].(int)
		if isOK {
			return service.FindUserById(userId)
		}

		return nil, errors.New("no userId")
	},
}

// fetch all user
var UserListField = &graphql.Field {
	Type:        graphql.NewList(types.UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		userList, _ := userRepo.NewUserRepository().FindUserList()

		return userList, nil
	},
}

// create user
var CreateUserField = &graphql.Field {
	Type:        types.UserType,
	Description: "Create User",
	Args: graphql.FieldConfigArgument {
		"firstName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig {
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		firstName, _ := params.Args["firstName"].(string)
		lastName, _ := params.Args["lastName"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := user.NewUser(firstName, lastName, email)
		if err != nil {
			panic(err)
		}

		userRepo.NewUserRepository().Save(newUser)
		return newUser, nil
	},
}
