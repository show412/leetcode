import "strings"

// Complete the countingValleys function below.
// https://www.hackerrank.com/challenges/counting-valleys/problem
// 遍历这个数组看看有没有和是负数的情况总数
func countingValleys(n int32, s string) int32 {
	var count int32
	var i int32
	var sum int32 = 0
	var intoValley int32 = 0
	str_array := strings.Split(s, "")
	for i = 0; i < n; i++ {
		if sum < 0 && intoValley == 0 {
			intoValley = 1
		}
		if str_array[i] == "U" {
			sum += 1
		} else {
			sum -= 1
		}
		if sum == 0 && intoValley == 1 {
			count++
			intoValley = 0
		}
	}
	return count
}

