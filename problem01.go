package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)
	// fmt.Println(len(slice01))

	fmt.Println(problem01(slice01))
}

func problem01(slice []int) int {
	// for i, v := range slice01 {
	// 	if i != len(slice01)-1 && slice01[i] < slice01[i+1] {
	// 	}

	sum := 0

	for _, v := range slice {
		sum += v
	}

	return sum
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
