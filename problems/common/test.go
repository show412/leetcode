package main

import (
	"fmt"
	"strings"
	// "math"
	// "sort"
)

func main() {
	// [1,1],[1,3],[3,1],[3,3],[4,1],[4,3]
	// [3,2],[0,0],[3,3],[3,4],[4,4],[2,1],[4,3],[1,0],[4,1],[0,2]
	// a := [][]int{[]int{3, 2}, []int{0, 0}, []int{3, 3}, []int{3, 4}, []int{4, 4}, []int{2, 1}, []int{4, 3}, []int{1, 0}, []int{4, 1}, []int{0, 2}}
	// b := pickFruits([]int{1, 2, 1, 2, 1, 2, 1})
	// b := pickFruits([]int{1, 2, 1, 3, 4, 3, 5, 1, 2})
	// v32 := "-354634382"
	// if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
	// 	fmt.Printf("%T, %v\n", s, s)
	// }
	// if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
	// 	fmt.Printf("%T, %v\n", s, s)
	// }

	// v64 := "-354634382"
	// if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
	// 	// fmt.Println(err)
	// 	fmt.Printf("%T, %v\n", s, s)
	// } else {
	// 	fmt.Println(err)
	// }
	// if s, err := strconv.ParseInt("4", 2, 64); err == nil {
	// 	fmt.Printf("%T, %v\n", s, s)
	// } else {
	// 	fmt.Println(err)
	// }
	// sss := "4"
	// var binString string
	// s := strconv.Itoa(100)
	// if err == nil {
	// fmt.Println(s)
	// }

	// for _, c := range sss {
	// 	fmt.Println(c)
	// 	binString = fmt.Sprintf("%s%b", binString, c)
	// }
	// fmt.Println(binString)
	// fmt.Println(reflect.TypeOf(binString).String())
	// fmt.Println(b)
	res := queryString("0110", 4)
	fmt.Println(res)
}

func queryString(str string, n int) string {
	// Write your code here.
	if str == "" {
		return "no"
	}
	for i := 0; i <= n; i++ {
		bi := fmt.Sprintf("%b", i)
		// s := strconv.Itoa(bi)
		if strings.Contains(str, bi) == false {
			return "no"
		}
	}
	return "yes"
}

func pickFruits(arr []int) int {
	// Write your code here.
	res := []int{}
	m := make(map[int]int)
	start := 0
	// [1,2,1,3,4,3,5,1,2]
	for end := 0; end < len(arr); end++ {
		if _, ok := m[arr[end]]; ok {
			m[arr[end]] += 1
		} else {
			m[arr[end]] = 1
		}
		for len(m) > 2 {
			m[arr[start]] -= 1
			if m[arr[start]] == 0 {
				delete(m, arr[start])
			}
			if start < end {
				start++
			}
		}
		if len(m) == 2 {
			// fmt.Println(m)
			// fmt.Println(end)
			// fmt.Println(start)
			res = append(res, end-start+1)
		}
	}

	max := 1
	for i := 0; i < len(res); i++ {
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}
