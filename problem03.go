package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)
	// fmt.Println(len(slice01))

	fmt.Println(problem03(slice01))
}

func problem03(slice []int) int {

	max := 0

	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}

func randomArrayGenerator(limit int, rangeNumber int) []int {
	slice := []int{}
	for i := 0; i < limit; i++ {
		slice = append(slice, randomNumber(rangeNumber))
	}
	return slice
}

func randomNumber(rangeNumber int) int {
	return rand.Intn(rangeNumber)
}
