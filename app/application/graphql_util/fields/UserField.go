package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/application/graphql_util/types"
	"github.com/inagacky/go_graphql_sample/app/domain/model/user"
	"github.com/inagacky/go_graphql_sample/app/domain/service"
	"github.com/inagacky/go_graphql_sample/app/infrastructure"
)

// fetch single user
var UserField = &graphql.Field {
	Type:        types.UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument {
		"id": &graphql.ArgumentConfig {
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userId, isOK := params.Args["id"].(string)
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
		return infrastructure.NewUserRepository().UserList(), nil
	},
}

// create user
var CreateUserField = &graphql.Field {
	Type:        types.UserType,
	Description: "Create User",
	Args: graphql.FieldConfigArgument {
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig {
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userName, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := user.NewUser(userName, email)
		if err != nil {
			panic(err)
		}
		infrastructure.NewUserRepository().Store(newUser)
		return newUser, nil
	},
}
