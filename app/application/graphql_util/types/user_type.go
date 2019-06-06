package types

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields {
		"id": &graphql.Field {
			Type: graphql.String,
		},
		"firstName": &graphql.Field {
			Type: graphql.String,
		},
		"lastName": &graphql.Field {
			Type: graphql.String,
		},
		"email": &graphql.Field {
			Type: graphql.String,
		},
	},
})

var UserInput = graphql.NewInputObject(graphql.InputObjectConfig {
	Name: "user",
	Fields: graphql.InputObjectConfigFieldMap {
		"id": &graphql.InputObjectFieldConfig {
			Type: graphql.String,
		},
		"firstName": &graphql.InputObjectFieldConfig {
			Type: graphql.String,
		},
		"lastName": &graphql.InputObjectFieldConfig {
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig {
			Type: graphql.String,
		},
	},
	Description: "user input type",
})
