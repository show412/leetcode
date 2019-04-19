// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var max int64 = 0
	array := make([]int64, n)
	for i := 0; i < len(queries); i++ {
		for j := queries[i][0] - 1; j < queries[i][1]; j++ {
			array[j] += int64(queries[i][2])
			if max < array[j] {
				max = array[j]
			}
		}
	}
	return max
}

