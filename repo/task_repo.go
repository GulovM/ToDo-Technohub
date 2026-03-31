package repo

import "fmt"

func CreateTask(title, description, deadline string, userID int) Task {
	var t Task
	nextTaskID++
	UserTaskIDs[userID] += 1
	t.Title = title
	t.Description = description
	t.Status = false
	t.Deadline = deadline
	t.ID = nextTaskID
	t.UserTaskID = UserTaskIDs[userID]
	t.UserID = userID
	Tasks[t.ID] = t
	return t
}

func SaveTask(t Task) {
	Tasks[t.ID] = t
}

func DeleteTask(id int) {
	delete(Tasks, id)
}

func ReadTask(userID, taskID int) (Task, error) {
	for _, v := range Tasks {
		if v.UserTaskID == taskID && v.UserID == userID {
			return v, nil
		}
	}
	return Task{}, fmt.Errorf("Задача не найдена.")
}

func ListUserTasks(userID int) []Task {
	tasks := make([]Task, 0)
	for _, v := range Tasks {
		if v.UserID == userID {
			tasks = append(tasks, v)
		}
	}
	return tasks
}
