package repo

import "fmt"

func CreateUser(login, password, name string) User {
	nextUserID++
	u := User{
		ID:       nextUserID,
		Login:    login,
		Name:     name,
		Password: password,
	}
	Users[login] = u
	UserTaskIDs[u.ID] = 0
	return u
}

func SaveUser(u User) {
	Users[u.Login] = u
}

func DeleteUser(login string) {
	delete(Users, login)
}

func ReadUser(login string) (User, error) {
	_, ok := Users[login]
	if !ok {
		return User{}, fmt.Errorf("Пользователь с таким логином не существует")
	}
	return Users[login], nil
}
