/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-06 20:46:47
 * @Description: file content
 */
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
// 定义全局变量 省的用指针传递
var res [][]int 
var entry []int 
func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0, len(nums))
	entry = make([]int, 0, len(nums))
	// sort firstly
	sort.Ints(nums)
	dfs(0, nums)
	return res
}
func dfs(start int, nums []int) {
	// 在这还是得copy 因为entry本质还是一个指针 还是会影响
    cpy := make([]int, len(entry))
    copy(cpy, entry)
	res = append(res, cpy)
	for i := start; i < len(nums); i++ {
		entry = append(entry, nums[i])
		dfs(i+1, nums)
		entry = entry[:len(entry)-1]
		// 后面的如果和当前的相等，那已经在上面的dfs(i+1)里选过了，所以这里要跳过，否则到了下个loop就重复了
		// 注意golang里没有next, 只能用i++
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}
}


// 这是用指针传递的方法
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