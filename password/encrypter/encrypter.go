package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

// Encrypter - структура для шифрования и дешифрования строк
type Encrypter struct {
	Key string
}
// NewEncrypter - создает новый экземпляр Encrypter с ключом из переменной окружения SECRET_KEY
func NewEncrypter() *Encrypter {
	res := os.Getenv("SECRET_KEY")
	if res == "" {
		panic("Не передан параметр SECRET_KEY в параметры окружения")
	}
	return &Encrypter{Key: res}
}

// Encrypt - метод для шифрования строки
func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	block,err:=aes.newCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err:= io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}
// Decrypt - метод для дешифрования строки
func (enc *Encrypter) Decrypt(encryptedString []byte) []byte {
	
	return []
}
