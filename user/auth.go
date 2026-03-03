package user

import (
	"fmt"
)

func Login() (User, error) {
	var login, password string
	fmt.Print("Введите ваш логин:\n>>")
	fmt.Scan(&login)
	fmt.Print("Введите пароль:\n>>")
	fmt.Scan(&password)
	u, err := Read(login)
	if err != nil {
		return User{}, err
	}
	if u.Password != password {
		return User{}, fmt.Errorf("Неправильный пароль!")
	}
	return u, nil
}
func Register() error {
	var login, name, password string
	_, err := Read(login)
	if err == nil {
		return fmt.Errorf("Пользователь с таким логином уже существует")
	}

	fmt.Print("Введите логин:\n>>")
	fmt.Scan(&login)
	fmt.Print("Ваше Имя:\n>>")
	fmt.Scan(&name)
	fmt.Print("Пароль:\n>>")
	fmt.Scan(&password)
	Create(login, password, name)
	fmt.Println("Пользователь успешно создан!")
	return nil
}

func AuthFlow() (User, uint8) {
	for {
		var signIn uint8
		fmt.Print("1. Войти\n2. Зарегистрироваться\n>> ")
		fmt.Scan(&signIn)
		switch signIn {
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
			return User{}, 0
		default:
			fmt.Print("\nВведите 1 для Входа или 2 для Регистрации!")
		}
	}
}
