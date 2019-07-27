package main

import (
	"fmt"
	"sort"
	"strings"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := expand("{a,b}c{d,e}f")
	fmt.Println(res)
}

var result []string

func expand(S string) []string {
	res := make([]string, 0)
	expandStr(S, &res)
	return res
}

func expandStr(S string, res *[]string) {
	if strings.Index(S, "{") == -1 {
		*res = append(*res, S)
		return
	}
	openIndex := strings.Index(S, "{")
	closeIndex := strings.Index(S, "}")
	options := strings.Split(S[openIndex+1:closeIndex], ",")
	sort.Strings(options)
	for i := 0; i < len(options); i++ {
		expandStr(S[:openIndex]+options[i]+S[closeIndex+1:], res)
	}
}
