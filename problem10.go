package main

import "fmt"

func main() {
    array1 := []int{1, 2, 3, 4, 11}
    array2 := []int{3, 9, 5, 6, 7}

    finalArray := []int{}

    finalArray = appender(array1, finalArray)
    finalArray = appender(array2, finalArray)

    fmt.Println(sorter(finalArray))
}

func sorter(array []int) []int {
    for i := 0; i < len(array)-1; i++ {
        if array[i] > array[i+1] {
            array[i], array[i+1] = array[i+1], array[i]
            i = 0
        }
    }
    return array
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
