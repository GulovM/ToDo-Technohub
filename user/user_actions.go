package user

import (
	"fmt"
	"strconv"
	"todo/task"
)

func ActionsFlow(u User) {
	for {
		var choice string
		fmt.Println("Выберите действие:")
		fmt.Print("1. Задачи\n2. Настройки\n0. Выйти\n>> ")
		task.Reader(&choice)
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

func TasksChoice(u User) error {
	for {
		var choice string
		task.AllTasks(u.ID)
		fmt.Println("Выберите действие:")
		fmt.Print("1. Создать задачу\n2. Изменить задачу\n3. Удалить задачу\n0. Назад\n>>")
		task.Reader(&choice)
		ch, _ := strconv.Atoi(choice)
		switch ch {
		case 1:
			task.CreateTask(u.ID)
		case 2:
			var taskString string
			fmt.Print("Какую задачу хотите изменить(id задачи)?\n>>")
			task.Reader(&taskString)
			taskId, err := strconv.Atoi(taskString)
			if err != nil {
				return err
			}
			fmt.Println("-------------")
			t, err := task.Read(u.ID, taskId)
			if err != nil {
				return err
			}
			t.Update()
		case 3:
			var (
				taskString string
				isDelete   string
			)

			fmt.Print("Какую задачу хотите удалить(id задачи)?\n>>")
			task.Reader(&taskString)
			taskId, err := strconv.Atoi(taskString)
			if err != nil {
				return err
			}
			t, err := task.Read(u.ID, taskId)
			if err != nil {
				return err
			}
			fmt.Print("Вы уверены, что хотите удалить эту задачу(Y/N)?\n>>")
			task.Reader(&isDelete)
			switch isDelete {
			case "Y", "y", "yes", "Yes", "YES", "Да", "да", "д", "Д":
				t.Delete()
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

func SettingsChoice(u *User) {
	for {
		var mainChoice string
		fmt.Println("Выберите действие:")
		fmt.Print("1. Изменить свои данные аккаунта\n2. Удалить аккаунт\n0. Назад\n>> ")
		task.Reader(&mainChoice)
		mCh, _ := strconv.Atoi(mainChoice)
		switch mCh {
		case 1:
			fmt.Println("Аккаунт\n----------")
			var editChoice string
			var login, name, password string
			fmt.Println("Что хотите изменить?")
			fmt.Print("1. Логин\n2. Имя\n3. Пароль\n4. Все данные\n0. Назад\n>> ")
			task.Reader(&editChoice)
			eCh, _ := strconv.Atoi(editChoice)
			switch eCh {
			case 1:
				//login
				fmt.Print("Введите новый логин:\n>>")
				task.Reader(&login)
				_, err := Read(login)
				if err == nil {
					fmt.Println("Пользователь с таким логином уже существует")
					return
				}

				u.ChangeLogin(login)
			case 2:
				//name
				fmt.Print("Введите новое имя:\n>>")
				task.Reader(&name)
				u.Name = name
				u.Update()
			case 3:
				//password
				fmt.Print("Введите новый пароль:\n>>")
				task.Reader(&password)
				u.Password = password
				u.Update()
			case 4:
				//all
				fmt.Print("Введите новый логин:\n>>")
				task.Reader(&login)
				_, err := Read(login)
				if err == nil {
					fmt.Println("Пользователь с таким логином уже существует")
					return
				}

				u.ChangeLogin(login)

				fmt.Print("Введите новое имя:\n>>")
				task.Reader(&name)
				u.Name = name

				fmt.Print("Введите новый пароль:\n>>")
				task.Reader(&password)
				u.Password = password

				u.Update()
			case 0:
				//back
				break
			default:
				fmt.Print("\nВведите только число 0-4!")
			}
		case 2:
			var isDelete string
			fmt.Print("Вы уверены, что хотите удалить аккаунт(Y/N)?\n>>")
			task.Reader(&isDelete)
			switch isDelete {
			case "Y", "y", "yes", "Yes", "YES", "Да", "да", "д", "Д":
				u.Delete()
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
