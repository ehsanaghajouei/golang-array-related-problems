package main

import "math/rand"

func randomNumber(rangeNumber int) int {
	return rand.Intn(rangeNumber)
}
