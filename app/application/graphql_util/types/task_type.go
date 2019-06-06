package types

import (
	"github.com/graphql-go/graphql"
)

var TaskType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields {
		"id": &graphql.Field {
			Type: graphql.Int,
		},
		"userId": &graphql.Field {
			Type: graphql.Int,
		},
		"title": &graphql.Field {
			Type: graphql.String,
		},
		"content": &graphql.Field {
			Type: graphql.String,
		},
	},
})
