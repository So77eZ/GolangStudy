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
	outputResult(IMT,printUserDiagnosis(IMT))
}

func outputResult(imt float64, diagnosis string){
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f",imt)
	fmt.Println(diagnosis)
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

func printUserDiagnosis(IMT float64)string{
	switch {
	case IMT < 16:                  return "У вас сильный дефицит массы тела"
	case IMT >= 16   && IMT < 18.5: return "У вас дефицит массы тела"
	case IMT >= 18.5 && IMT < 25:	return "У вас нормальная масса тела"
	case IMT >= 25   && IMT < 30: 	return "У вас избыточная масса тела"
	case IMT >= 30   && IMT < 35: 	return "У вас 1-я степень ожирения"
	case IMT >= 35   && IMT < 40: 	return "У вас 2-я степень ожирения"
	default:                      	return "У вас 3-я степень ожирения"
	}
	// if IMT < 16{
	// 	return "У вас сильный дефицит массы тела"
	// }
	// if IMT >= 16 && IMT < 18.5{
	// 	return "У вас дефицит массы тела"
	// }
	// if IMT >= 18.5 && IMT < 25{
	// 	return "У вас нормальная масса тела"
	// }
	// if IMT >= 25 && IMT < 30{
	// 	return "У вас избыточная масса тела"
	// }
	// if IMT >= 30 && IMT < 35{
	// 	return "У вас 1-я степень ожирения"
	// }
	// if IMT >= 35 && IMT < 40{
	// 	return "У вас 2-я степень ожирения"
	// }
	// //if IMT > 40{
	// return "У вас 3-я степень ожирения"
	// //}
}