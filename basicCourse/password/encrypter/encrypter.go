package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// Encrypter - структура для шифрования и дешифрования строк
type Encrypter struct {
	Key []byte
}

// NewEncrypter - создает новый экземпляр Encrypter с ключом из переменной окружения KEY.
// KEY должен быть hex-строкой длиной 32, 48 или 64 символа (AES-128/192/256).
func NewEncrypter() *Encrypter {
	res := os.Getenv("KEY")
	if res == "" {
		panic("Не передан параметр KEY в параметры окружения")
	}
	key, err := hex.DecodeString(res)
	if err != nil {
		panic("Параметр KEY не является корректной hex-строкой: " + err.Error())
	}
	return &Encrypter{Key: key}
}

// Encrypt - метод для шифрования строки
func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}

// Decrypt - метод для дешифрования строки.
// Возвращает ошибку, если данные повреждены, зашифрованы другим ключом
// или сохранены в незашифрованном виде.
func (enc *Encrypter) Decrypt(encryptedString []byte) ([]byte, error) {
	block, err := aes.NewCipher(enc.Key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	if len(encryptedString) < nonceSize {
		return nil, errors.New("данные короче размера nonce, файл повреждён или не зашифрован")
	}
	nonce, cipherText := encryptedString[:nonceSize], encryptedString[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
