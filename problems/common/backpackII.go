// https://www.jiuzhang.com/solution/backpack-ii/
/*There are n items and a backpack with size m.
Given array A representing the size of each item and array V
representing the value of each item.
What's the maximum value can you put into the backpack?
*/
/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @param V: Given n items with value V[i]
 * @return: The maximum value
 */
func backPackII(m int, A []int, V []int) int {
	// write your code here
	// f[i][j] means the maximum value of i goods filled could be j size
	f := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 0
	f[1][0] = 0
	f[0][1] = 0
	for i := 1; i < len(A)+1; i++ {
		for j := 1; j < m+1; j++ {
			// 第i个物品达到的最大值起码可以和前i-1个的价值一样
			f[i][j] = f[i-1][j]
			// 前i-1个组成了j-A[i] 第i个进来的时候就能组成j了 因为第i个的大小是A[i-1]
			// 价值是V[i-1] 因为数组里的key 第i个对应的是i-1
			if A[i-1] <= j {
				f[i][j] = max(f[(i - 1)][j], f[(i - 1)][j-A[i-1]]+V[i-1])
			}
		}
	}
	return f[len(A)][m]
}
func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
