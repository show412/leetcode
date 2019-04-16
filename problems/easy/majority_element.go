// https://leetcode.com/problems/majority-element/submissions/
func majorityElement(nums []int) int {
	return nil if nums == nil || len(nums) == 0
	var candidate int
	var count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}