package user

import (
	"fmt"
	"todo/task"
)

var Users = make(map[string]User)
var nextUserID = 0

type User struct {
	ID       int
	Login    string
	Password string
	Name     string
}

func Create(login, password, name string) {
	nextUserID++
	var u User = User{
		ID:       nextUserID,
		Login:    login,
		Name:     name,
		Password: password}
	Users[login] = u
	task.UserTaskIDs[u.ID] = 0
}

func (u *User) Update() {
	Users[u.Login] = *u
}

func (u *User) ChangeLogin(newLogin string) {
	oldLogin := u.Login
	u.Login = newLogin
	u.Update()
	delete(Users, oldLogin)
}

func (u *User) Delete() {
	delete(Users, u.Login)
}

func Read(login string) (User, error) {
	_, ok := Users[login]
	if !ok {
		return User{}, fmt.Errorf("Пользователь с таким логином не существует")
	}
	return Users[login], nil
}
