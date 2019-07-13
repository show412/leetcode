package main

import (
	"fmt"
	"math"
	// "math"
)

func main() {
	res := new21Game(21, 17, 10)
	fmt.Println(res)
}

func new21Game(N int, K int, W int) float64 {
	// dp[x] = the answer when Alice has x points
	dp := make([]float64, N+W+1)

	for k := K; k <= N; k++ {
		dp[k] = 1.0
	}
	// S = dp[k+1] + dp[k+2] + ... + dp[k+W]
	S := math.Min(float64(N-K+1), float64(W))
	for k := K - 1; k >= 0; k-- {
		dp[k] = S / float64(W)
		S += dp[k] - dp[k+W]
	}
	return dp[0]
}
