/*
 * @Author: hongwei.sun
 * @Date: 2024-02-19 14:40:16
 * @LastEditors: your name
 * @LastEditTime: 2024-02-19 18:24:34
 * @Description: file content
 */
//  https://leetcode.com/problems/subsets/description/
//  Given an integer array nums of unique elements, return all possible
//  subsets
//   (the power set).

//  The solution set must not contain duplicate subsets. Return the solution in any order.

//  Example 1:

//  Input: nums = [1,2,3]
//  Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//  Example 2:

//  Input: nums = [0]
//  Output: [[],[0]]

//  Constraints:

//  1 <= nums.length <= 10
//  -10 <= nums[i] <= 10
//  All the numbers of nums are unique.
/*
it's one typical backtracking or dfs, convert to one tree to travers
                      []
		    1/                 \no 1
          [1]                  []
	  2	/     \no 2         2/    \no 2
    [1,2]     [1]          [2]    []
  3/ \no3    3/ \        3/ \    3/ \
[1,2,3][1,2][1,3] [1] [2,3] [2] [3] []

简单来说就是从集合里 包括i 和不包括i的 一个个的去枚举
看code 仔细想的话 在哪里加到结果里的逻辑其实不是很好理解 目前看也只能是记住
*/

// practice
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	subset := make([]int, 0)
	backtrack(0, nums, &subset, &ret)
	return ret
}

// start is which index start to backtrace from, subset is
func backtrack(start int, nums []int, subset *[]int, ret *[][]int) {
	// append this loop result into ret
	cpy := make([]int, len(*subset))
	copy(cpy, *subset)
	*ret = append(*ret, cpy)

	for i := start; i < len(nums); i++ {
		*subset = append(*subset, nums[i])
		backtrack(i+1, nums, subset, ret)
		// back to last level
		*subset = (*subset)[:len(*subset)-1]
	}
}

/*
                      []
		    1/                 \no 1
           [1]                  []
	2	/      \no 2         2/    \no 2
    [1,2]       [1]          [2]    []
  3/   \no3    3/ \        3/ \    3/ \
[1,2,3]    [1,3] [1]    [2,3] [2] [3] []
*/