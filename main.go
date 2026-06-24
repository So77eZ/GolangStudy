package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTpowder  = 2
	var userHeight float64 
	var userWeight float64

	//Ввод пользователя
	fmt.Print("Калькулятор индекса массы тела\n")
	fmt.Print("Введите свой рост в метрах (Пример - 1.71): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг (Пример - 62): ")
	fmt.Scan(&userWeight)

	IMT := userWeight / math.Pow(userHeight,IMTpowder)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(math.Floor(IMT))
}