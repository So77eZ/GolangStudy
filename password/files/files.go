package files

import (
	"fmt"
	"os"
)

//JSONdb - структура для работы с JSON файлами
type JSONdb struct {
	filename string
}

//NewJSONdb - Создает новый экземпляр JsonDB с указанным именем файла
func NewJSONdb(name string) *JSONdb {
	return &JSONdb{
		filename: name,
	}
}

//Write - Записать что-то в файл
func (db *JSONdb) Write(content []byte) error {
	file, err := os.Create(db.filename)
	if err != nil {
		//fmt.Println("Не удалось создать файл! ", err)
		return err
	}
	defer file.Close() // При выполнении всего стэка - произойдет вызов через defer ф-ции file.Close()

	_, err = file.Write(content)
	if err != nil {
		//fmt.Println("Не удалось записать файл! ", err)
		return err
	}
	//fmt.Println("Файл успешно создан и записан!")
	return nil
}

//Read - Чтение файла и запись его куда-то
func (db *JSONdb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
