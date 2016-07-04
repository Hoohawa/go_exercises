package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	hash1 := sha256.Sum256([]byte{1, 2, 3, 4, 5})
	hash2 := sha256.Sum256([]byte{5, 4, 3, 2, 1})

	fmt.Printf("H1: ")
	PrintHashBinary(hash1)
	fmt.Printf("H2: ")
	PrintHashBinary(hash2)

	fmt.Println("Differing bits: ", countDifferentBits(hash1, hash2))
}

func countDifferentBits(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += int(pc[hash1[i]^hash2[i]])
	}
	return count
}

func PrintHashBinary(hash [32]byte) {
	for _, b := range hash {
		fmt.Printf("%08b", b)
	}
	fmt.Println()
}
