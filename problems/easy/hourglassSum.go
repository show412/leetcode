// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	const INT_MAX = int32(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	var max int32 = INT_MIN
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			hourglass := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if max < hourglass {
				max = hourglass
			}
		}
	}
	return max
}
