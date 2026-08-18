package main

import (
	"fmt"
	"os"
	"strings"

	"GolangCourse/password/account"
	"GolangCourse/password/output"

	//"GolangCourse/password/cloude"
	"GolangCourse/password/files"
	"GolangCourse/password/utils"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithdb){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
	"5": exitFromMenu,
}

func main() {
	vault := account.NewVault(files.NewJSONdb("data.json"))
	err := godotenv.Load("../.env")
	if err != nil {
		output.PrintError("Не прочитался файл окружения: ")
		output.PrintError(err.Error())
	}
	res := os.Getenv("VAR")
	fmt.Println("env.: ", res)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], " = ", pair[1])
	}
	// или
	//vault := account.NewVault(cloude.NewCloudeDb("https://exampleCloudStorage.com"))
	fmt.Println("Вход в личный кабинет")
Menu:
	for {
		variant := utils.GetUserInput(
			" ",
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по Login",
			"4. Удалить аккаунт",
			"5. Выход",
			" ",
			"Выбранный вариант",
		)
		menuFunc := menu[variant]
		if menuFunc == nil {
			output.PrintError("Пожалуйста, введите корректное число для выбора пункта меню.")
			break Menu
		}
		menuFunc(vault)

		// _, err := fmt.Scanln(&variant)
		// if err != nil {
		// 	output.PrintError("Пожалуйста, введите корректное число для выбора пункта меню.")
		// 	continue
		// }
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

// findAccountByURL ф-ция нахождения аккаунта по URl
func findAccountByURL(vault *account.VaultWithdb) {
	fmt.Println("\nНахождение аккаунта по URL")
	userURLInput := utils.GetUserInput("Введите URL-ссылку на аккаунт")
	// анонимная ф-ция нахождения аккаунта по переданному URL
	foundedAccounts := vault.FindAccounts(userURLInput, func(acc account.Account, str string) bool {
		return strings.Contains(acc.URL, str)
	})
	outputResults(&foundedAccounts)
}

func outputResults(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		fmt.Println("")
		account.PrintAccount()
	}
}

// findAccountByLogin ф-ция нахождения аккаунта по URl
func findAccountByLogin(vault *account.VaultWithdb) {
	fmt.Println("\nНахождение аккаунта по логину")
	userLoginInput := utils.GetUserInput("Введите Login аккаунта")
	// анонимная ф-ция нахождения аккаунта по переданному URL
	foundedAccounts := vault.FindAccounts(userLoginInput, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResults(&foundedAccounts)
}

// checkLogin ф-ция нахождения аккаунта по переданному логину
func checkLogin(acc account.Account, str string) bool {
	return strings.Contains(acc.Login, str)
}

func deleteAccount(vault *account.VaultWithdb) {
	fmt.Println("\nУдаление аккаунта")
	userURLInput := utils.GetUserInput("Введите URL-ссылку на аккаунт")
	if vault.DeleteAccountByURL(userURLInput) {
		color.Green("Аккаунт по URL: " + userURLInput + " Успешно удален")
	} else {
		output.PrintError("Аккаунт по URL: " + userURLInput + " не был найден")
	}
}

func createAccount(vault *account.VaultWithdb) {
	Login := utils.GetUserInput("Введите логин: ")
	Password := utils.GetUserInput("Введите пароль: ")
	URL := utils.GetUserInput("Введите URL: ")
	myAccount, err := account.NewAccount(Login, Password, URL)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}
	vault.AddAccount(*myAccount)

	fmt.Println("\nАккаунт успешно создан")
}

func exitFromMenu(vault *account.VaultWithdb) {
	color.HiGreen("Выход из меню")
}
