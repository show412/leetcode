/**
 * @param arr: The array you should find shortest subarray length which has duplicate elements.
 * @return: Return the shortest subarray length which has duplicate elements.
 */
// It's O(n^2)
func getLength(arr []int) int {
	// Write your code here.
	// INI_MAX = int(^uint(0)>>1)
	result := -1
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				if result == -1 || (j-i+1) < result {
					result = j - i + 1
				}
			}
		}
	}
	return result
}

// It's O(n) with hash
func getLength(arr []int) int {
	// Write your code here.
	m := make(map[int]int, len(arr))
	result := -1
	for i := 0; i < len(arr)-1; i++ {
		if v, ok := m[arr[i]]; ok {
			if result == -1 || (i-v+1) < result {
				result = i - v + 1
			}
		}
		m[arr[i]] = i
	}
	return result
}
