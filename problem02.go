package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)
	// fmt.Println(len(slice01))

	fmt.Println(problem02(slice01))
}

func problem02(slice []int) []int {
	revese := []int{}

	for i := len(slice) - 1; i >= 0; i-- {
		revese = append(revese, slice[i])
	}

	return revese
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
