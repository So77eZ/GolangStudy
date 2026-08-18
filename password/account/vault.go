package account

import (
	"GolangCourse/password/encrypter"
	"GolangCourse/password/output"
	"encoding/json"
	"strings"
	"time"
)

//Vault хранилище аккаунтов
type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

//VaultWithdb хранилище аккаунтов с привязкой к файлу data.json
type VaultWithdb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

// Db интерфейс для работы с хранилищем данных (например, JSON-файл или облачное хранилище)
type Db interface {
	ByteReader
	ByteWriter
}

// ByteReader интерфейс для чтения данных в виде байтов
type ByteReader interface {
	Read() ([]byte, error)
}

// ByteWriter интерфейс для записи данных в виде байтов
type ByteWriter interface {
	Write([]byte) error
}

// NewVault создаёт Vault: пытается загрузить данные из data.json,
// при ошибке чтения или декодирования возвращает пустое хранилище.
func NewVault(db Db, enc encrypter.Encrypter) *VaultWithdb {
	// Чтение существующих значений файла data.json
	file, err := db.Read()
	if err != nil {
		return &VaultWithdb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	// Декодирование содержимого файла в структуру Vault
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError("Ошибка при декодировании data.json")
		output.PrintError(err.Error())
		return &VaultWithdb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithdb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

// AddAccount добавляет acc в список аккаунтов, обновляет UpdatedAt
// и сохраняет хранилище на диск в data.json.
func (vault *VaultWithdb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	if err := vault.save(); err != nil {
		output.PrintError("Ошибка при сохранении хранилища vault")
		output.PrintError(err.Error())
	}
}

//FindAccounts поиск аккаунта по переданным параметрам
func (vault *VaultWithdb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var foundedAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
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
				output.PrintError("Ошибка при удалении аккаунта по переданному URL: " + URL)
				output.PrintError(err.Error())
			}
			return true
		}
	}
	return false
}

func (vault *VaultWithdb) save() error {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	//vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError(err.Error())
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
