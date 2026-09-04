// !Вывести сумму баланса в консоль

package main

import (
	"fmt"
	"strconv"
)

func main() {

	transactions := []float64{}
	fmt.Println("Программа учета транзакций")
	fmt.Println("Вводите транзакции (или 'q' для выхода): ")
	for {
		input := getUserInput()
		if input == "q" || input == "Q" || input == "" {
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
		fmt.Printf("Сумма баланса: %.2f\n", getUserWalletBalance(transactions))
	}
}

func getUserInput() string {
	var userInputNumber string
	fmt.Scan(&userInputNumber)
	return userInputNumber
}

func getUserWalletBalance(transactions []float64) float64 {
	balance := 0.0
	for _, transaction := range transactions {
		balance += transaction
	}
	return balance
}
