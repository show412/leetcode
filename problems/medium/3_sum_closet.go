import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/
// Sort the vector and then no need to run O(N^3) algorithm as each index has a direction to move.

// The code starts from this formation.

// ----------------------------------------------------
// ^  ^                                               ^
// |  |                                               |
// |  +- second                                     third
// +-first
// if nums[first] + nums[second] + nums[third] is smaller than the target, we know we have to increase the sum. so only choice is moving the second index forward.

// ----------------------------------------------------
// ^    ^                                             ^
// |    |                                             |
// |    +- second                                   third
// +-first
// if the sum is bigger than the target, we know that we need to reduce the sum. so only choice is moving 'third' to backward. of course if the sum equals to target, we can immediately return the sum.

// ----------------------------------------------------
// ^    ^                                          ^
// |    |                                          |
// |    +- second                                third
// +-first
// when second and third cross, the round is done so start next round by moving 'first' and resetting second and third.

// ----------------------------------------------------
//   ^    ^                                           ^
//   |    |                                           |
//   |    +- second                                 third
//   +-first
// while doing this, collect the closest sum of each stage by calculating and comparing delta. Compare abs(target-newSum) and abs(target-closest). At the end of the process the three indexes will eventually be gathered at the end of the array.

// ----------------------------------------------------
//                                          ^    ^    ^
//                                          |    |    `- third
//                                          |    +- second
//                                          +-first
// if no exactly matching sum has been found so far, the value in closest will be the answer.

// type arryNums []int

// func (a arryNums) Len() int           { return len(a) }
// func (a arryNums) Less(i, j int) bool { return a[i] < a[j] }
// func (a arryNums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	closet := nums[0] + nums[1] + nums[2]
	var curSum int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		z := len(nums) - 1
		curSum = nums[i] + nums[j] + nums[z]
		for j < z {
			if curSum == target {
				return target
			}
			if math.Abs(float64(target-curSum)) < math.Abs(float64(target-closet)) {
				closet = curSum
			}
			if curSum < target {
				j++
			} else {
				z--
			}
		}
	}
	return curSum
}
