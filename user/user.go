package user

import "fmt"

var Users = make(map[string]User)
var ID = 0

type User struct {
	ID       int
	Login    string
	Password string
	Name     string
}

func Create(login, password, name string) {
	for _, v := range Users {
		if v.ID > 0 {
			ID = v.ID + 1
		} else {
			ID += 1
		}
	}
	var u User = User{
		ID:       ID,
		Login:    login,
		Name:     name,
		Password: password}
	Users[login] = u
}

func (u *User) Update() {
	Users[u.Login] = *u
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
