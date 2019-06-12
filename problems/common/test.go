package main

import (
	"fmt"
	"strconv"
)

func main() {
	// [[0, 30],[5, 10],[15, 20]]
	// [[9,10],[4,9],[4,17]]
	res := generateParenthesis(6)
	fmt.Println(res)
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	nums := make([]int, 0)
	visit := make(map[int]bool)
	uniq := make(map[string]bool)
	res := make([][]int, 0)
	entry := make([]int, 0)
	result := make([]string, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, 1)
	}
	for i := 0; i < n; i++ {
		nums = append(nums, -1)
	}
	dfs(nums, visit, entry, &res, uniq)
	for _, item := range res {
		str := ""
		for _, i := range item {
			if i == 1 {
				str += "("
			} else if i == -1 {
				str += ")"
			}
		}
		result = append(result, str)
	}
	return result
}

func dfs(nums []int, visit map[int]bool, entry []int, res *[][]int, uniq map[string]bool) {
	if len(entry) == len(nums) {
		cpy := make([]int, len(entry))
		copy(cpy, entry)
		str := ""
		for i := 0; i < len(cpy); i++ {
			str += strconv.Itoa(cpy[i])
		}
		// fmt.Println(str)
		if v, ok := uniq[str]; ok && v == true {
			return
		}
		*res = append(*res, cpy)
		uniq[str] = true
		return
	}
	for i := 0; i < len(nums); i++ {
		// if i != 0 && nums[i-1] == nums[i] {
		// 	continue
		// }
		if v, ok := visit[i]; ok && v == true {
			continue
		}
		entry = append(entry, nums[i])
		visit[i] = true
		if isValid(entry) == true {
			dfs(nums, visit, entry, res, uniq)
		}
		entry = entry[:len(entry)-1]
		visit[i] = false
	}
}

func isValid(arr []int) bool {
	cur := 0
	for _, i := range arr {
		cur += i
		if cur < 0 {
			return false
		}
	}
	return true
}
