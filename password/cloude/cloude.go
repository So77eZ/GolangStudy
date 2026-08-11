package cloude

import (
	"fmt"
	"os"
)

//CloudeDb - Создает новый экземпляр CloudeDb с указанным ссылкой на облачное хранилище
type CloudeDb struct {
	url string
}

// NewCloudeDb - Создает новый экземпляр CloudeDb с указанным URL
func NewCloudeDb(url string) *CloudeDb {
	return &CloudeDb{
		url: url,
	}
}

// Write - Записать что-то в файл
func (db *CloudeDb) Write(content []byte) error {
	file, err := os.Create(db.url)
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

// Read - Чтение файла и запись его куда-то
func (db *CloudeDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
