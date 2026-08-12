package output

import (
	"github.com/fatih/color"
)

// PrintError - ф-ция для вывода ошибки в консоль с использованием цветного типизирования вывода
func PrintError(value any) {
	typeDefineV1(value)
}

// typeDefineV1 - функция для определения типа значения и вывода соответствующего сообщения об ошибке
func typeDefineV1(value any) {
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

// typeDefineV2 - функция для определения типа значения и вывода соответствующего сообщения об ошибке
func typeDefineV2(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		color.Red(stringValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")
}
