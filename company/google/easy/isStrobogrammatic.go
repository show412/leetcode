// https://leetcode.com/problems/strobogrammatic-number/
/*
A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

Write a function to determine if a number is strobogrammatic. The number is represented as a string.

Example 1:

Input:  "69"
Output: true
Example 2:

Input:  "88"
Output: true
Example 3:

Input:  "962"
Output: false
*/

/*
	use cases:
	"968018661806000118986811000908199810896"
	"96"
	"88"
	"101"
*/
// 没有说num的范围 不像rotated_digits那道题的条件是0 - 10000
func isStrobogrammatic(num string) bool {
	if num == "0" {
		return true
	}
	rotated := make([]byte, 0)
	for i := len(num) - 1; i >= 0; i-- {
		digit := num[i]
		if digit == '6' {
			rotated = append(rotated, '9')
			continue
		}
		if digit == '9' {
			rotated = append(rotated, '6')
			continue
		}
		if digit == '8' {
			rotated = append(rotated, '8')
			continue
		}
		if digit == '1' {
			rotated = append(rotated, '1')
			continue
		}
		if digit == '0' {
			rotated = append(rotated, '0')
			continue
		}
		return false
	}

	if string(rotated) == num {
		return true
	}
	return false
}
