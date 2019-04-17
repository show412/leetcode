// https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=warmup&isFullScreen=true
// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	m := make(map[int32]int32)
	var count int32 = 0
	var i int32
	for i = 0; i < n; i++ {
		if _, ok := m[ar[i]]; ok {
			m[ar[i]]++
			if m[ar[i]]%2 == 0 {
				count++
			}
		} else {
			m[ar[i]] = 1
		}
	}
	return count
}
