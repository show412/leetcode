package main

import (
	"fmt"
	// "math"
)

func main() {
	res := findLongestWord("abpcplea", []string{"a", "b", "c"})
	fmt.Println(res)
}

func findLongestWord(s string, d []string) string {
	if len(d) == 0 || len(s) == 0 {
		return ""
	}
	candidateList := make([]string, 0)
	// find which word is in s
	for i := 0; i < len(d); i++ {
		word := d[i]
		pointS := 0
		pointW := 0
		for pointW < len(word) && pointS < len(s) {
			if pointW == len(word)-1 && s[pointS] == word[pointW] {
				candidateList = append(candidateList, word)
				break
			}
			if s[pointS] == word[pointW] {
				pointS++
				pointW++
				continue
			} else {
				if (len(s) - pointS) < (len(word) - pointW) {
					break
				}
				pointS++
			}
		}
	}
	// find out the result
	if len(candidateList) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(candidateList); i++ {
		word := candidateList[i]
		if len(word) > len(res) {
			res = word
			continue
		}
		if len(word) == len(res) {
			res = checkLex(word, res)
		}
	}
	return res
}

func checkLex(word1 string, word2 string) string {
	s := 0
	for s < len(word1) {
		if word1[s] < word2[s] {
			return word1
		} else if word1[s] > word2[s] {
			return word2
		}
		s++
	}
	return word1
}
