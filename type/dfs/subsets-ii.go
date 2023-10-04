import "sort"

// https://leetcode.com/problems/subsets-ii/
/*
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	entry := make([]int, 0, len(nums))
	// sort firstly
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
		// when the next num is equal to current num, it should skip it
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}
}