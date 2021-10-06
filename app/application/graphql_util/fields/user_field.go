package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/application/graphql_util/types"
	"github.com/inagacky/go_graphql_sample/app/presentation/resolver"
)

var UserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolver.GetUser,
}

var UserListField = &graphql.Field{
	Type:        graphql.NewList(types.UserType),
	Description: "List of users",
	Resolve:     resolver.GetUserList,
}

var CreateUserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Create User",
	Args: graphql.FieldConfigArgument{
		"firstName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolver.CreateUser,
}
