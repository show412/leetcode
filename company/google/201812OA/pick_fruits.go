/**
 * @param arr: the arr
 * @return: the length of the longset subarray
 */
func pickFruits(arr []int) int {
	// Write your code here.
	res := []int{}
	m := make(map[int]int)
	start := 0
	// [1,2,1,3,4,3,5,1,2]
	for end := 0; end < len(arr); end++ {
		if _, ok := m[arr[end]]; ok {
			m[arr[end]] += 1
		} else {
			m[arr[end]] = 1
		}
		for len(m) > 2 {
			m[arr[start]] -= 1
			if m[arr[start]] == 0 {
				delete(m, arr[start])
			}
			if start < end {
				start++
			}
		}
		if len(m) == 2 {
			// fmt.Println(m)
			// fmt.Println(end)
			// fmt.Println(start)
			res = append(res, end-start+1)
		}
	}

	max := 1
	for i := 0; i < len(res); i++ {
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}
