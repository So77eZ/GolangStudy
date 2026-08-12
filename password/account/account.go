package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"GolangCourse/password/utils"

	"github.com/fatih/color"
)

var symbols = []rune("1abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+")

//Account объявление структуры для хранения данных аккаунта
type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

//InputAccountData Ввод данных аккаунта
func (account *Account) InputAccountData() {
	account.Login = utils.GetUserInput([]string{"Input login: "})
	account.generatePassword()
	account.URL = utils.GetUserInput([]string{"Input url: "})
}

//FormatAccount - собрать данные аккаунта в текстовое представление
func (account Account) FormatAccount() string {
	return fmt.Sprintf("Login: %s\nPassword: %s\nURL: %s\nCreated at: %s\n",
		account.Login,
		account.Password,
		account.URL,
		account.CreatedAt)
}

//PrintAccount - вывести данные аккаунта в консоль с цветом
func (account *Account) PrintAccount() {

	color.Blue("Login: " + account.Login)
	color.Cyan("Password: " + account.Password)
	color.Green("URL: " + account.URL)
	color.Yellow("Created at: " + string(account.CreatedAt.Format("02.01.2006")))
}

// ф-ция генерации пароля для поля "password" в стуктуре Account.
func (account *Account) generatePassword() {
	generatedPassword := make([]rune, 12) //Длинна пароля
	for i := range generatedPassword {
		generatedPassword[i] = symbols[rand.IntN(len(symbols))]
	}
	account.Password = string(generatedPassword)
}

//NewAccount ф-ция создания аккаунта со временем указания создания и обновления
func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		URL:       urlString,
		Login:     login,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		fmt.Println("Пароль не был введен, генерируем случайный")
		newAcc.generatePassword()
		//newAcc.accountData.generatePassword()
	}
	return newAcc, nil
}

func generatePassword() []rune {
	generatedPassword := make([]rune, 12) //Длинна пароля
	for i := range generatedPassword {
		generatedPassword[i] = symbols[rand.IntN(len(symbols))]
	}
	return generatedPassword
}
