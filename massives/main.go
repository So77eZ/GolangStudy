package main

import (
	"fmt"
	"strconv"
)

func main() {
	var transactions []float64
	fmt.Print("Программа учета транзакций")
	for {
		fmt.Print("Введите транзакцию (или 'q' для выхода): ")
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
		fmt.Println("Баланс кошелька:", getUserWalletBlance(transactions))
	}
}

func getUserInput() string {
	var userInputNumber string
	fmt.Scan(&userInputNumber)
	return userInputNumber
}

func getUserWalletBlance(transactions []float64) float64 {
	balance := 0.0
	for _, transaction := range transactions {
		balance += transaction
	}
	return balance
}
