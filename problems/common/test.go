package main

import (
	"fmt"
	// "math"
)

func main() {
	/*
		"heeellooo"
		["axxxrrzzz"]
	*/
	/*
		"zzzzzyyyyy"
		["zzyy","zy","zyy"]
	*/
	/*
		"heeellooo"
		"hello", "hi", "helo"
	*/
	/*
			"dddiiiinnssssssoooo"
		  ["dinnssoo","ddinso","ddiinnso","ddiinnssoo","ddiinso","dinsoo","ddiinsso","dinssoo","dinso"]
	*/
	res := expressiveWords("heeellooo", []string{"axxxrrzzz"})
	fmt.Println(res)
}

func expressiveWords(S string, words []string) int {
	if len(words) == 0 {
		return 0
	}
	res := 0
	SByteArray, SByteNum := generatePair(S)
	for i := 0; i < len(words); i++ {
		word := words[i]
		wordByteArray, wordByteNum := generatePair(word)
		// fmt.Println("****")
		// fmt.Println(SByteArray)
		// fmt.Println(wordByteArray)
		// fmt.Println(SByteNum)
		// fmt.Println(wordByteNum)
		if sliceEqual(SByteArray, wordByteArray) == true {
			for j := 0; j < len(wordByteNum); j++ {
				if SByteNum[j] < 3 && SByteNum[j] != wordByteNum[j] || SByteNum[j] < wordByteNum[j] {
					break
				}
				// fmt.Println(j)
				if j == len(wordByteNum)-1 {
					res++
				}

			}
		}
	}
	return res
}

func generatePair(s string) ([]byte, []int) {
	var byteArray []byte
	var byteNum []int
	byteArray = append(byteArray, s[0])
	num := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			num++
		} else {
			byteArray = append(byteArray, s[i])
			byteNum = append(byteNum, num)
			num = 1
		}
		if i == len(s)-1 {
			byteNum = append(byteNum, num)
		}
	}
	return byteArray, byteNum
}

func sliceEqual(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	if (s1 == nil) != (s2 == nil) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
