package task

import (
	"errors"
)

type Task struct {
	Id            int `json:"id"`
	UserId        int `json:"user_id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
}

// constructor
func NewTask(
	userId int,
	title string,
	content string) (*Task, error) {

	// validate
	if userId == 0 {
		return &Task{0, 0, "", ""}, errors.New("userId is empty")
	}
	if title == "" {
		return &Task{0, 0, "", ""}, errors.New("title is empty")
	}
	if content == "" {
		return &Task{0, 0, "", ""}, errors.New("content is empty")
	}

	return &Task{
		UserId: userId,
		Title:  title,
		Content: content,
	}, nil
}
