package main

import (
	"fmt"
	// "math"
)

/*
test case:
[0,0]
-1
true

[0,1,0]
0
false

[23, 2, 6, 4, 7],  k=0
false

[5,0,0]
0
true

{23, 2, 6, 4, 7}, 6
true

[23, 2, 4, 6, 7], 6
true
*/
func main() {
	res := checkSubarraySum([]int{23, 2, 4, 6, 7}, 6)
	fmt.Println(res)
}

func checkSubarraySum(nums []int, k int) bool {
	modMap := make(map[int]int, 0)
	sumArray := make([]int, len(nums))
	sumArray[0] = nums[0]

	if k != 0 {
		modMap[nums[0]%k] = 0
	} else {
		modMap[nums[0]] = 0
	}
	for i := 1; i < len(nums); i++ {
		sumArray[i] = nums[i] + sumArray[i-1]
		if sumArray[i] == k || (k != 0 && sumArray[i]%k == 0) {
			return true
		}
		mod := sumArray[i]
		if k != 0 {
			mod = sumArray[i] % k
		}
		if v, ok := modMap[mod]; ok {
			if i-v > 1 {
				return true
			} else {
				continue
			}
		}
		modMap[mod] = i

	}
	// fmt.Println(sumArray)
	// fmt.Println(modMap)
	return false
}
