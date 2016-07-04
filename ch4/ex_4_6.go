package main

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	str1 := []byte("AA\u00A0\tb沒精  v\u0085c     ")
	str2 := []byte("A\u00A0\tb沒精   c")

	fmt.Println(string(str1))
	fmt.Println(string(str2))

	fmt.Println(string(compressSpaces(str1)))
	fmt.Println(string(compressSpaces(str2)))
}

func compressSpaces(str []byte) []byte {
	cnt := 0
	for i := 0; i < len(str); i++ {
		str[cnt] = str[i]

		r, _ := utf8.DecodeRuneInString(string(str[i:]))
		if unicode.IsSpace(r) {
			nextNonSpace, err := findFirstNonSpace(str[i:])
			if err != nil { // Ends with spaces
				str[cnt] = ' '
				return str[:cnt+1]
			} else {
				i += nextNonSpace - 1
				str[cnt] = ' ' // replace with ASCII space
			}
		}
		cnt++
	}
	return str[:cnt]
}

func findFirstNonSpace(str []byte) (int, error) {
	for i, r := range string(str) {
		if !unicode.IsSpace(r) {
			return i, nil
		}
	}
	return len(str), errors.New("No non-space found")
}
