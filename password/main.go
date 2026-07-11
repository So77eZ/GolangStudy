package main

import (
	"fmt"
	"math/rand/v2"
)

//объявление структуры для хранения данных аккаунта
type accoutData struct {
	login    string
	password string
	url      string
}

var symbols = []rune("1abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+")

func main() {

	rand.IntN(100)

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
			printAccountData(&UserAccountData)
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
	password := generatePassword(12) //getUserInput("Input password: ")
	url := getUserInput("Input url: ")
	*account = accoutData{
		login:    login,
		password: password,
		url:      url,
	}
}

func printAccountData(account *accoutData) {
	fmt.Println("Login:", account.login)
	fmt.Println("Password:", account.password)
	fmt.Println("URL:", account.url)
}

func generatePassword(length int) string {
	password := make([]rune, length)
	for i := range password {
		password[i] = symbols[rand.IntN(len(symbols))]
	}
	return string(password)
}
