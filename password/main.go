package main

import (
	"fmt"
)

//объявление структуры для хранения данных аккаунта
type accoutData struct {
	login    string
	password string
	url      string
}

func main() {
	UserAccountData := accoutData{}

	fmt.Println("Вход в личный кабинет")
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить данные аккаунта")
		fmt.Println("2. Просмотреть данные аккаунта")
		fmt.Println("3. Выход")
		var choice int
		fmt.Scan(&choice)
		if choice == 1 {
			getAccountData(&UserAccountData)
			fmt.Println("\nДанные аккаунта добавлены успешно")
		}
		if choice == 2 {
			fmt.Println("\nПросмотр данных аккаунта")
			fmt.Println("Login:", UserAccountData.login)
			fmt.Println("Password:", UserAccountData.password)
			fmt.Println("URL:", UserAccountData.url)
		}
		if choice == 3 {
			fmt.Println("Выход из программы")
			break
		}
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scan(&result)
	return result
}

func getAccountData(account *accoutData) {
	login := getUserInput("Input login: ")
	password := getUserInput("Input password: ")
	url := getUserInput("Input url: ")
	*account = accoutData{
		login:    login,
		password: password,
		url:      url,
	}
}
