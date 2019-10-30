package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
test case:
[[1,1,4],[9,4,9],[9,1,9],[2,3,5],[4,1,5],[10,4,5]]
33
false

*/

func main() {
	res := generateAbbreviations("word")
	// 998001
	fmt.Println(res)
}
func generateAbbreviations(word string) []string {
	res := make([]string, 0)
	for i := 0; i < int(math.Pow(2.0, float64(len(word)))); i++ {
		out := ""
		cnt := 0
		for j := 0; j < len(word); j++ {
			if (i>>uint(j))&1 == 1 {
				cnt++
			} else {
				if cnt > 0 {
					out += strconv.Itoa(cnt)
					cnt = 0
				}
				out += string(word[j])
			}
		}
		if cnt > 0 {
			out += strconv.Itoa(cnt)
		}
		res = append(res, out)
	}
	return res
}
