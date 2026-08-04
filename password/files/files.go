package files

import (
	"fmt"
	"os"
)

//WriteFile - Записать что-то в файл
func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Не удалось создать файл! ", err)
		return err
	}
	defer file.Close() // При выполнении всего стэка - произойдет вызов через defer ф-ции file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Не удалось записать файл! ", err)
		return err
	}
	fmt.Println("Файл успешно создан и записан!")
	return nil
}

//ReadFile - Чтение файла и запись его куда-то
func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
