package handler

import (
	"fmt"
	"strconv"
	"todo/repo"
	"todo/service"
)

func Login() (repo.User, error) {
	var login, password string
	fmt.Print("Введите ваш логин:\n>>")
	Reader(&login)
	fmt.Print("Введите пароль:\n>>")
	Reader(&password)
	return service.Login(login, password)
}

func Register() error {
	var login, name, password string

	fmt.Print("Введите логин:\n>>")
	Reader(&login)

	_, err := service.ReadUser(login)
	if err == nil {
		return fmt.Errorf("Пользователь с таким логином уже существует")
	}

	fmt.Print("Ваше Имя:\n>>")
	Reader(&name)
	fmt.Print("Пароль:\n>>")
	Reader(&password)

	err = service.Register(login, password, name)
	if err != nil {
		return err
	}
	fmt.Println("Пользователь успешно создан!")
	return nil
}

func AuthFlow() (repo.User, uint8) {
	for {
		var signIn string
		fmt.Print("1. Войти\n2. Зарегистрироваться\n>> ")
		Reader(&signIn)
		sIn, _ := strconv.Atoi(signIn)
		switch sIn {
		case 1:
			fmt.Println("Вход")
			u, err := Login()
			if err != nil {
				fmt.Println(err)
				continue
			}
			return u, 1
		case 2:
			fmt.Println("Регистрация")
			err := Register()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Регистрация успешна. Выполните вход.")
		case 0:
			fmt.Println("\nВыход из программы...")
			return repo.User{}, 0
		default:
			fmt.Print("\nВведите 1 для Входа или 2 для Регистрации!")
		}
	}
}
