package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)

	fmt.Println(problem07(slice01))
}

func problem07(slice []int) int {
	for {
		swapped := false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] < slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return slice[1]
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
