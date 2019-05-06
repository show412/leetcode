func budgetShopping(n int32, bundleQuantities []int32, bundleCosts []int32) int32 {
    // Write your code here
		// It's a backpack question
		// use DP
		// bool f[][] = make([int32][int32]bool)
		f := make([][]bool, len(A)+1)
		for i := 0; i < len(A)+1; i++ {
			f[i] = make([]bool, m+1)
		}
		for i := 0; i<=len(bundleQuantities) ; i++ {
			for j:=0; j<=n; j++ {
				f[i][j] = false
			}
		}

				f[0][0] = true;
				for i := 1; i<=len(bundleQuantities) ; i++ {
					for j:=0; j<=n; j++ {
						f[i][j] = f[i - 1][j];
                if (j >= bundleQuantities[i-1] && f[i-1][j - bundleQuantities[i-1]]) {
                    f[i][j] = true;
                }
					}
				}


        for (int i = m; i >= 0; i--) {
            if (f[A.length][i]) {
                return i;
            }
        }

        return 0;
}
