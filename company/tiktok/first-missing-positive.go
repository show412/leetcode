/*
 * @Author: hongwei.sun
 * @Date: 2024-04-09 15:26:15
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 16:48:48
 * @Description: file content
 */
// https://leetcode.com/problems/first-missing-positive/description/
/*
Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

 

Example 1:

Input: nums = [1,2,0]
Output: 3
Explanation: The numbers in the range [1,2] are all in the array.
Example 2:

Input: nums = [3,4,-1,1]
Output: 2
Explanation: 1 is in the array but 2 is missing.
Example 3:

Input: nums = [7,8,9,11,12]
Output: 1
Explanation: The smallest positive integer 1 is missing.
 

Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
*/
// TC O(n) and SC O(1)
/*
1, 把所有是负数的都置为0， 因为我们不用care 负数
2，miss的smallest positive number 一定是在[1-len(nums)+1]
我们的要做的就是找到1-len(nums)+1 有哪些不存在数组里
3，理想情况我们希望nums是 [1.... len(nums)] 1在index 0， 2在index 1 那就可以利用这个来处理
nums是没有排序的，因为我们已经把所有负值置为0了，那可以利用数组的index和负值来标识相应的index对应的值是不是存在在数组里
比如nums第一个数是3， 它可以 通过 数组index 2是不是负值来标识， 如果是负值说明3是存在在这个数组的
4，然后在遍历 [1-len(nums)+1] 看大于等于0的最小nums index是什么 这个就是结果
*/
func firstMissingPositive(nums []int) int {
	for index, num := range nums {
		if num < 0 {
			nums[index] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if index >= 0 && index < len(nums) {
			/*
			nums[index] < 0 说明被标识过了 nums里有重复的正整数
			*/
			if nums[index] > 0 {
				nums[index] = -nums[index]
			} else if nums[index] == 0 {
				// 等于0说明原来是负数或者0， 标识一个超出范围的负值，说明相应的正整值存在
				// 同时当遍历到这个位置的时候又不用再去找 所以要标识一个超出范围的负值
				nums[index] = -(len(nums)+1)
			}
		}
	}
	res := len(nums) + 1
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] >= 0 {
			res = i
			break
		}
	}
	return res
}



//TC O(n) and SC O(n)
func firstMissingPositive(nums []int) int {
    m := make(map[int]bool)
	for _, num := range nums {
		if num > 0 {
			m[num] = true
		}
	}
	res := len(nums)+1
	for i := 1; i <= len(nums); i++ {
		if m[i] == false {
			res = i
            break
		}
	} 
	return res
}