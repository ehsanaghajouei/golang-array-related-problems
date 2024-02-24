package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)
	// fmt.Println(len(slice01))

	fmt.Println(problem05(slice01, 2))
}

func problem05(slice []int, left int) []int {
	leftedSlice := []int{}
	leftedSlice = append(leftedSlice, slice...)

	for i := 0; i < len(slice); i++ {
		if left >= len(slice) {
			left = 0
		}
		leftedSlice[i] = slice[left]
		left++
	}
	return leftedSlice
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
