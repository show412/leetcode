/*
 * @Author: hongwei.sun
 * @Date: 2024-02-19 14:40:16
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-17 22:54:11
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
	// see which level current is in
	// l := 0
	backtrack(0, nums, &subset, &ret)
	return ret
}

// start is which index start to backtrace from, subset is
func backtrack(start int, nums []int, subset *[]int, ret *[][]int) {
	// append this loop result into ret
	cpy := make([]int, len(*subset))
	copy(cpy, *subset)
	*ret = append(*ret, cpy)
	// go to next level
	// l++
	// fmt.Println("*****")
	// fmt.Println("level,", l)
	// fmt.Println("ret is,", *ret)
	for i := start; i < len(nums); i++ {
		*subset = append(*subset, nums[i])
		// check which i is now
		// fmt.Println("current i is,", i)
		// fmt.Println("subset is,", *subset)
		backtrack(i+1, nums, subset, ret)
		// back to last level
		*subset = (*subset)[:len(*subset)-1]
		// fmt.Println("backtrack subset is,", *subset)
	}
	// go to last level
	// l--
}

// another writing
var res [][]int
var candidate []int

func subsets(nums []int) [][]int {
   res = make([][]int, 0)
   candidate = make([]int, 0)
   dfs(nums, 0)
   return res
}

func dfs(nums []int, start int) {
   cpy := make([]int, len(candidate))
   copy(cpy, candidate)
   res = append(res, cpy)
   for i := start; i < len(nums); i++ {
	   candidate = append(candidate, nums[i])
	   dfs(nums, i+1)
	   candidate = candidate[:len(candidate)-1]
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

/*
*****
level, 1
ret is, [[]]
current i is, 0
subset is, [1]
*****
level, 2
ret is, [[] [1]]
current i is, 1
subset is, [1 2]
*****
level, 3
ret is, [[] [1] [1 2]]
current i is, 2
subset is, [1 2 3]
*****
level, 4
ret is, [[] [1] [1 2] [1 2 3]]
backtrack subset is, [1 2]
backtrack subset is, [1]
current i is, 2
subset is, [1 3]
*****
level, 3
ret is, [[] [1] [1 2] [1 2 3] [1 3]]
backtrack subset is, [1]
backtrack subset is, []
current i is, 1
subset is, [2]
*****
level, 2
ret is, [[] [1] [1 2] [1 2 3] [1 3] [2]]
current i is, 2
subset is, [2 3]
*****
level, 3
ret is, [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3]]
backtrack subset is, [2]
backtrack subset is, []
current i is, 2
subset is, [3]
*****
level, 2
ret is, [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
backtrack subset is, []
*/