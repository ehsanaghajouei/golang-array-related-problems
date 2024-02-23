package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)

	println(problem07(slice01))
}

func problem07(slice []int) int {
	biggest := findBiggestElement(slice)
	secondBiggest := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > secondBiggest && slice[i] < biggest {
			secondBiggest = slice[i]
		}
	}
	return secondBiggest
}

func findBiggestElement(slice []int) int {
	biggest := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] > biggest {
			biggest = slice[i]
		}
	}
	return biggest
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
