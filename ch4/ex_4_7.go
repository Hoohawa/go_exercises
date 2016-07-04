package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	original := []byte("A沒B精打C采D")

	str := original[:]
	fmt.Println(string(str))

	reverseUnicode(str[:])
	fmt.Println(string(str))

	reverseUnicode(str[:])
	fmt.Println(string(str))

	if string(original) != string(str) {
		fmt.Println("ERROR: Strings do not match after reversing twice")
	}
}

func reverseUnicode(str []byte) {
	// Reverse all unicode characters
	for i := 0; i < len(str); {
		_, size := utf8.DecodeRune(str[i:])
		reverseBytes(str, i, i+size-1)
		i += size
	}
	// Reverse the entire slice
	reverseBytes(str, 0, len(str)-1)
}

func reverseBytes(str []byte, from, to int) {
	for i, j := from, to; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}
