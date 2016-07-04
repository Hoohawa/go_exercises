package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hashFuncFlag = flag.String("h", "256", "-h [256|384|512]")

func main() {
	hashFunc := "256"
	if len(os.Args) > 1 {
		flag.Parse()
		hashFunc = *hashFuncFlag
	}

	var s string
	fmt.Println("Input a string:")
	fmt.Scanf("%s", &s)
	switch {
	case hashFunc == "256":
		fmt.Printf("%v\n", sha256.Sum256([]byte(s)))
	case hashFunc == "384":
		fmt.Printf("%v\n", sha512.Sum384([]byte(s)))
	case hashFunc == "512":
		fmt.Printf("%v\n", sha512.Sum512([]byte(s)))
	case true:
		fmt.Println("Unrecognized hash function")
	}
}
