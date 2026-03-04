package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Tasks = make(map[int]Task)

var UserTaskIDs = make(map[int]int)
var nextTaskID = 0

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
	nextTaskID++
	UserTaskIDs[UserID] += 1
	t.Title = title
	t.Description = description
	t.Status = false
	t.Deadline = deadline
	t.ID = nextTaskID
	t.UserTaskID = UserTaskIDs[UserID]
	t.UserID = UserID
	Tasks[t.ID] = t
}

func CreateTask(UserID int) {
	var title, description, deadline string
	fmt.Print("Заголовок задачи:\n>>")
	Reader(&title)
	fmt.Print("Описание задачи:\n>>")
	Reader(&description)
	fmt.Print("Дедлайн задачи:\n>>")
	Reader(&deadline)
	Create(title, description, deadline, UserID)
	fmt.Println("Задача успешно создана!")
}

func (t *Task) Update() {
	var title, description, deadline, status string
	fmt.Print("Заголовок задачи(enter - пропустить):\n>>")
	Reader(&title)
	fmt.Print("Описание задачи(enter - пропустить):\n>>")
	Reader(&description)
	fmt.Print("Дедлайн задачи(enter - пропустить):\n>>")
	Reader(&deadline)
	fmt.Print("Статус задачи(true - выполнен, false - не выполнен:\n>>")
	Reader(&status)
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
		fmt.Println(err)
		return
	}
	t.Status = st
	Tasks[t.ID] = *t
}

func (t *Task) Delete() {
	delete(Tasks, t.ID)
}

func Read(userID, taskID int) (Task, error) {
	for _, v := range Tasks {
		if v.UserTaskID == taskID && v.UserID == userID {
			return v, nil
		}
	}
	return Task{}, fmt.Errorf("Задача не найдена.")
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

func Reader(input *string) string {
	reader := bufio.NewReader(os.Stdin)
	*input, _ = reader.ReadString('\n')
	*input = strings.TrimSpace(*input)
	return *input
}
