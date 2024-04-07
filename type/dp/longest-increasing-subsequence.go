/*
 * @Author: hongwei.sun
 * @Date: 2024-04-07 17:10:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-07 21:49:20
 * @Description: file content
 */
//  https://leetcode.com/problems/longest-increasing-subsequence/description/
/*
Given an integer array nums, return the length of the longest strictly increasing 
subsequence.

Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1
 

Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104
 

Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/
// 不能直接的brute force O(n^2)
/* 
因为要的是递增，双层循环是 for i for j, j 都是和i去比。 而这里不仅j要大于i是要找j++往前时和j的对比
也不能说拿个数记cur 或者上一个值，因为这个值是会变的，也许cur是一个大于i的值 但不一定就是递增里最大的
比如 [0,1,0,3,2,3]， i = 0, 下一个大于i的cur 是1 然后cur为1， 然后下一个大于1的是3，cur又为3
这样LIS就是3。 但其实应该是 0,1,2,3。
可以用dp， dp[i] 代表在i前面的LIS是多少
*/
func lengthOfLIS(nums []int) int {
	// resMax := 1
	// dp[i] 代表包括i的LIS是多少
	// 不能用dp[i]代表 i前面的LIS是什么 因为这么定义就和i没有关系了，i是否大于j 不能推导出 dp[i] = dp[j]+1 or dp[i] = dp[j]
	res := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }

 // use the binary search to make the TC O(nlogn)
/* It means we replace the corresponding number with the smaller number and
add the bigger number at the bottome of dp array
to keep on the following elements could be checked.
It's a hard to understand solution.
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	const MAX = math.MaxInt64
	f := make([]int, len(nums))
	// init the every element is MAX in f array as dp
	for i := 1; i < len(f); i++ {
		f[i] = MAX
	}
	f[0] = nums[0]
	// binary search: find the smallest element which is bigger or equal to nums[i]
	// the index is j and replace the f[j] with nums[i]
	// it's to find the longest increasing subsequence.
	for i := 1; i < len(nums); i++ {
		j := sort.Search(len(f), func(j int) bool { return f[j] >= nums[i] })
		f[j] = nums[i]
	}
	// the f length which is not MAX is the result
	for i := len(f) - 1; i >= 0; i-- {
		if f[i] != MAX {
			return i + 1
		}
	}
	return 0
}
