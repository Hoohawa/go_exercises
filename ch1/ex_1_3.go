package main

import (
	"fmt"
	"os"
	"strings"
)

func Approach1(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func Approach2(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(Approach1(os.Args[:]))
	fmt.Println(Approach2(os.Args[:]))
}
