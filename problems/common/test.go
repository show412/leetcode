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
	res := checkWord("abcdbbdbadadacbcccabbcadddbca", []string{"abcdbbdbadadacbcccabbcadddbca", "abcdbbdbadadacbcccabbcadddca", "bdd", "abbb", "dbbcd", "abcdbbdbadadacbcccabbca", "cbcacbd", "accacaca", "cbbcabcdd", "abdacaadbb", "bccaccbbdbb", "acaadbccdaca", "ddcdbddccccab", "aaaddcdbcaabba", "ccabdddadcbbdaa", "dddaccdbbbdcdbbd", "aaddacccaacbccdac", "bacaddaacacadbdbab", "aaaacbbcabbdbbbdadd", "bbbabadccaacbdaaaaab", "cbccaaddddcdcddccbaca", "abcdbbdbadadacbcccab", "bcadadaaddcbadcbdcabdcd", "acdcbcbdccbaacdcccdacbbb", "cadbaabbaccbbccccccbcaacc", "bddddbddbdcdbaaddccbacaccc", "abcacbbaaddcdabcdbdaddbdbdb", "dbdabdcaaccdccdcbccadddacdda", "abbdcaacdcacadbbbcbccaacdcdbb", "aadddadcbda", "abdbdbbadb", "addccacaaabccdc", "abcdbbdbadadacbccb", "dddcbdacccbaaacacccdbdd", "cbabcacbccbddabaadbdbdcbaacd", "cdcbcdddbcabbccbcbbbdabbadd", "a", "cabbcacaadbaddcbaacc", "cddbdba", "abcdbbdbadadcb", "abcdbbdbadadcbb", "ad", "aabadadcadcdcbdbbacdaabacaad", "abcd", "abcdbd", "abcdbbd", "abcdbbdd", "adddbcab", "abcdbbdbad", "abcdbbdbadd", "abcdbbdbadad", "abcdbbdbadadc", "ddabdbbacbcadbcacdacddaababc", "cdcbddcabcbdacbbaddba", "abcdbbdbadadacbcccabbcadca", "abcdbbdbadadacbcccabbcaddca", "abcdbbdbadadacbcccabbcaca", "abcdbbdbadadacbcccabbcca", "aababdddbdc", "dbdbcb", "aaddaabcaadcd", "daaaddbb", "bccdbdbadabcdbccbdcbacabc", "bcadadcbb", "babcdadadddabdaaadd", "cccbadbbabaadaadcadccccdc", "b", "cdcbcabbbcaa", "cbaccaddddbbb", "aabbcbbbaaccabdbabbcbddbdacb", "abcdbbdbadadacbcccabba", "abcdbbdbadadacbcccabb", "badabd", "abccbccbbdbbcdbb", "caadbcacbbcdabacca", "ddacdcdbbcbb", "cacbcabcdccda", "abda", "dcaaddaddadaddcdbbbbb", "caddb", "bbcdbddbdcbcccabb", "badb", "cabcdccbadbbabbbdbbcdbad", "ddcaddcdbacdcbadbbbbdbbcdc", "b", "cccdcaabdcabcbbcaabababddda", "dcadabcadadcbbcacdaccbb", "abcdbbdbadadacbcccb", "cbacbbacacbabdadc", "acbdacbaacaac", "aacbccbbbbcacddaa", "bcdbaab", "caadcaadbaadababddcbbabaacbdd", "badbbacbabcdabbcaddddc", "abcdbbdbadadcbcb", "abcdbbdbadadcbccb", "abd", "abcdd", "addcabbbdabaa", "abcdbbdbd"})
	fmt.Println(res)
}

func checkWord(s string, str []string) bool {
	// Write your code here
	slen := len(s)
	if slen > len(str) {
		return false
	}
	if slen == 0 {
		return false
	}
	// result := make([]string, 0)
	m := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		m[str[i]] = true
	}
	strMap := make(map[string]bool, len(s))
	// sort.Strings(str)
	return dfs(s, m, strMap)
}
func dfs(s string, m map[string]bool, strMap map[string]bool) bool {
	// result = append(result, s)
	// fmt.Println(result)
	// if len(s) == 1 || len(s) == len(m) {

	// for k, v := range m {
	if exist, ok := m[s]; !ok || exist == false {
		return false
	}
	// }
	// return true
	// }
	if len(s) == 1 {
		return true
	}
	if v, ok := strMap[s]; ok && v == true {
		return false
	}
	for i := 0; i < len(s); i++ {
		sub := s
		if i == 0 {
			sub = s[1:]
		} else {
			sub = s[0:i] + s[(i+1):len(s)]
		}
		if dfs(sub, m, strMap) == true {
			return true
		}
		// result = result[0 : len(result)-1]
	}
	strMap[s] = true
	return false
}
