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
	res := isIsomorphic("paper", "title")
	fmt.Println(res)
}

func isIsomorphic(s string, t string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	ms := make(map[byte]byte, 0)
	mt := make(map[byte]byte, 0)
	for i := 0; i < l; i++ {
		vs, oks := ms[s[i]]
		vt, okt := mt[t[i]]
		if !oks && !okt {
			ms[s[i]] = t[i]
			mt[t[i]] = s[i]
		} else if oks && okt {
			if vs != t[i] || vt != s[i] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
