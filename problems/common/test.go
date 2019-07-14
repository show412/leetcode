package main

import (
	"fmt"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := reorganizeString("bbbbaaaaababaababab")
	fmt.Println(res)
}

func reorganizeString(S string) string {
	if len(S) == 1 {
		return S
	}
	res := make([]byte, len(S))
	m := make(map[byte]int)
	max := 0
	var maxLetter byte
	for i := 0; i < len(S); i++ {
		m[S[i]] += 1
		if m[S[i]] > max {
			max = m[S[i]]
			maxLetter = S[i]
		}
	}
	if max > (len(S)+1)/2 {
		return ""
	}
	index := 0
	for m[maxLetter] > 0 {
		res[index] = maxLetter
		m[maxLetter]--
		index += 2
	}
	for k, _ := range m {
		for m[k] > 0 {
			if index >= len(S) {
				index = 1
			}
			res[index] = k
			m[k]--
			index += 2
		}
	}
	return string(res)
}
