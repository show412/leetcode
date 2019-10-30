package main

import (
	"fmt"
	"sort"
)

/*
test case:
[[1,1,4],[9,4,9],[9,1,9],[2,3,5],[4,1,5],[10,4,5]]
33
false

*/

func main() {
	res := subsetsWithDup([]int{1, 2, 2})
	// 998001
	fmt.Println(res)
}
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	entry := make([]int, 0, len(nums))
	sort.Ints(nums)
	backtrack(0, nums, &entry, &res)
	return res
}

func backtrack(start int, nums []int, entry *[]int, res *[][]int) {
	cpy := make([]int, len(*entry))
	copy(cpy, *entry)
	*res = append(*res, cpy)

	for i := start; i < len(nums); i++ {
		*entry = append(*entry, nums[i])
		backtrack(i+1, nums, entry, res)
		*entry = (*entry)[:len(*entry)-1]
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}
}
