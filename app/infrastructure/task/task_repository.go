package user

import (
	"fmt"
	"github.com/inagacky/go_graphql_sample/app/application/db"
	"github.com/inagacky/go_graphql_sample/app/domain/model/task"
	"github.com/jinzhu/gorm"
)

type TaskRepository interface {
	Save(task *task.Task) (*task.Task, error)
	FindById(taskId int) (*task.Task, error)
	FindTaskList() ([]*task.Task, error)
}

type taskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

// タスク作成
func (u *taskRepository) Save(task *task.Task) (*task.Task, error) {

	db := db.GetDB()
	if err := db.Create(&task).Error; err != nil {
		fmt.Printf("Task Save Error: %v", err.Error())
		return nil, err
	}

	return task, nil
}

// IDを元にタスク取得
func (u *taskRepository) FindById(id int) (*task.Task, error) {

	var task = task.Task{}
	task.Id = id
	db := db.GetDB()
	if err := db.First(&task).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("FindById Task errors: %v", err.Error())
			return nil, err
		} else {
			fmt.Printf("FindById Task Notfound ID: %v", id)
			return nil, nil
		}
	}

	return &task, nil
}

// タスク一覧取得
func (u *taskRepository) FindTaskList() ([]*task.Task, error) {

	// 構造体のインスタンス化
	taskList := []*task.Task{}
	db := db.GetDB()
	if err := db.Find(&taskList).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Printf("FindUserList Task errors: %v", err.Error())
			return nil, err
		}
	}

	return taskList, nil
}
