func permute(nums []int) [][]int {
	res := [][]int{}
	visit := make(map[int]bool)
	entry := []int{}
	dfs(nums, visit, &entry, &res)
	return res
}

func dfs(nums []int, visit map[int]bool, entry *[]int, res *[][]int) {
	if len(*entry) == len(nums) {
		cpy := make([]int, len(*entry))
		copy(cpy, *entry)
		*res = append(*res, cpy)
		return
	}

	for i := 0; i < len(nums); i++ {

		if visit[nums[i]] == true {
			continue
		} else {
			visit[nums[i]] = true
			*entry = append(*entry, nums[i])
			dfs(nums, visit, entry, res)
			visit[nums[i]] = false
			*entry = (*entry)[:len(*entry)-1]
		}
	}
	return
}
