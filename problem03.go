package main

func problem03(slice []int) int {

	max := 0

	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}
