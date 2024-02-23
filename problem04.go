package main

func problem04(slice []int) []int {
	noneDuplicates := []int{}
	for _, v := range slice {
		if !existCheck(noneDuplicates, v) {
			noneDuplicates = append(noneDuplicates, v)
		}
	}
	return noneDuplicates
}

func existCheck(slice []int, x int) bool {
	for i := 0; i < len(slice); i++ {
		if x == slice[i] {
			return true
		}
	}
	return false
}
