package main

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
