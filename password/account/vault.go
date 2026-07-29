package account

import (
	"GolangCourse/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

//Vault хранилище аккаунтов
type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewVault создаёт Vault: пытается загрузить данные из data.json,
// при ошибке чтения или декодирования возвращает пустое хранилище.
func NewVault() *Vault {
	// Чтение существующих значений файла data.json
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	// Декодирование содержимого файла в структуру Vault
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Cyan("Ошибка при декодировании data.json")
		color.Red(err.Error())
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

// AddAccount добавляет acc в список аккаунтов, обновляет UpdatedAt
// и сохраняет хранилище на диск в data.json.
func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Cyan("Ошибка при преобразовании хранилища vault")
		color.Red(err.Error())
		return
	}
	files.WriteFile(data, "data.json")
}

//FindAccountByURL поиск аккаунта по переданному URL
func (vault *Vault) FindAccountByURL(URL string) []Account {
	var foundedAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, URL)
		if isMatched {
			foundedAccounts = append(foundedAccounts, account)
		}
	}
	return foundedAccounts
}

// ToBytes сериализует Vault в JSON-байты для последующей записи в файл.
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
