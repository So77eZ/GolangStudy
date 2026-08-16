package main

import (
	"fmt"
	"strings"

	"GolangCourse/password/account"
	"GolangCourse/password/output"

	//"GolangCourse/password/cloude"
	"GolangCourse/password/files"
	"GolangCourse/password/utils"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithdb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

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
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		_, err := fmt.Scanln(&variant)
		if err != nil {
			output.PrintError("Пожалуйста, введите корректное число для выбора пункта меню.")
			continue
		}
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// case "4":
		// 	fmt.Println("\nВыход из программы")
		// 	break Menu
		// default:
		// 	fmt.Printf("Введено неверное значение номера пункта меню: %s\n", variant)
		// }
	}
}

// findAccount ф-ция нахождения аккаунта по переданным параметрам
func findAccount(vault *account.VaultWithdb) {
	fmt.Println("\nНахождение аккаунта")
	userURLInput := utils.GetUserInput([]string{"Введите URL-ссылку на аккаунт"})
	// анонимная ф-ция нахождения аккаунта по переданному URL
	foundedAccounts := vault.FindAccounts(userURLInput, func(acc account.Account, str string) bool {
		return strings.Contains(acc.URL, str)
	})
	if len(foundedAccounts) == 0 {
		output.PrintError("Аккаунты с таким URL не найдены")
	}
	for _, account := range foundedAccounts {
		fmt.Println("")
		account.PrintAccount()
	}
}

// checkLogin ф-ция нахождения аккаунта по переданному логину
func checkLogin(acc account.Account, str string) bool {
	return strings.Contains(acc.Login, str)
}

func deleteAccount(vault *account.VaultWithdb) {
	fmt.Println("\nУдаление аккаунта")
	userURLInput := utils.GetUserInput([]string{"Введите URL-ссылку на аккаунт"})
	if vault.DeleteAccountByURL(userURLInput) {
		color.Green("Аккаунт по URL: " + userURLInput + " Успешно удален")
	} else {
		output.PrintError("Аккаунт по URL: " + userURLInput + " не был найден")
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
