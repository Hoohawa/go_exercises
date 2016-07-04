package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Printf("Input number: ")
	fmt.Scanf("%s", &n)
	fmt.Println(nonRecursiveComma(n))
}

func nonRecursiveComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	mod := n % 3
	result := s[:mod]
	for i := mod; i < n; i += 3 {
		if i > 0 {
			result += ","
		}
		result += s[i : i+3]
	}
	return result
}
