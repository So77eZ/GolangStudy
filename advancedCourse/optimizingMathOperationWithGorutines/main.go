package main

//1. summ of elems slice
import (
	"fmt"
	"time"
)

func main() {
	//numberArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numberArray := make([]int, 50_000_001)
	for i := range numberArray {
		numberArray[i] = i
	}
	//0       1        2        3
	numGoroutines := 4

	sumChannel := make(chan int, numGoroutines)

	partSize := len(numberArray) / numGoroutines
	startTimeRutines := time.Now()
	for i := 0; i < numGoroutines; i++ {

		startIndex := i * partSize        //0*(12/4)=0i//1*3=3i//2*3=6i//3*3=9i
		endIndex := startIndex + partSize //0+(12/4)=3i//3+3=3i//6+3=6i//9+3=12i
		if i == numGoroutines-1 {
			endIndex = len(numberArray)
		}
		go sumParts(numberArray[startIndex:endIndex], sumChannel)
	}
	resultSumRutines := 0
	for i := 0; i < numGoroutines; i++ {
		resultSumRutines += <-sumChannel
	}
	executionTimeRutines := time.Since(startTimeRutines)

	startTimeSequential := time.Now()
	resultSumSequential := 0
	for _, num := range numberArray {
		resultSumSequential += num
	}
	executionTimeSequential := time.Since(startTimeSequential)

	fmt.Println("Execution time (Routines): ", executionTimeRutines)
	fmt.Println("sum (Routines): ", resultSumRutines)
	fmt.Println("Execution time (Sequential): ", executionTimeSequential)
	fmt.Println("sum (Sequential): ", resultSumSequential)
}

func sumParts(arr []int, sumChannel chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sumChannel <- sum
}
