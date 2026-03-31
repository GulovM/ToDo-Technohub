package service

import (
	"fmt"
	"todo/repo"
)

func Login(login, password string) (repo.User, error) {
	u, err := repo.ReadUser(login)
	if err != nil {
		return repo.User{}, err
	}
	if u.Password != password {
		return repo.User{}, fmt.Errorf("Неправильный пароль!")
	}
	return u, nil
}

func Register(login, password, name string) error {
	_, err := repo.ReadUser(login)
	if err == nil {
		return fmt.Errorf("Пользователь с таким логином уже существует")
	}

	repo.CreateUser(login, password, name)
	return nil
}

func ReadUser(login string) (repo.User, error) {
	return repo.ReadUser(login)
}

func UpdateUser(u *repo.User) {
	repo.SaveUser(*u)
}

func ChangeLogin(u *repo.User, newLogin string) {
	oldLogin := u.Login
	u.Login = newLogin
	UpdateUser(u)
	repo.DeleteUser(oldLogin)
}

func DeleteUser(u *repo.User) {
	repo.DeleteUser(u.Login)
}
