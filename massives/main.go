package main

import (
	"fmt"
	"strconv"
)

func main() {
	var transactions []float64
	for {
		fmt.Print("Введите транзакцию (или 'q' для выхода): ")
		input := getUserInput()
		if input == "q" {
			break
		} else {
			input, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Ошибка: введите числовое значение транзакции")
				continue
			}
			transactions = append(transactions, input)
		}
	}
	if len(transactions) > 0 {
		fmt.Println("Массив введенных транзакций:")
		fmt.Println(transactions)
	}
}

func getUserInput() string {
	var userInputNumber string
	fmt.Scan(&userInputNumber)
	return userInputNumber
}
