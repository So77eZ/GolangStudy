package main

import (
	"flag"
	"fmt"
	"strings"

	"GolangCourse/password/account"
	"GolangCourse/password/encrypter"
	"GolangCourse/password/output"

	//"GolangCourse/password/cloude"
	"GolangCourse/password/files"
	"GolangCourse/password/utils"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// menuAction выполняет пункт меню и возвращает false, если нужно выйти из меню.
type menuAction func(*account.VaultWithdb) bool

var menu = map[string]menuAction{
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
	"5": exitFromMenu,
}

// loadEnv подхватывает .env, если он лежит в текущей директории или на уровень выше.
// Отсутствие файла не фатально: KEY может быть задан прямо в переменных окружения.
func loadEnv() {
	for _, path := range []string{".env", "../.env"} {
		if err := godotenv.Load(path); err == nil {
			return
		}
	}
	output.PrintError("Файл .env не найден, KEY ожидается в переменных окружения")
}

func main() {
	vaultPath := flag.String("vault", "data.vault", "путь к файлу хранилища паролей")
	flag.Parse()

	loadEnv()
	// NewEncrypter паникует при отсутствии KEY — до того, как будет затронуто хранилище
	enc := encrypter.NewEncrypter()

	vault := account.NewVault(files.NewJSONdb(*vaultPath), *enc)
	// или
	//vault := account.NewVault(cloude.NewCloudeDb("https://exampleCloudStorage.com"))
	fmt.Println("Вход в личный кабинет")
Menu:
	for {
		variant, ok := utils.GetUserInput(
			" ",
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по Login",
			"4. Удалить аккаунт",
			"5. Выход",
			" ",
			"Выбранный вариант",
		)
		// ввод закончился (EOF) — дальше спрашивать нечего
		if !ok {
			break Menu
		}
		action := menu[variant]
		if action == nil {
			output.PrintError("Пожалуйста, введите корректное число для выбора пункта меню.")
			continue
		}
		// пункт меню вернул false — выходим из цикла
		if !action(vault) {
			break Menu
		}

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
func findAccountByURL(vault *account.VaultWithdb) bool {
	fmt.Println("\nНахождение аккаунта по URL")
	userURLInput, ok := utils.GetUserInput("Введите URL-ссылку на аккаунт")
	if !ok {
		return false
	}
	// анонимная ф-ция нахождения аккаунта по переданному URL
	foundedAccounts := vault.FindAccounts(userURLInput, func(acc account.Account, str string) bool {
		return strings.Contains(acc.URL, str)
	})
	outputResults(&foundedAccounts)
	return true
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
func findAccountByLogin(vault *account.VaultWithdb) bool {
	fmt.Println("\nНахождение аккаунта по логину")
	userLoginInput, ok := utils.GetUserInput("Введите Login аккаунта")
	if !ok {
		return false
	}
	// анонимная ф-ция нахождения аккаунта по переданному URL
	foundedAccounts := vault.FindAccounts(userLoginInput, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResults(&foundedAccounts)
	return true
}

// checkLogin ф-ция нахождения аккаунта по переданному логину
func checkLogin(acc account.Account, str string) bool {
	return strings.Contains(acc.Login, str)
}

func deleteAccount(vault *account.VaultWithdb) bool {
	fmt.Println("\nУдаление аккаунта")
	userURLInput, ok := utils.GetUserInput("Введите URL-ссылку на аккаунт")
	if !ok {
		return false
	}
	if vault.DeleteAccountByURL(userURLInput) {
		color.Green("Аккаунт по URL: " + userURLInput + " Успешно удален")
	} else {
		output.PrintError("Аккаунт по URL: " + userURLInput + " не был найден")
	}
	return true
}

func createAccount(vault *account.VaultWithdb) bool {
	Login, ok := utils.GetUserInput("Введите логин: ")
	if !ok {
		return false
	}
	Password, ok := utils.GetUserInput("Введите пароль: ")
	if !ok {
		return false
	}
	URL, ok := utils.GetUserInput("Введите URL: ")
	if !ok {
		return false
	}
	myAccount, err := account.NewAccount(Login, Password, URL)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return true
	}
	vault.AddAccount(*myAccount)

	fmt.Println("\nАккаунт успешно создан")
	return true
}

// exitFromMenu сообщает циклу меню, что работу нужно завершить.
func exitFromMenu(vault *account.VaultWithdb) bool {
	color.HiGreen("Выход из меню")
	return false
}
