package main

import (
	"fmt"
	"math/rand/v2"
)

//объявление структуры для хранения данных аккаунта
type accountData struct {
	login    string
	password string
	url      string
}

func (account *accountData) inputAccountData() {
	account.login = getUserInput("Input login: ")
	account.generatePassword()
	account.url = getUserInput("Input url: ")
}

func (account accountData) printAccountData() {
	fmt.Println("Login:", account.login)
	fmt.Println("Password:", account.password)
	fmt.Println("URL:", account.url)
}

// ф-ция генерации пароля для поля "password" в стуктуре accountData.
func (account *accountData) generatePassword() {
	var symbols = []rune("1abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+")
	generatedPassword := make([]rune, 12) //Длинна пароля
	for i := range generatedPassword {
		generatedPassword[i] = symbols[rand.IntN(len(symbols))]
	}
	account.password = string(generatedPassword)
}

func main() {

	userAccountData := accountData{}

	fmt.Println("Вход в личный кабинет")
	for {
		printMenu()
		var choice int
		fmt.Scan(&choice)
		if choice < 1 || choice > 3 {
			fmt.Println("\nОшибка ввода: выберите пункт меню от 1 до 3")
			continue
		}
		if choice == 1 {
			userAccountData.inputAccountData()
			fmt.Println("\nДанные аккаунта добавлены успешно")
		}
		if choice == 2 {
			fmt.Println("\nПросмотр данных аккаунта")
			userAccountData.printAccountData()
		}
		if choice == 3 {
			fmt.Println("\nВыход из программы")
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

func printMenu() {
	fmt.Println("\nВыберите действие:")
	fmt.Println("1. Добавить данные аккаунта")
	fmt.Println("2. Просмотреть данные аккаунта")
	fmt.Println("3. Выход")
}
