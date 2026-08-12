package main

import (
	"fmt"

	"GolangCourse/password/account"
	"GolangCourse/password/output"

	//"GolangCourse/password/cloude"
	"GolangCourse/password/files"
	"GolangCourse/password/utils"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJSONdb("data.json"))
	// или
	//vault := account.NewVault(cloude.NewCloudeDb("https://exampleCloudStorage.com"))
	fmt.Println("Вход в личный кабинет")
Menu:
	for {
		variant := utils.GetUserInput([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		_, err := fmt.Scanln(&variant)
		if err != nil {
			output.PrintError("Пожалуйста, введите корректное число для выбора пункта меню.")
			continue
		}
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			fmt.Println("\nНахождение аккаунта")
			userURLInput := utils.GetUserInput([]string{"Введите URL-ссылку на аккаунт"})
			findAccount(userURLInput, vault)
		case "3":
			fmt.Println("\nУдаление аккаунта")
			userURLInput := utils.GetUserInput([]string{"Введите URL-ссылку на аккаунт"})
			deleteAccount(userURLInput, vault)
		case "4":
			fmt.Println("\nВыход из программы")
			break Menu
		default:
			fmt.Printf("Введено неверное значение номера пункта меню: %s\n", variant)
		}
	}
}

func findAccount(URL string, vault *account.VaultWithdb) {
	foundedAccounts := vault.FindAccountByURL(URL)
	if len(foundedAccounts) == 0 {
		output.PrintError("Аккаунты с таким URL не найдены")
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
		output.PrintError("Аккаунт по URL: " + URL + " не был найден")
	}
}

func createAccount(vault *account.VaultWithdb) {
	Login := utils.GetUserInput([]string{"Введите логин: "})
	Password := utils.GetUserInput([]string{"Введите пароль: "})
	URL := utils.GetUserInput([]string{"Введите URL: "})
	myAccount, err := account.NewAccount(Login, Password, URL)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}
	vault.AddAccount(*myAccount)

	fmt.Println("\nАккаунт успешно создан")
}
