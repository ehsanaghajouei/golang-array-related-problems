package main

import "fmt"

func main() {
	slice01 := randomArrayGenerator(5, 100)
	fmt.Println(slice01)
	// fmt.Println(len(slice01))

	// fmt.Println(problem01(slice01))
	// fmt.Println(problem02(slice01))
	// fmt.Println(problem03(slice01))
	// fmt.Println(problem04(slice01))
	fmt.Println(problem05(slice01, 2))
}
