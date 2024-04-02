/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 11:06:50
 * @Description: file content
 */
// can practice with this problem https://leetcode.com/problems/sort-an-array/solutions/
/*
思路: 两两比较把 最大的往后挪
TC 平均是 O(n^2)  最好是O（n）
*/
func bubbleSort(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		swap := false
		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
	return nums
}