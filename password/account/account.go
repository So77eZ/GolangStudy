package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"GolangCourse/password/utils"
)

var symbols = []rune("1abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+")

//Account объявление структуры для хранения данных аккаунта
type Account struct {
	login    string
	password string
	url      string
}

//AccountWithTimeStamp Аккаунт с учетом времени создания и обновления
type AccountWithTimeStamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

//InputAccountData Ввод данных аккаунта
func (account *AccountWithTimeStamp) InputAccountData() {
	account.login = utils.GetUserInput("Input login: ")
	account.generatePassword()
	account.url = utils.GetUserInput("Input url: ")
}

//PrintAccount Вывод данных об аккаунте
func (account AccountWithTimeStamp) PrintAccount() {
	fmt.Println("Login:", account.login)
	fmt.Println("Password:", account.password)
	fmt.Println("URL:", account.url)
	fmt.Println("Created at: ", account.createdAt.Weekday())
}

// ф-ция генерации пароля для поля "password" в стуктуре Account.
func (account *AccountWithTimeStamp) generatePassword() {
	generatedPassword := make([]rune, 12) //Длинна пароля
	for i := range generatedPassword {
		generatedPassword[i] = symbols[rand.IntN(len(symbols))]
	}
	account.password = string(generatedPassword)
}

//NewAccountWithTimeStamp ф-ция создания аккаунта со временем указания создания и обновления
func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
		// panic("Пустой логин")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimeStamp{
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
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
