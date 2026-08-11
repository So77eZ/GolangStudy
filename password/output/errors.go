package output

import (
	"github.com/fatih/color"
)

// PrintError - ф-ция для вывода ошибки в консоль с использованием цветного типизирования вывода
func PrintError(value any) {
	typeDefine(value)
}

// typeDefine - функция для определения типа значения и вывода соответствующего сообщения об ошибке
func typeDefine(value any) {
	switch t := value.(type) {
	case string:
		color.Black(t)
	case error:
		color.White(t.Error())
	case int:
		color.Cyan("Код ошибки: %d", t)
	default:
		color.Red("Неизвестный тип ошибки: %T", t)
	}
}
