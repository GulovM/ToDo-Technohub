package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Tasks = make(map[int]Task)
var ID, UserTaskID = 0, 0

type Task struct {
	Status      bool
	ID          int //общий айди
	UserTaskID  int //айди для задач конкретных пользоватлей
	UserID      int
	Title       string
	Description string
	Deadline    string
}

func Create(title, description, deadline string, UserID int) {
	var t Task
	for _, v := range Tasks {
		if v.UserID == UserID {
			if v.UserTaskID > 0 {
				UserTaskID = v.UserTaskID + 1
			} else {
				UserTaskID += 1
			}
		}
	}
	if UserTaskID == 0 {
		UserTaskID += 1
	}
	ID += 1
	t.Title = title
	t.Description = description
	t.Status = false
	t.Deadline = deadline
	t.ID = ID
	t.UserTaskID = UserTaskID
	t.UserID = UserID
	Tasks[t.ID] = t
}

func CreateTask(UserID int) {
	var title, description, deadline string
	fmt.Print("Заголовок задачи:\n>>")
	fmt.Scan(&title)
	fmt.Print("Описание задачи:\n>>")
	fmt.Scan(&description)
	fmt.Print("Дедлайн задачи:\n>>")
	fmt.Scan(&deadline)
	Create(title, description, deadline, UserID)
	fmt.Println("Задача успешно создана!")
}

func (t *Task) Update() {
	var title, description, deadline string
	var status bool
	fmt.Print("Заголовок задачи(enter - пропустить):\n>>")
	reader(&title)
	fmt.Print("Описание задачи(enter - пропустить):\n>>")
	reader(&description)
	fmt.Print("Дедлайн задачи(enter - пропустить):\n>>")
	reader(&deadline)
	fmt.Print("Статус задачи(true - выполнен, false - не выполнен:\n>>")
	fmt.Scan(&status)
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if deadline != "" {
		t.Deadline = deadline
	}
	t.Status = status
	Tasks[t.ID] = *t
}

func (t *Task) Delete() {
	delete(Tasks, t.ID)
}

func Read(taskID int) Task {
	for _, v := range Tasks {
		if v.UserTaskID == taskID {
			return v
		}
	}
	return Task{}
}

func AllTasks(UserID int) {
	for _, v := range Tasks {
		if v.UserID == UserID {
			fmt.Println("-------------")
			fmt.Printf("№:%v\nЗаголовок:%v\nОписание:%v\nСтатус:%v\nДедлайн:%v\n", v.UserTaskID, v.Title, v.Description, v.Status, v.Deadline)
			fmt.Println("-------------")
		}
	}
}

func reader(input *string) string {
	reader := bufio.NewReader(os.Stdin)
	*input, _ = reader.ReadString('\n')
	*input = strings.TrimSpace(*input)
	return *input
}
