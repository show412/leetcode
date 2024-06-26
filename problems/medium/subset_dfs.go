/*
 * @Author: hongwei.sun
 * @Date: 2024-02-19 17:10:19
 * @LastEditors: your name
 * @LastEditTime: 2024-02-19 17:31:12
 * @Description: file content
 */
// https://leetcode.com/problems/subsets/
func subsets(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	entry := make([]int, 0, len(nums))

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
	}
}


func subsets(nums []int) [][]int {
	results := make([][]int, len(nums))
	if nums == nil {
		return results
	}
	if len(nums) == 0 {
		results = append(results, nums)
		return results
	}
	// sort.Sort(nums)
	subsets := make([]int, len(nums))
	dfs(&subsets, nums, 0, &results)
	return results
}
func dfs(subsets *[]int, nums []int, start int, *results [][]int) {
	cpy := make([]int, len(*subsets))
	copy(cpy, *subsets)
	*results = append(*results, cpy)
	for i := start; i < len(nums); i++ {
		*subsets = append(*subsets,nums[i])
		dfs(subsets, nums, i+1, results)
        *subsets = (*subsets)[:len(*subsets)-1]
	}
}

