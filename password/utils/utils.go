package utils

import (
	"fmt"
)

//GetUserInput - получение пользовательского ввода из консоли
func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}
