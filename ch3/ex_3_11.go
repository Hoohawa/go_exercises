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
	sign, str := splitSign(s)
	str, dec := splitDecimalPoint(str)
	n := len(str)
	if n <= 3 {
		return sign + str + dec
	}
	mod := n % 3
	result := str[:mod]
	for i := mod; i < n; i += 3 {
		if i > 0 {
			result += ","
		}
		result += str[i : i+3]
	}
	return sign + result + dec
}

func splitDecimalPoint(s string) (string, string) {
	pointIdx := -1
	for i := range s {
		if s[i] == '.' {
			pointIdx = i
			break
		}
	}
	if pointIdx > -1 {
		return s[:pointIdx], s[pointIdx:]
	}
	return s, ""
}

func splitSign(s string) (string, string) {
	if len(s) > 0 && s[0] == '-' || s[0] == '+' {
		return string(s[0]), s[1:]
	}
	return "", s
}
