package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// stdin - общий буферизованный ридер для всего ввода программы.
// Создавать новый на каждый вызов нельзя: часть уже вычитанных в буфер
// данных терялась бы между вызовами.
var stdin = bufio.NewReader(os.Stdin)

// GetUserInput - получение пользовательского ввода из консоли по параметру запроса.
// Второе возвращаемое значение равно false, если ввод закончился (EOF)
// и запрашивать что-либо дальше бессмысленно.
func GetUserInput(prompt ...any) (string, bool) {
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v: ", value)
		} else {
			fmt.Println(value)
		}
	}
	line, err := stdin.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	// Последняя строка без перевода строки приходит вместе с io.EOF,
	// поэтому пустой результат отличаем от непустого
	if err != nil && line == "" {
		fmt.Println()
		return "", false
	}
	return line, true
}
