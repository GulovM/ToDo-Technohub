package main

import (
	"fmt"
	"todo/user"
)

func main() {
	for {
		fmt.Println("Добро пожаловать в трекер задач!\n0 - выйти из программы")
		u, i := user.AuthFlow()
		if i == 0 {
			return
		}
		user.ActionsFlow(u)
	}
}
