package task

import (
	"github.com/inagacky/go_graphql_sample/app/domain/model/task"
	taskRepo "github.com/inagacky/go_graphql_sample/app/infrastructure/task"
)

func FindTaskById(taskId int) (*task.Task, error) {
	repository := taskRepo.NewTaskRepository()
	return repository.FindById(taskId)
}

func Save(newTask *task.Task) (*task.Task, error) {

	return taskRepo.NewTaskRepository().Save(newTask)
}

func FindTaskList() ([]*task.Task, error) {

	return taskRepo.NewTaskRepository().FindTaskList()
}
