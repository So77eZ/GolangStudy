package main

import (
	"fmt"

	"GolangCourse/password/account"
	"GolangCourse/password/utils"
)

func main() {

	fakeLogin := utils.GetUserInput("Введите логин: ")
	fakePassword := utils.GetUserInput("Введите пароль: ")
	fakeURL := utils.GetUserInput("Введите URL: ")
	userAccountData, err := account.NewAccountWithTimeStamp(fakeLogin, fakePassword, fakeURL) //accountData{}
	if err != nil {
		fmt.Println(err)
		return
	}

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
			userAccountData.InputAccountData()
			fmt.Println("\nДанные аккаунта добавлены успешно")
		}
		if choice == 2 {
			fmt.Println("\nПросмотр данных аккаунта")
			userAccountData.PrintAccount()
		}
		if choice == 3 {
			fmt.Println("\nВыход из программы")
			break
		}
	}
}

func printMenu() {
	fmt.Println("\nВыберите действие:")
	fmt.Println("1. Добавить данные аккаунта")
	fmt.Println("2. Просмотреть данные аккаунта")
	fmt.Println("3. Выход")
}
