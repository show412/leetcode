func kDifference(a []int32, k int32) int32 {
	// Write your code here
	// Time complexity is O(N), use a hashmap to store count of each integer in the nums array.

	// increment count when the a[i] - k or the a[i] + k is in the hashmap.

	if a == nil || len(a) < 2 || k < 0 {
		return 0
	}
	m, count := make(map[int32]int32), 0
	for _, i := range a {
		if k > 0 && m[i] == 0 {
			if m[i-k] > 0 {
				count++
			}
			if m[i+k] > 0 {
				count++
			}
		}
		m[i]++
	}

	if k == 0 {
		for _, v := range m {
			if v > 1 {
				count++
			}
		}
	}

	return count
}

func findPairs(nums []int, k int) int {
	if nums == nil || len(nums) < 2 || k < 0 {
		return 0
	}

	dict, count := make(map[int]int), 0

	for _, n := range nums {
		if k > 0 && dict[n] == 0 {
			if dict[n-k] > 0 {
				count++
			}
			if dict[n+k] > 0 {
				count++
			}
		}
		dict[n]++
	}

	if k == 0 {
		for _, v := range dict {
			if v > 1 {
				count++
			}
		}
	}

	return count
}
