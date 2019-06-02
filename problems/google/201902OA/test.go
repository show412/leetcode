package main

import "fmt"

func main() {
	target := "hello"
	a := []string{"heylol"}
	b := getAns(target, a)
	fmt.Println(b)
}

func getAns(target string, s []string) []string {
	// Write your code here
	if len(target) == 0 || len(s) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	tLen := len(target)
	for i := 0; i < len(s); i++ {
		str := s[i]
		k := 0
		for j := 0; j < len(str); j++ {
			fmt.Println(string(target[k]))
			fmt.Println(string(str[j]))
			if string(str[j]) == string(target[k]) {
				if k < tLen-1 {
					k++
				}
			}
			if k == tLen-1 && j == len(str)-1 {
				res = append(res, str)
			}

		}
	}
	return res
}
