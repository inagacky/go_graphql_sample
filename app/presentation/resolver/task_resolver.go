package resolver

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/domain/model/task"
	service "github.com/inagacky/go_graphql_sample/app/domain/service/task"
)

// GetTask タスク取得
func GetTask(params graphql.ResolveParams) (interface{}, error) {
	userId, isOK := params.Args["id"].(int)
	if isOK {
		return service.FindTaskById(userId)
	}

	return nil, errors.New("no userId")
}

// GetTaskList タスク一覧取得
func GetTaskList(p graphql.ResolveParams) (interface{}, error) {

	return service.FindTaskList()
}

// CreateTask タスク作成
func CreateTask(params graphql.ResolveParams) (interface{}, error) {
	userId, _ := params.Args["userId"].(int)
	title, _ := params.Args["title"].(string)
	content, _ := params.Args["content"].(string)

	newUser, err := task.NewTask(userId, title, content)
	if err != nil {
		panic(err)
	}

	service.Save(newUser)

	return newUser, nil
}
