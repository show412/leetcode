
func maxSum(n int, array []int) int {
	if len(array) == 0 {
		return 0
	}
	maxSum := array[0]
	max := array[0]
	for i := 0; i < n; i++ {
		if max <= 0 {
			max = array[i]
		} else {
			max += array[i]
		}
		if max > maxSum {
			maxSum = max
		}
	}
	return maxSum
}
