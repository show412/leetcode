package main

import (
	"fmt"
	"math"
	"strings"
	// "math"
	// "sort"
)

func main() {

	res := minimumTotal([][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}})
	fmt.Println(res)
}

func minimumTotal(triangle [][]int) int {
	// write your code here
	// DP to solve the issue
	f := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		f[i] = make([]int, len(triangle[i]))
	}
	res := int(^uint(0) >> 1)
	f[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		f[i][0] += f[i-1][0] + triangle[i][0]
	}

	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i]); j++ {
			if j == len(triangle[i])-1 {
				f[i][j] = triangle[i][j] + f[i-1][j-1]
			} else {
				f[i][j] = triangle[i][j] + int(math.Min(float64(f[i-1][j-1]), float64(f[i-1][j])))
			}

		}
	}
	bottom := len(triangle) - 1
	// fmt.Println(triangle[bottom])
	// fmt.Println(f)
	for j := 0; j < len(triangle[bottom]); j++ {
		if res > f[bottom][j] {
			res = f[bottom][j]
		}
	}
	return res
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
