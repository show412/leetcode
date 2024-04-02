/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 11:10:06
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 11:16:05
 * @Description: file content
 */
 // can practice with this problem https://leetcode.com/problems/sort-an-array/solutions/
/*
思路: 每次的元素都排序到当前数组中，像插扑克牌一样
TC 平均是 O(n^2)  最好是O（n）
*/
func insertSort(nums []int) []int {
	for i := 1; i <= len(nums) - 1; i++ {
		num := nums[i]
		j := i-1
		for  j >= 0 &&  nums[j] > num {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = num 
	}
	return nums
}
