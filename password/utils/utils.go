package utils

import (
	"fmt"
)

//GetUserInput - получение пользовательского ввода из консоли по параметру запроса
func GetUserInput(prompt ...any) string {
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	var result string
	fmt.Scanln(&result)
	return result
}
