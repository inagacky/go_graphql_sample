package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/application/graphql_util/types"
	"github.com/inagacky/go_graphql_sample/app/presentation/resolver"
)

var TaskField = &graphql.Field{
	Type:        types.TaskType,
	Description: "Get single task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolver.GetTask,
}

var TaskListField = &graphql.Field{
	Type:        graphql.NewList(types.TaskType),
	Description: "Task of tasks",
	Resolve:     resolver.GetTaskList,
}

var CreateTaskField = &graphql.Field{
	Type:        types.TaskType,
	Description: "Create Task",
	Args: graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolver.CreateTask,
}
