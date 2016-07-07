package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordFreq := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordFreq[scanner.Text()]++
	}

	fmt.Println("Word frequencies:")
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
