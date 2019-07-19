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
	res := canTransform("XXXXXLXXXX", "LXXXXXXXXX")
	fmt.Println(res)
}

func canTransform(start string, end string) bool {
	if len(start) == 1 {
		if start == end {
			return true
		} else {
			return false
		}
	}
	m := map[string]string{"XL": "LX", "RX": "XR"}
	for i := 0; i < len(start)-1; i++ {
		if start[i] == end[i] {
			continue
		} else {
			s := string(start[i]) + string(start[i+1])
			e := string(end[i]) + string(end[i+1])
			fmt.Println("*****")
			fmt.Println(i)
			fmt.Println(s)
			fmt.Println(e)
			if v, _ := m[s]; v == e {
				i++
			} else {
				return false
			}
		}
	}
	if start[len(start)-1] != end[len(end)-1] {
		return false
	}
	return true
}
