package main

import (
	"fmt"
	"math"
)

const IMTpower = 2

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	for {
		userWeight, userHeight := getUserInput()
		IMT := calculateIMT(userWeight, userHeight)
		// if err != nil {
		// 	//fmt.Println("Ошибка:", err)
		// 	fmt.Print("Ошибка: ")
		// 	panic(err)
		// 	//continue
		// }
		outputResult(IMT, printUserDiagnosis(IMT))
		if !repeatCalculation() {
			break
		}
	}
}

func repeatCalculation() bool {
	var userAnswer string
	fmt.Println("Продолжить рассчет? (y/n)")
	fmt.Scan(&userAnswer)
	return userAnswer == "y"
}

func outputResult(imt float64, diagnosis string) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println("\n" + diagnosis)
	fmt.Println(result)
}

func calculateIMT(userWeight float64, userHeight float64) float64 {
	return userWeight / math.Pow(userHeight/100, IMTpower)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	for {
		fmt.Print("\nВведите свой рост в сантиметрах (Пример - 171): ")
		if _, err := fmt.Scan(&userHeight); err != nil {
			fmt.Println("Ошибка: введите числовое значение роста")
			fmt.Scanln()
			continue
		}
		if userHeight <= 0 {
			fmt.Println("Ошибка: рост должен быть положительным числом")
			continue
		}
		break
	}

	for {
		fmt.Print("Введите свой вес в кг (Пример - 62): ")
		if _, err := fmt.Scan(&userWeight); err != nil {
			fmt.Println("Ошибка: введите числовое значение веса")
			fmt.Scanln()
			continue
		}
		if userWeight <= 0 {
			fmt.Println("Ошибка: вес должен быть положительным числом")
			continue
		}
		break
	}

	return userHeight, userWeight
}

func printUserDiagnosis(IMT float64) string {
	switch {
	case IMT < 16:
		return "У вас сильный дефицит массы тела"
	case IMT >= 16 && IMT < 18.5:
		return "У вас дефицит массы тела"
	case IMT >= 18.5 && IMT < 25:
		return "У вас нормальная масса тела"
	case IMT >= 25 && IMT < 30:
		return "У вас избыточная масса тела"
	case IMT >= 30 && IMT < 35:
		return "У вас 1-я степень ожирения"
	case IMT >= 35 && IMT < 40:
		return "У вас 2-я степень ожирения"
	default:
		return "У вас 3-я степень ожирения"
	}
}
