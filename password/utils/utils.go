package utils

import (
	"fmt"
)

//GetUserInput - получение пользовательского ввода из консоли по параметру запроса
func GetUserInput(prompt string) string {
	fmt.Println(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}
