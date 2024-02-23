package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)
	// fmt.Println(len(slice01))

	fmt.Println(problem04(slice01))
}

func problem04(slice []int) []int {
	noneDuplicates := []int{}

	for _, v := range slice {
		if !existCheck(noneDuplicates, v) {
			noneDuplicates = append(noneDuplicates, v)
		}
	}

	return noneDuplicates
}

func existCheck(slice []int, x int) bool {
	for i := 0; i < len(slice); i++ {
		if x == slice[i] {
			return true
		}
	}
	return false
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
