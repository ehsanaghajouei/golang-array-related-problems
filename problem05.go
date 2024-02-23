package main

func problem05(slice []int, left int) []int {
	leftedSlice := []int{}
	leftedSlice = append(leftedSlice, slice...)
	for i := 0; i < len(slice); i++ {
		if left >= len(slice) || len(slice) < len(leftedSlice) {
			left = -1
			left++
		}
		leftedSlice[i] = slice[left]
		left++
	}
	return leftedSlice
}
