package main

import "fmt"

func main() {
    array1 := []int{1, 2, 3, 4, 5}
    array2 := []int{3, 4, 5, 6, 7}

    finalArray := []int{}

    finalArray = appender(array1, finalArray)
    finalArray = appender(array2, finalArray)

    fmt.Println(finalArray)
}

func appender(source []int, destination []int) []int {
    for i := 0; i < len(source); i++ {
        if !checkExist(destination, source[i]) {
            destination = append(destination, source[i])
        }
    }
    return destination
}

func checkExist(array []int, x int) bool {
    for i := 0; i < len(array); i++ {
        if x == array[i] {
            return true
        }
    }
    return false
}
