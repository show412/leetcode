// 在n个物品中挑选若干物品装入背包，最多能装多满？
// 假设背包的大小为m，每个物品的大小为A[i]
// https://www.jiuzhang.com/solutions/backpack/
func backPack(m int, A []int) int {
	// f[i][j] means first i items could fill pack whose size is j
	// It's the way to define a two-dimension array
	f := make([][]bool, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i < len(A); i++ {
		for j := 0; j <= m; j++ {
			// 两种方案
			// 1， 前i-1个就组成了j了
			// 2， 前i-1个组成了j-A[i] 第i个进来的时候就能组成j了 因为第i个的大小是A[i]，
			f[i+1][j] = f[i][j]
			if j >= A[i] && f[i][j-A[i]] {
				f[i+1][j] = true
			}
		}
	}

	for i := m; i >= 0; i-- {
		if f[len(A)][i] {
			return i
		}
	}
	return 0
}
