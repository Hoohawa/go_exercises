package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	categories := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		// Count various categories of runes
		categories["control"] = Btoi(unicode.IsControl(r))
		categories["digit"] += Btoi(unicode.IsDigit(r))
		categories["graphic"] += Btoi(unicode.IsGraphic(r))
		categories["letter"] += Btoi(unicode.IsLetter(r))
		categories["lower"] += Btoi(unicode.IsLower(r))
		categories["mark"] += Btoi(unicode.IsMark(r))
		categories["number"] += Btoi(unicode.IsNumber(r))
		categories["print"] += Btoi(unicode.IsPrint(r))
		categories["punct"] += Btoi(unicode.IsPunct(r))
		categories["space"] += Btoi(unicode.IsSpace(r))
		categories["symbol"] += Btoi(unicode.IsSymbol(r))
		categories["title"] += Btoi(unicode.IsTitle(r))

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Println("Categories: ")
	for k, v := range categories {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
