// Works with unicode as well,try with:
// s1: 沒精打采
// s2: 打采沒精
package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	if isAnagram(s1, s2) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("NOT Anagram")
	}
}

func isAnagram(s1, s2 string) bool {
	count1 := getLetterCounts(s1)
	count2 := getLetterCounts(s2)
	for i := 0; i < 256; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}
	return true
}

func getLetterCounts(str string) [256]int {
	var cnt [256]int
	for i := range str {
		cnt[int(str[i])]++
	}
	return cnt
}
