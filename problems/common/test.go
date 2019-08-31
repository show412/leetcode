package main

import (
	"fmt"
	// "math"
)

/*
test case:
[0,0]
-1
true

[0,1,0]
0
false

[23, 2, 6, 4, 7],  k=0
false

[5,0,0]
0
true

{23, 2, 6, 4, 7}, 6
true

[23, 2, 4, 6, 7], 6
true
*/
func main() {
	res := isOneEditDistance("1203", "1213")
	fmt.Println(res)
}

func isOneEditDistance(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls > lt {
		return isOneEditDistance(t, s)
	}
	if lt-ls > 1 {
		return false
	}
	for i := 0; i < ls; i++ {
		if s[i] != t[i] {
			if ls == lt {
				return s[(i+1):] == t[(i+1):]
			} else {
				return s[i:] == t[(i+1):]
			}
		}
	}
	return ls+1 == lt
}
