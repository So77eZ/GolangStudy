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
	if err := vault.save(); err != nil {
		color.Cyan("Ошибка при сохранении хранилища vault")
		color.Red(err.Error())
	}
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

//DeleteAccountByURL удаляет аккаунт по переданному URL
func (vault *Vault) DeleteAccountByURL(URL string) bool {
	for i, account := range vault.Accounts {
		if strings.Contains(account.URL, URL) {
			// Удаляем аккаунт из среза
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			if err := vault.save(); err != nil {
				color.Red("Ошибка при сохранении хранилища vault")
			}
			return true
		}
	}
	return false
}

func (vault *Vault) save() error {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		return err
	}
	files.WriteFile(data, "data.json")
	return nil
}

// ToBytes сериализует Vault в JSON-байты для последующей записи в файл.
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
