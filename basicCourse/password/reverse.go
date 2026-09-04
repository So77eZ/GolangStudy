package main

import "fmt"

func in() {
	massive := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Исходный массив:", massive)
	reverse(&massive)
	fmt.Println("Массив после реверса:", massive)
}

func reverse(revMas *[6]int) {
	for index, value := range *revMas {
		(*revMas)[len(*revMas)-1-index] = value
	}
}
