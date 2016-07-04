package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for idx, arg := range os.Args {
		s += sep + strconv.Itoa(idx) + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
