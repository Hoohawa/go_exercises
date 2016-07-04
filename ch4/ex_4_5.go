package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b", "b", "a", "c", "c", "a", "b"}
	slice = removeAdjacentDuplicates(slice)
	fmt.Println(slice)
}

func removeAdjacentDuplicates(slice []string) []string {
	cnt := 0
	for pos := 0; pos < len(slice); pos++ {
		slice[cnt] = slice[pos]
		cnt++
		for i := pos + 1; i < len(slice); i++ {
			if slice[pos] != slice[i] {
				pos = i - 1
				break
			}
		}
	}
	return slice[:cnt]
}
