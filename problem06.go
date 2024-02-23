package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 5, 3, 2}
	arr2 := []int{1, 2, 3, 4}

	fmt.Println(arr1)
	fmt.Println(arr2)

	if problem06(arr1, arr2) {
		fmt.Println("Equal.")
	} else {
		fmt.Println("Not Equal.")
	}
}

func sortSlice(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			i = 0
		}
	}
	return slice
}

func problem06(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		fmt.Println("Length is not equal.")
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if sortSlice(slice1)[i] != sortSlice(slice2)[i] {
			fmt.Println(slice1[i], "is not equal to", slice2[i])
			return false
		}
	}
	return true
}
