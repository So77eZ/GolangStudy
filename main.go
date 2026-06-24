package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTpowder float64 = 2
	userHeight := 1.8
	userWeight := 100.0
	IMT := userWeight / math.Pow(userHeight,IMTpowder)
	fmt.Print(IMT)
}