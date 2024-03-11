package main

import (
    "fmt"
)

func main() {
    array := []int{5,2,7,7,5,5,2,7,8,2,2,2,5,7,8}

    fmt.Println(frequency(sorter(array)))
}

func sorter(array []int) []int {
    for i := 0; i < len(array)-1; i++ {
        if array[i] > array[i+1] {
            array[i],array[i+1] = array[i+1],array[i]
            i = 0
        }
    }
    return array
}

func frequency(array []int) map[int]int {
    count := map[int]int{}
    itemCount := 1
    for i := 0; i < len(array)-1; i++ {
        if array[i] == array[i+1] {
            itemCount++
            count[array[i]] = itemCount
        } else {
            itemCount = 1
        }
    }
    return count
}
