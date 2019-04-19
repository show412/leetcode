// Complete the jumpingOnClouds function below.
// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem

func jumpingOnClouds(c []int32) int32 {
	// intsalize
	var jump int32 = 0
	// function
	// 0 0 0 1 0 0
	for i := 0; i < len(c); i++ {
		if i == len(c)-1 {
			break
		}
		if i == len(c)-3 || i == len(c)-2 {
			jump++
			break
		}
		if c[i+1] == 1 || c[i+2] == 0 {
			i++
		}
		jump++
	}
	// result
	return jump
}
