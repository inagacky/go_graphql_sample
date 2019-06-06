package graphql_util

import (
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/application/graphql_util/fields"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"user":      fields.UserField,
		"userList":  fields.UserListField,
	},
})

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"createUser":  fields.CreateUserField,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query,
	Mutation: mutation,
})
