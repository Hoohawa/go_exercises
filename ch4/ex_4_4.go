package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(slice, 2))
}

func rotate(slice []int, n int) []int {
	for i := 0; i < n; i++ {
		slice = append(slice, slice[i])
	}
	return slice[n:]
}
