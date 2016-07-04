// Run with:
// $ go run ex_1_4.go $(ls ex_1_4_files/ | awk '{s="ex_1_4_files/"$1;print s}')
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fileNames := strings.Join(lineFiles[line], " ")
			fmt.Printf("%d\t%s Files: %s\n", n, line, fileNames)
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		lineFiles[input.Text()] = append(lineFiles[input.Text()], f.Name())
	}
}
