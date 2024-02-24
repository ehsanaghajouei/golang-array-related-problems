package main

import "fmt"

func main() {
	array01 := [4]int{4, 11, 9, 23}
	array02 := [4]int{1, 4, 7, 23}

	fmt.Println(array01)
	fmt.Println(array02)

	fmt.Println(problem08(array01[:], array02[:]))
}

func problem08(array1 []int, array2 []int) []int {
	intersection := []int{}
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				intersection = append(intersection, array1[i])
			}
		}
	}
	return intersection
}
