// https://leetcode.com/problems/flipping-an-image/
/*
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

Example 1:

Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
Example 2:

Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Notes:

1 <= A.length = A[0].length <= 20
0 <= A[i][j] <= 1
*/
// 这个题涉及到in place的reverse 要考虑长度是奇数还是偶数的情况
func flipAndInvertImage(A [][]int) [][]int {
	if len(A) <= 0 {
		return A
	}
	r := len(A)
	c := len(A[0])
	// cl := c/2 + 1
	// if len(A[0])%2 == 0 {
	// 	cl = c / 2
	// }
	for i := 0; i < r; i++ {
		// notice there is (c+1)/2 to replace the cl
		for j := 0; j < (c+1)/2; j++ {
			A[i][j] ^= 1
			if j < c/2 {
				A[i][c-j-1] ^= 1
				A[i][j], A[i][c-1-j] = A[i][c-1-j], A[i][j]
			}
			/*
				when c is odd, we could use a tmp to reverse to remove the if j < c/2
				tmp := A[i][j] ^ 1
				A[i][j] = A[i][c-j-1] ^ 1
				A[i][c-j-1] = tmp
			*/
		}
	}
	// fmt.Println(A)
	return A
}
