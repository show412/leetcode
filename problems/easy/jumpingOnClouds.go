// Complete the jumpingOnClouds function below.
// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem

func jumpingOnClouds(c []int32) int32 {
	// f[i] represent that when it's at i, it needs how many steps at least
	var f []int32
	// intsalize
	f[0] = 0
	// function
	for i := 1 ; i < len(c); i++ {
		if c[i] == 1 {
			if c[i-1] == 1 {
				f[i] = 
			} else {

			}
			f[i] = f[i-1] + 2
		} else {
			if f[i-1] + 1	< f[i-1]
		}

	}
	// result

}