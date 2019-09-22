package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
test case:
"a, a, a, a, b,b,b,c, c"
["a"]
"b"

paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
*/
func main() {
	res := mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"})
	// res := canFinish(4, [][]int{[]int{0, 1}})
	fmt.Println(res)
}

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]bool, 0)
	for _, v := range banned {
		bannedMap[v] = true
	}
	pureText := make([]byte, 0)
	for i := 0; i < len(paragraph); i++ {
		letter := paragraph[i]
		if unicode.IsLetter(rune(letter)) || unicode.IsSpace(rune(letter)) {
			pureText = append(pureText, letter)
		} else {
			pureText = append(pureText, ' ')
		}
	}
	wordArray := strings.Split(strings.ToLower(string(pureText)), " ")
	fmt.Println(wordArray)
	wordMap := make(map[string]int, 0)
	maxWord := ""
	maxNum := 0
	for i := 0; i < len(wordArray); i++ {
		word := strings.TrimSpace(wordArray[i])

		if bannedMap[word] == true || len(word) == 0 {
			continue
		}
		if _, ok := wordMap[word]; !ok {
			wordMap[word] = 0
		}
		wordMap[word] += 1
		if wordMap[word] > maxNum {
			maxWord = word
			maxNum = wordMap[word]
		}
	}
	return maxWord
}
