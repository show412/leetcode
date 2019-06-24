package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := firstUniqChar("leetcode")
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	lenny := len(s)
	if lenny == 1 {
		return 0
	}

	cc := []byte(s)
	slow := 0
	fast := 1
	count := make([]int, 256)
	count[cc[slow]]++
	for fast < lenny {
		count[cc[fast]]++
		for slow < lenny && count[cc[slow]] > 1 {
			slow++
		}
		if slow >= lenny {
			return -1
		}
		if count[cc[slow]] == 0 {
			count[cc[slow]]++
			fast = slow
		}
		fast++
		fmt.Println(slow)
	}

	return slow
}
