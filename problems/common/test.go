package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := firstUniqChar("aaaa")
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	m := make(map[byte]int, 0)
	a := make(map[byte]bool, len(s))
	for i := 0; i < len(s); i++ {
		_, ok1 := m[s[i]]
		_, ok2 := a[s[i]]
		if ok1 {
			delete(m, s[i])
			a[s[i]] = true
		} else if !ok1 && !ok2 {
			m[s[i]] = i
		} else if !ok1 && ok2 {
			continue
		}
	}
	if len(m) == 0 {
		return -1
	}
	res := len(s)
	for _, v := range m {
		if v < res {
			res = v
		}
	}
	return res
}
