package main

import (
	"fmt"
	// "math"
	// "sort"
)

func main() {

	// res := minimumTotal([][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}})
	// a := []string{"abcdbbdbadadacbcccabbcadddbca", "abcdbbdbadadacbcccabbcadddca", "bdd", "abbb", "dbbcd", "abcdbbdbadadacbcccabbca", "cbcacbd", "accacaca", "cbbcabcdd", "abdacaadbb", "bccaccbbdbb", "acaadbccdaca", "ddcdbddccccab", "aaaddcdbcaabba", "ccabdddadcbbdaa", "dddaccdbbbdcdbbd", "aaddacccaacbccdac", "bacaddaacacadbdbab", "aaaacbbcabbdbbbdadd", "bbbabadccaacbdaaaaab", "cbccaaddddcdcddccbaca", "abcdbbdbadadacbcccab", "bcadadaaddcbadcbdcabdcd", "acdcbcbdccbaacdcccdacbbb", "cadbaabbaccbbccccccbcaacc", "bddddbddbdcdbaaddccbacaccc", "abcacbbaaddcdabcdbdaddbdbdb", "dbdabdcaaccdccdcbccadddacdda", "abbdcaacdcacadbbbcbccaacdcdbb", "aadddadcbda", "abdbdbbadb", "addccacaaabccdc", "abcdbbdbadadacbccb", "dddcbdacccbaaacacccdbdd", "cbabcacbccbddabaadbdbdcbaacd", "cdcbcdddbcabbccbcbbbdabbadd", "a", "cabbcacaadbaddcbaacc", "cddbdba", "abcdbbdbadadcb", "abcdbbdbadadcbb", "ad", "aabadadcadcdcbdbbacdaabacaad", "abcd", "abcdbd", "abcdbbd", "abcdbbdd", "adddbcab", "abcdbbdbad", "abcdbbdbadd", "abcdbbdbadad", "abcdbbdbadadc", "ddabdbbacbcadbcacdacddaababc", "cdcbddcabcbdacbbaddba", "abcdbbdbadadacbcccabbcadca", "abcdbbdbadadacbcccabbcaddca", "abcdbbdbadadacbcccabbcaca", "abcdbbdbadadacbcccabbcca", "aababdddbdc", "dbdbcb", "aaddaabcaadcd", "daaaddbb", "bccdbdbadabcdbccbdcbacabc", "bcadadcbb", "babcdadadddabdaaadd", "cccbadbbabaadaadcadccccdc", "b", "cdcbcabbbcaa", "cbaccaddddbbb", "aabbcbbbaaccabdbabbcbddbdacb", "abcdbbdbadadacbcccabba", "abcdbbdbadadacbcccabb", "badabd", "abccbccbbdbbcdbb", "caadbcacbbcdabacca", "ddacdcdbbcbb", "cacbcabcdccda", "abda", "dcaaddaddadaddcdbbbbb", "caddb", "bbcdbddbdcbcccabb", "badb", "cabcdccbadbbabbbdbbcdbad", "ddcaddcdbacdcbadbbbbdbbcdc", "b", "cccdcaabdcabcbbcaabababddda", "dcadabcadadcbbcacdaccbb", "abcdbbdbadadacbcccb", "cbacbbacacbabdadc", "acbdacbaacaac", "aacbccbbbbcacddaa", "bcdbaab", "caadcaadbaadababddcbbabaacbdd", "badbbacbabcdabbcaddddc", "abcdbbdbadadcbcb", "abcdbbdbadadcbccb", "abd", "abcdd", "addcabbbdabaa", "abcdbbdbd"}
	// fmt.Println(len(a))
	// res := checkWord("abcdbbdbadadacbcccabbcadddbca", []string{"abcdbbdbadadacbcccabbcadddbca", "abcdbbdbadadacbcccabbcadddca", "bdd", "abbb", "dbbcd", "abcdbbdbadadacbcccabbca", "cbcacbd", "accacaca", "cbbcabcdd", "abdacaadbb", "bccaccbbdbb", "acaadbccdaca", "ddcdbddccccab", "aaaddcdbcaabba", "ccabdddadcbbdaa", "dddaccdbbbdcdbbd", "aaddacccaacbccdac", "bacaddaacacadbdbab", "aaaacbbcabbdbbbdadd", "bbbabadccaacbdaaaaab", "cbccaaddddcdcddccbaca", "abcdbbdbadadacbcccab", "bcadadaaddcbadcbdcabdcd", "acdcbcbdccbaacdcccdacbbb", "cadbaabbaccbbccccccbcaacc", "bddddbddbdcdbaaddccbacaccc", "abcacbbaaddcdabcdbdaddbdbdb", "dbdabdcaaccdccdcbccadddacdda", "abbdcaacdcacadbbbcbccaacdcdbb", "aadddadcbda", "abdbdbbadb", "addccacaaabccdc", "abcdbbdbadadacbccb", "dddcbdacccbaaacacccdbdd", "cbabcacbccbddabaadbdbdcbaacd", "cdcbcdddbcabbccbcbbbdabbadd", "a", "cabbcacaadbaddcbaacc", "cddbdba", "abcdbbdbadadcb", "abcdbbdbadadcbb", "ad", "aabadadcadcdcbdbbacdaabacaad", "abcd", "abcdbd", "abcdbbd", "abcdbbdd", "adddbcab", "abcdbbdbad", "abcdbbdbadd", "abcdbbdbadad", "abcdbbdbadadc", "ddabdbbacbcadbcacdacddaababc", "cdcbddcabcbdacbbaddba", "abcdbbdbadadacbcccabbcadca", "abcdbbdbadadacbcccabbcaddca", "abcdbbdbadadacbcccabbcaca", "abcdbbdbadadacbcccabbcca", "aababdddbdc", "dbdbcb", "aaddaabcaadcd", "daaaddbb", "bccdbdbadabcdbccbdcbacabc", "bcadadcbb", "babcdadadddabdaaadd", "cccbadbbabaadaadcadccccdc", "b", "cdcbcabbbcaa", "cbaccaddddbbb", "aabbcbbbaaccabdbabbcbddbdacb", "abcdbbdbadadacbcccabba", "abcdbbdbadadacbcccabb", "badabd", "abccbccbbdbbcdbb", "caadbcacbbcdabacca", "ddacdcdbbcbb", "cacbcabcdccda", "abda", "dcaaddaddadaddcdbbbbb", "caddb", "bbcdbddbdcbcccabb", "badb", "cabcdccbadbbabbbdbbcdbad", "ddcaddcdbacdcbadbbbbdbbcdc", "b", "cccdcaabdcabcbbcaabababddda", "dcadabcadadcbbcacdaccbb", "abcdbbdbadadacbcccb", "cbacbbacacbabdadc", "acbdacbaacaac", "aacbccbbbbcacddaa", "bcdbaab", "caadcaadbaadababddcbbabaacbdd", "badbbacbabcdabbcaddddc", "abcdbbdbadadcbcb", "abcdbbdbadadcbccb", "abd", "abcdd", "addcabbbdabaa", "abcdbbdbd"})
	// res := getAns([][]int{[]int{1, 1, 1, 1}, []int{1, 1, 0, 1}, []int{1, 0, 1, 0}})
	res := isValid("{[]}")
	fmt.Println(res)
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	bracket := map[string]string{"(": ")", "{": "}", "[": "]"}
	stack := make([]string, 0)
	for i := len(s) - 1; i >= 0; i-- {
		sub := string(s[i])
		if v, ok := bracket[sub]; ok {
			close := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v != close {
				return false
			}
		} else {
			stack = append(stack, sub)
		}
	}
	return true
}
