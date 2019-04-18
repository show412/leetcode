// Complete the formingMagicSquare function below.
// It turns out there are only 8 possible solutions to the problem
// as given on the blog post
// https://mindyourdecisions.com/blog/2015/11/08/how-many-3x3-magic-squares-are-there-sunday-puzzle/

// the question is in https://www.mathblog.dk/hackerrank-forming-a-magic-square/
// One common magic grid althgram is https://www.geeksforgeeks.org/magic-square/
func formingMagicSquare(s [][]int32) int32 {
	allPossible := [][9]int32{
		[9]int32{8, 1, 6, 3, 5, 7, 4, 9, 2},
		[9]int32{6, 1, 8, 7, 5, 3, 2, 9, 4},
		[9]int32{4, 9, 2, 3, 5, 7, 8, 1, 6},
		[9]int32{2, 9, 4, 7, 5, 3, 6, 1, 8},
		[9]int32{8, 3, 4, 1, 5, 9, 6, 7, 2},
		[9]int32{4, 3, 8, 9, 5, 1, 2, 7, 6},
		[9]int32{6, 7, 2, 1, 5, 9, 8, 3, 4},
		[9]int32{2, 7, 6, 9, 5, 1, 4, 3, 8}}
	flatS := [9]int32{s[0][0], s[0][1], s[0][2], s[1][0], s[1][1], s[1][2], s[2][0], s[2][1], s[2][2]}
	adjustNum := []int32{}
	for i := 0; i < 8; i++ {
		var sum int32 = 0
		for j := 0; j < 9; j++ {
			if (allPossible[i][j] - flatS[j]) > 0 {
				sum += int32((allPossible[i][j] - flatS[j]))
			} else {
				sum += int32((flatS[j] - allPossible[i][j]))
			}
		}
		adjustNum = append(adjustNum, sum)
	}
	var minCost int32 = 1000
	for _, v := range adjustNum {
		if minCost > v {
			minCost = v
		}
	}
	return minCost
}
