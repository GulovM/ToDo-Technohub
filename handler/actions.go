package handler

import (
	"fmt"
	"strconv"
	"todo/repo"
	"todo/service"
)

func ActionsFlow(u repo.User) {
	for {
		var choice string
		fmt.Println("Выберите действие:")
		fmt.Print("1. Задачи\n2. Настройки\n0. Выйти\n>> ")
		Reader(&choice)
		ch, _ := strconv.Atoi(choice)
		switch ch {
		case 1:
			fmt.Println("Задачи")
			TasksChoice(u)
		case 2:
			fmt.Println("Настройки")
			SettingsChoice(&u)
		case 0:
			fmt.Println("\nВыход из программы...")
			return
		default:
			fmt.Print("\nВведите 1 для работы с Задачами или 2 для Настройки!")
		}
	}
}

func TasksChoice(u repo.User) error {
	for {
		var choice string
		printTasks(u.ID)
		fmt.Println("Выберите действие:")
		fmt.Print("1. Создать задачу\n2. Изменить задачу\n3. Удалить задачу\n0. Назад\n>>")
		Reader(&choice)
		ch, _ := strconv.Atoi(choice)
		switch ch {
		case 1:
			createTask(u.ID)
		case 2:
			var taskString string
			fmt.Print("Какую задачу хотите изменить(id задачи)?\n>>")
			Reader(&taskString)
			taskID, err := strconv.Atoi(taskString)
			if err != nil {
				return err
			}
			fmt.Println("-------------")
			t, err := service.ReadTask(u.ID, taskID)
			if err != nil {
				return err
			}
			updateTask(&t)
		case 3:
			var (
				taskString string
				isDelete   string
			)

			fmt.Print("Какую задачу хотите удалить(id задачи)?\n>>")
			Reader(&taskString)
			taskID, err := strconv.Atoi(taskString)
			if err != nil {
				return err
			}
			t, err := service.ReadTask(u.ID, taskID)
			if err != nil {
				return err
			}
			fmt.Print("Вы уверены, что хотите удалить эту задачу(Y/N)?\n>>")
			Reader(&isDelete)
			switch isDelete {
			case "Y", "y", "yes", "Yes", "YES", "Да", "да", "д", "Д":
				service.DeleteTask(&t)
			case "N", "n", "no", "No", "NO", "Нет", "нет", "н", "Н":
				return nil
			default:
				fmt.Println("Введите Y или N!")
			}
		case 0:
			return nil
		default:
			fmt.Println("Введите 0-3!")
		}
	}
}

func SettingsChoice(u *repo.User) {
	for {
		var mainChoice string
		fmt.Println("Выберите действие:")
		fmt.Print("1. Изменить свои данные аккаунта\n2. Удалить аккаунт\n0. Назад\n>> ")
		Reader(&mainChoice)
		mCh, _ := strconv.Atoi(mainChoice)
		switch mCh {
		case 1:
			fmt.Println("Аккаунт\n----------")
			var editChoice string
			var login, name, password string
			fmt.Println("Что хотите изменить?")
			fmt.Print("1. Логин\n2. Имя\n3. Пароль\n4. Все данные\n0. Назад\n>> ")
			Reader(&editChoice)
			eCh, _ := strconv.Atoi(editChoice)
			switch eCh {
			case 1:
				fmt.Print("Введите новый логин:\n>>")
				Reader(&login)
				_, err := service.ReadUser(login)
				if err == nil {
					fmt.Println("Пользователь с таким логином уже существует")
					return
				}

				service.ChangeLogin(u, login)
			case 2:
				fmt.Print("Введите новое имя:\n>>")
				Reader(&name)
				u.Name = name
				service.UpdateUser(u)
			case 3:
				fmt.Print("Введите новый пароль:\n>>")
				Reader(&password)
				u.Password = password
				service.UpdateUser(u)
			case 4:
				fmt.Print("Введите новый логин:\n>>")
				Reader(&login)
				_, err := service.ReadUser(login)
				if err == nil {
					fmt.Println("Пользователь с таким логином уже существует")
					return
				}

				service.ChangeLogin(u, login)

				fmt.Print("Введите новое имя:\n>>")
				Reader(&name)
				u.Name = name

				fmt.Print("Введите новый пароль:\n>>")
				Reader(&password)
				u.Password = password

				service.UpdateUser(u)
			case 0:
				break
			default:
				fmt.Print("\nВведите только число 0-4!")
			}
		case 2:
			var isDelete string
			fmt.Print("Вы уверены, что хотите удалить аккаунт(Y/N)?\n>>")
			Reader(&isDelete)
			switch isDelete {
			case "Y", "y", "yes", "Yes", "YES", "Да", "да", "д", "Д":
				service.DeleteUser(u)
			case "N", "n", "no", "No", "NO", "Нет", "нет", "н", "Н":
				return
			default:
				fmt.Println("Введите Y или N!")
			}
		case 0:
			fmt.Println("\nНазад...")
			return
		default:
			fmt.Print("\nВведите только 1, либо 2, либо 0!")
		}
	}
}

func createTask(userID int) {
	var title, description, deadline string
	fmt.Print("Заголовок задачи:\n>>")
	Reader(&title)
	fmt.Print("Описание задачи:\n>>")
	Reader(&description)
	fmt.Print("Дедлайн задачи:\n>>")
	Reader(&deadline)
	service.CreateTask(title, description, deadline, userID)
	fmt.Println("Задача успешно создана!")
}

func updateTask(t *repo.Task) {
	var title, description, deadline, status string
	fmt.Print("Заголовок задачи(enter - пропустить):\n>>")
	Reader(&title)
	fmt.Print("Описание задачи(enter - пропустить):\n>>")
	Reader(&description)
	fmt.Print("Дедлайн задачи(enter - пропустить):\n>>")
	Reader(&deadline)
	fmt.Print("Статус задачи(true - выполнен, false - не выполнен:\n>>")
	Reader(&status)
	err := service.UpdateTask(t, title, description, deadline, status)
	if err != nil {
		fmt.Println(err)
	}
}

func printTasks(userID int) {
	for _, v := range service.ListTasks(userID) {
		fmt.Println("-------------")
		fmt.Printf("№:%v\nЗаголовок:%v\nОписание:%v\nСтатус:%v\nДедлайн:%v\n", v.UserTaskID, v.Title, v.Description, v.Status, v.Deadline)
		fmt.Println("-------------")
	}
}
