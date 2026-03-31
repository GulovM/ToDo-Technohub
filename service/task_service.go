package service

import (
	"strconv"
	"todo/repo"
)

func CreateTask(title, description, deadline string, userID int) repo.Task {
	return repo.CreateTask(title, description, deadline, userID)
}

func ReadTask(userID, taskID int) (repo.Task, error) {
	return repo.ReadTask(userID, taskID)
}

func ListTasks(userID int) []repo.Task {
	return repo.ListUserTasks(userID)
}

func UpdateTask(t *repo.Task, title, description, deadline, status string) error {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if deadline != "" {
		t.Deadline = deadline
	}
	st, err := strconv.ParseBool(status)
	if err != nil {
		return err
	}
	t.Status = st
	repo.SaveTask(*t)
	return nil
}

func DeleteTask(t *repo.Task) {
	repo.DeleteTask(t.ID)
}
