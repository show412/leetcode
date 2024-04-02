/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 10:58:15
 * @Description: selection sort
 */
// can practice with this problem https://leetcode.com/problems/sort-an-array/solutions/
/*
思路: 从后到前一个一个找出最大的
TC O(n^2)
*/
func selectionSort(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		maxIndex := 0
		for j := 0; j <= i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
	}
	return nums
}


