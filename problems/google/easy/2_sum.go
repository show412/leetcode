// https://leetcode.com/problems/two-sum/
// once pass hash
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
