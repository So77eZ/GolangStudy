package main

import (
	"fmt"

	"GolangCourse/password/account"
	"GolangCourse/password/files"
	"GolangCourse/password/utils"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJSONdb("data.json"))
	var choise int //Пользователських ввод пункта меню
	fmt.Println("Вход в личный кабинет")
Menu:
	for {
		printMenu()
		_, err := fmt.Scanln(&choise)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		switch choise {
		case 1:
			createAccount(vault)
		case 2:
			fmt.Println("\nНахождение аккаунта")
			userURLInput := utils.GetUserInput("Введите URL-ссылку на аккаунт")
			findAccount(userURLInput, vault)
		case 3:
			fmt.Println("\nУдаление аккаунта")
			userURLInput := utils.GetUserInput("Введите URL-ссылку на аккаунт")
			deleteAccount(userURLInput, vault)
		case 4:
			fmt.Println("\nВыход из программы")
			break Menu
		default:
			fmt.Printf("Введено неверное значение номера пункта меню: %d\n", choise)
		}
	}
}

func findAccount(URL string, vault *account.VaultWithdb) {
	foundedAccounts := vault.FindAccountByURL(URL)
	if len(foundedAccounts) == 0 {
		color.Red("Аккаунты с таким URL не найдены")
	}
	for _, account := range foundedAccounts {
		fmt.Println("")
		account.PrintAccount()
	}
}

func deleteAccount(URL string, vault *account.VaultWithdb) {
	if vault.DeleteAccountByURL(URL) {
		color.Green("Аккаунт по URL: " + URL + " Успешно удален")
	} else {
		color.Red("Аккаунт по URL: " + URL + " не был найден")
	}
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

func createAccount(vault *account.VaultWithdb) {
	Login := utils.GetUserInput("Введите логин: ")
	Password := utils.GetUserInput("Введите пароль: ")
	URL := utils.GetUserInput("Введите URL: ")
	myAccount, err := account.NewAccount(Login, Password, URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	//vault = account.NewVault(files.NewJSONdb("data.json"))
	vault.AddAccount(*myAccount)

	fmt.Println("\nАккаунт успешно создан")
}
