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
	res := getAns([][]int{[]int{1, 1, 1, 1}, []int{1, 1, 0, 1}, []int{1, 0, 1, 0}})
	fmt.Println(res)
}

func getAns(mp [][]int) int {
	// Write your code here.
	fa := make([]int, 5005)
	line := make([]int, 3000)
	column := make([]int, 3000)
	cnt := 0
	m := len(mp)
	n := len(mp[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 1 {
				fa[i*n+j+1] = i*n + j + 1
				line[i] = i*n + j + 1   //该点为1，给行存入编号
				column[j] = i*n + j + 1 //该点为1，给列存入编号
				cnt++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 1 {
				unity(i*n+j+1, line[i], fa)
				unity(i*n+j+1, column[j], fa)
			}
		}
	}
	// fmt.Println(fa[0:12])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 1 && fa[i*n+j+1] == i*n+j+1 {
				cnt--
			}
		}
	}
	return cnt
}

func unity(x int, y int, fa []int) []int {
	fmt.Println("*******")
	fmt.Println(x, y)
	x = find(x, fa)
	y = find(y, fa)
	fmt.Println(fa[0:13])
	fa[x] = y
	return fa
}

func find(x int, fa []int) int {
	if fa[x] == x {
		return x
	} else {
		fa[x] = find(fa[x], fa)
	}
	return fa[x]
}
