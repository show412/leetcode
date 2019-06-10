package main

import (
	"fmt"
	"math"
)

func main() {

	// res := minimumTotal([][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}})
	// a := []string{"abcdbbdbadadacbcccabbcadddbca", "abcdbbdbadadacbcccabbcadddca", "bdd", "abbb", "dbbcd", "abcdbbdbadadacbcccabbca", "cbcacbd", "accacaca", "cbbcabcdd", "abdacaadbb", "bccaccbbdbb", "acaadbccdaca", "ddcdbddccccab", "aaaddcdbcaabba", "ccabdddadcbbdaa", "dddaccdbbbdcdbbd", "aaddacccaacbccdac", "bacaddaacacadbdbab", "aaaacbbcabbdbbbdadd", "bbbabadccaacbdaaaaab", "cbccaaddddcdcddccbaca", "abcdbbdbadadacbcccab", "bcadadaaddcbadcbdcabdcd", "acdcbcbdccbaacdcccdacbbb", "cadbaabbaccbbccccccbcaacc", "bddddbddbdcdbaaddccbacaccc", "abcacbbaaddcdabcdbdaddbdbdb", "dbdabdcaaccdccdcbccadddacdda", "abbdcaacdcacadbbbcbccaacdcdbb", "aadddadcbda", "abdbdbbadb", "addccacaaabccdc", "abcdbbdbadadacbccb", "dddcbdacccbaaacacccdbdd", "cbabcacbccbddabaadbdbdcbaacd", "cdcbcdddbcabbccbcbbbdabbadd", "a", "cabbcacaadbaddcbaacc", "cddbdba", "abcdbbdbadadcb", "abcdbbdbadadcbb", "ad", "aabadadcadcdcbdbbacdaabacaad", "abcd", "abcdbd", "abcdbbd", "abcdbbdd", "adddbcab", "abcdbbdbad", "abcdbbdbadd", "abcdbbdbadad", "abcdbbdbadadc", "ddabdbbacbcadbcacdacddaababc", "cdcbddcabcbdacbbaddba", "abcdbbdbadadacbcccabbcadca", "abcdbbdbadadacbcccabbcaddca", "abcdbbdbadadacbcccabbcaca", "abcdbbdbadadacbcccabbcca", "aababdddbdc", "dbdbcb", "aaddaabcaadcd", "daaaddbb", "bccdbdbadabcdbccbdcbacabc", "bcadadcbb", "babcdadadddabdaaadd", "cccbadbbabaadaadcadccccdc", "b", "cdcbcabbbcaa", "cbaccaddddbbb", "aabbcbbbaaccabdbabbcbddbdacb", "abcdbbdbadadacbcccabba", "abcdbbdbadadacbcccabb", "badabd", "abccbccbbdbbcdbb", "caadbcacbbcdabacca", "ddacdcdbbcbb", "cacbcabcdccda", "abda", "dcaaddaddadaddcdbbbbb", "caddb", "bbcdbddbdcbcccabb", "badb", "cabcdccbadbbabbbdbbcdbad", "ddcaddcdbacdcbadbbbbdbbcdc", "b", "cccdcaabdcabcbbcaabababddda", "dcadabcadadcbbcacdaccbb", "abcdbbdbadadacbcccb", "cbacbbacacbabdadc", "acbdacbaacaac", "aacbccbbbbcacddaa", "bcdbaab", "caadcaadbaadababddcbbabaacbdd", "badbbacbabcdabbcaddddc", "abcdbbdbadadcbcb", "abcdbbdbadadcbccb", "abd", "abcdd", "addcabbbdabaa", "abcdbbdbd"}
	// fmt.Println(len(a))
	// res := checkWord("abcdbbdbadadacbcccabbcadddbca", []string{"abcdbbdbadadacbcccabbcadddbca", "abcdbbdbadadacbcccabbcadddca", "bdd", "abbb", "dbbcd", "abcdbbdbadadacbcccabbca", "cbcacbd", "accacaca", "cbbcabcdd", "abdacaadbb", "bccaccbbdbb", "acaadbccdaca", "ddcdbddccccab", "aaaddcdbcaabba", "ccabdddadcbbdaa", "dddaccdbbbdcdbbd", "aaddacccaacbccdac", "bacaddaacacadbdbab", "aaaacbbcabbdbbbdadd", "bbbabadccaacbdaaaaab", "cbccaaddddcdcddccbaca", "abcdbbdbadadacbcccab", "bcadadaaddcbadcbdcabdcd", "acdcbcbdccbaacdcccdacbbb", "cadbaabbaccbbccccccbcaacc", "bddddbddbdcdbaaddccbacaccc", "abcacbbaaddcdabcdbdaddbdbdb", "dbdabdcaaccdccdcbccadddacdda", "abbdcaacdcacadbbbcbccaacdcdbb", "aadddadcbda", "abdbdbbadb", "addccacaaabccdc", "abcdbbdbadadacbccb", "dddcbdacccbaaacacccdbdd", "cbabcacbccbddabaadbdbdcbaacd", "cdcbcdddbcabbccbcbbbdabbadd", "a", "cabbcacaadbaddcbaacc", "cddbdba", "abcdbbdbadadcb", "abcdbbdbadadcbb", "ad", "aabadadcadcdcbdbbacdaabacaad", "abcd", "abcdbd", "abcdbbd", "abcdbbdd", "adddbcab", "abcdbbdbad", "abcdbbdbadd", "abcdbbdbadad", "abcdbbdbadadc", "ddabdbbacbcadbcacdacddaababc", "cdcbddcabcbdacbbaddba", "abcdbbdbadadacbcccabbcadca", "abcdbbdbadadacbcccabbcaddca", "abcdbbdbadadacbcccabbcaca", "abcdbbdbadadacbcccabbcca", "aababdddbdc", "dbdbcb", "aaddaabcaadcd", "daaaddbb", "bccdbdbadabcdbccbdcbacabc", "bcadadcbb", "babcdadadddabdaaadd", "cccbadbbabaadaadcadccccdc", "b", "cdcbcabbbcaa", "cbaccaddddbbb", "aabbcbbbaaccabdbabbcbddbdacb", "abcdbbdbadadacbcccabba", "abcdbbdbadadacbcccabb", "badabd", "abccbccbbdbbcdbb", "caadbcacbbcdabacca", "ddacdcdbbcbb", "cacbcabcdccda", "abda", "dcaaddaddadaddcdbbbbb", "caddb", "bbcdbddbdcbcccabb", "badb", "cabcdccbadbbabbbdbbcdbad", "ddcaddcdbacdcbadbbbbdbbcdc", "b", "cccdcaabdcabcbbcaabababddda", "dcadabcadadcbbcacdaccbb", "abcdbbdbadadacbcccb", "cbacbbacacbabdadc", "acbdacbaacaac", "aacbccbbbbcacddaa", "bcdbaab", "caadcaadbaadababddcbbabaacbdd", "badbbacbabcdabbcaddddc", "abcdbbdbadadcbcb", "abcdbbdbadadcbccb", "abd", "abcdd", "addcabbbdabaa", "abcdbbdbd"})
	// res := getAns([][]int{[]int{1, 1, 1, 1}, []int{1, 1, 0, 1}, []int{1, 0, 1, 0}})
	// res := merge([][]int{[]int{0, 4}, []int{1, 4}})
	// [1,3],[2,6],[8,10],[15,18]
	// res := merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}})
	// [[2,3],[5,5],[2,2],[3,4],[3,4]]
	// expected [[2,4],[5,5]]
	// res := merge([][]int{[]int{1, 4}, []int{0, 0}})
	// [[2,3],[2,2],[3,3],[1,3],[5,7],[2,2],[4,6]]
	// expected [[1,3],[4,7]]
	res := merge([][]int{[]int{2, 3}, []int{5, 5}, []int{2, 2}, []int{3, 4}, []int{3, 4}})
	fmt.Println(res)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	// m := make(map[int]int)
	// intervalArray := make([]int, 0)
	res := make([][]int, 0)
	intervals = quickSort(intervals, 0, len(intervals)-1)
	// fmt.Println(intervals)
	for i := 0; i < len(intervals); i++ {
		item := intervals[i]
		if len(res) == 0 || res[len(res)-1][1] < item[0] {
			res = append(res, item)
		} else {
			res[len(res)-1][1] = int(math.Max(float64(res[len(res)-1][1]), float64(item[1])))
		}
	}

	return res
}

func quickSort(intervalArray [][]int, s int, e int) [][]int {
	if s >= e {
		return intervalArray
	}

	// copy()
	pivot := intervalArray[s][0]
	left := s
	right := e
	for s != e {
		for s < e && intervalArray[e][0] >= pivot {
			e--
		}
		for s < e && intervalArray[s][0] <= pivot {
			s++
		}
		if s < e {
			intervalArray[s], intervalArray[e] = intervalArray[e], intervalArray[s]
		}
	}
	temp := intervalArray[left]
	intervalArray[left] = intervalArray[s]
	intervalArray[s] = temp
	// fmt.Println("******")
	// fmt.Println(s, e)
	// fmt.Println(left, right)
	quickSort(intervalArray, left, s-1)
	quickSort(intervalArray, s+1, right)
	return intervalArray
}
