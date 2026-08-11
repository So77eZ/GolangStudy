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

//VaultWithdb хранилище аккаунтов с привязкой к файлу data.json
type VaultWithdb struct {
	Vault
	db *files.JSONdb
}

// NewVault создаёт Vault: пытается загрузить данные из data.json,
// при ошибке чтения или декодирования возвращает пустое хранилище.
func NewVault(db *files.JSONdb) *VaultWithdb {
	// Чтение существующих значений файла data.json
	file, err := db.Read()
	if err != nil {
		return &VaultWithdb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	// Декодирование содержимого файла в структуру Vault
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Cyan("Ошибка при декодировании data.json")
		color.Red(err.Error())
		return &VaultWithdb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithdb{
		Vault: vault,
		db:    db,
	}
}

// AddAccount добавляет acc в список аккаунтов, обновляет UpdatedAt
// и сохраняет хранилище на диск в data.json.
func (vault *VaultWithdb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	if err := vault.save(); err != nil {
		color.Cyan("Ошибка при сохранении хранилища vault")
		color.Red(err.Error())
	}
}

//FindAccountByURL поиск аккаунта по переданному URL
func (vault *VaultWithdb) FindAccountByURL(URL string) []Account {
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
func (vault *VaultWithdb) DeleteAccountByURL(URL string) bool {
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

func (vault *VaultWithdb) save() error {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		return err
	}
	return vault.db.Write(data)
}

// ToBytes сериализует Vault в JSON-байты для последующей записи в файл.
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
