package main

func problem02(slice []int) []int {
	revese := []int{}

	for i := len(slice) - 1; i >= 0; i-- {
		revese = append(revese, slice[i])
	}

	return revese
}
