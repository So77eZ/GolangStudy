package main

import (
	"fmt"

	"GolangCourse/password/account"
	"GolangCourse/password/utils"
)

func main() {
	// fakeLogin := utils.GetUserInput("Введите логин: ")
	// fakePassword := utils.GetUserInput("Введите пароль: ")
	// fakeURL := utils.GetUserInput("Введите URL: ")
	// userAccountData, err := account.NewAccount(fakeLogin, fakePassword, fakeURL) //accountData{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //Первая запись данных аккаунта
	// files.WriteFile(userAccountData.FormatAccount(), "userAccountData.txt")
	fmt.Println("Вход в личный кабинет")
Menu:
	for {
		choice := printMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			createAccount()
		case 2:
			fmt.Println("\nНахождение аккаунта, on constuction")
			findAccount()
		case 3:
			fmt.Println("\nУдаление аккаунта, on constuction")
			deleteAccount()
			break Menu
		default:
			fmt.Println("\nВыход из программы")
			break Menu
		}
	}
}

func findAccount() {

}

func deleteAccount() {

}

func printMenu() int {
	var userChoice int
	fmt.Println("\nВыберите действие:")
	fmt.Println("\n1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	return userChoice
}

func createAccount() {
	Login := utils.GetUserInput("Введите логин: ")
	Password := utils.GetUserInput("Введите пароль: ")
	URL := utils.GetUserInput("Введите URL: ")
	myAccount, err := account.NewAccount(Login, Password, URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	// data, err := vault.ToBytes()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// files.WriteFile(data, "data.json")
	fmt.Println("\nАккаунт успешно создан")
}
