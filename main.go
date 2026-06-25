package main

import (
	"fmt"
	"math"
)

const IMTpower  = 2

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	userWeight,userHeight := getUserInput()
	IMT := calculateIMT(userWeight,userHeight)
	outputResult(IMT)
}

func outputResult(imt float64){
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f",imt)
	fmt.Print(result)
}

func calculateIMT(userWeight float64, userHeight float64) float64{
	return userWeight / math.Pow(userHeight/100,IMTpower)
}

func getUserInput() (float64,float64){
	var userHeight float64 
	var userWeight float64

	fmt.Print("Введите свой рост в сантиметрах (Пример - 171): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг (Пример - 62): ")
	fmt.Scan(&userWeight)
	return userWeight,userHeight
}