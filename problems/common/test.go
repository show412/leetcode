package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := numSquares(12)
	// [1, 1, 4, 2, 1, 1, 0, 0]
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
func numSquares(n int) int {
	if n == 1 || n == 4 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	var list []int
	for i := 1; i < n; i++ {
		if i*i > n {
			break
		}
		list = append(list, i)
	}
	if list[len(list)-1]*list[len(list)-1] == n {
		return 1
	}
	var sub []int
	cnt := 0
	var res [][]int
	dfs(list, 0, sub, cnt, &res, n)

	min := int(^uint(0) >> 1)
	for i := 0; i < len(res); i++ {
		if min > len(res[i]) {
			min = len(res[i])
		}
	}
	return min
}

func dfs(list []int, start int, sub []int, cnt int, res *[][]int, max int) {
	if cnt == max {
		var cpy []int
		copy(cpy, sub)
		*res = append(*res, cpy)
	}

	for i := start; i < len(list); i++ {
		sub = append(sub, list[i])
		cnt += list[i] * list[i]
		if cnt > max {
			sub = sub[:len(sub)-1]
			cnt -= list[i] * list[i]
			break
		}
		dfs(list, i, sub, cnt, res, max)
		sub = sub[:len(sub)-1]
		cnt -= list[i] * list[i]
	}

}
