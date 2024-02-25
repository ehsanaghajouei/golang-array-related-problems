package main

import "fmt"

func main() {
	array01 := []int{5, 10, 12, 14, 16}
	// array02 := []int{99, 58, 11, 10, 1}
	// array03 := []int{1, 2, 19, 10, 3}

	if peoblem09(array01) {
		fmt.Println("non-decreasing")
	}
}

func peoblem09(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}
