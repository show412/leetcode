package main

import (
	"fmt"
	"strings"
)

/*
test case:
[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3
10

[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
6
*/
func main() {
	res := toGoatLatin("The quick brown fox jumped over the lazy dog")
	// 998001
	fmt.Println(res)
}

func toGoatLatin(S string) string {
	arr := strings.Split(S, " ")
	vowel := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for i := 0; i < len(arr); i++ {
		if _, ok := vowel[arr[i][0]]; ok {
			arr[i] += "ma"
		} else {
			arr[i] = arr[i][1:] + string(arr[i][0]) + "ma"
		}
		arr[i] += strings.Repeat("a", i+1)
	}
	return strings.Join(arr, " ")
}
