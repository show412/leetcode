// 假设buff[i][j]的含义为：将第i个数改成j的调整代价。如果改成j之后不符合调整的目标则buff[i][j]=INT_MAX。

// 初始化缓存将第0个数改为j的代价是

// buff[0][j] = abs(A[0]-j);
// 1
// 为了符合调整目标，将第i个数改成j的话，第i-1个数k的可取值范围是[j-target, j+target]。在符合调整目标的情况下的状态转移方程为：

// buff[i][j] = min(buff[i][j], buff[i-1][k]+abs(j-A[i]));


	func  MinAdjustmentCost(A []int, target int) res int {
		// write your code here
		var len int = len(A)

		buff := make([][]int, len(A)+1)
		for i := 0; i < len(A)+1; i++ {
			buff[i] = make([]int, 101)
		}
		// var buff [len][101] int
		const INT_MAX = int32(^uint(0) >> 1)

		for j := 1; j <= 100; j++ {
			buff[0][j] := int(math.Abs(float64(j-A[0])))
		}
		for i := 1; i < len; i++ {
				for j := 1; j <= 100; j++{
					buff[i][j] = INT_MAX
					lo := int(math.Max(float64(1), float64(j - target)))
					hi := int(math.Min(float64(100), float64(j + target)))
					for k := lo; k <= hi; k++ {
							buff[i][j] := int(math.Min(float64(buff[i][j]), float64(buff[i-1][k]+abs(j-A[i]))))
					}
				}
		}
		res := INT_MAX
		for i := 1; i <= 100; i++ {
			res := int(math.Min(float64(res), float64(buff[len-1][i]+ int(math.Abs(float64(j-A[i]))))))
		}
		return res;
	}
