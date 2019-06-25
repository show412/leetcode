// https://leetcode.com/problems/plus-one/
/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.

*/
// TC is O(n), SC is O(1)
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	// res := make([]int, len(digits)+1)
	length := len(digits)
	digits[length-1] = digits[length-1] + 1
	digits = append(digits, 0)
	for i := length - 1; i >= 0; i-- {
		digits[i+1] = digits[i] % 10
		if i == 0 {
			if digits[i] >= 10 {
				digits[i] = digits[i] / 10
			} else {
				digits[i] = 0
			}
		} else {
			digits[i-1] += digits[i] / 10
		}

	}
	if digits[0] != 0 {
		return digits
	}
	return digits[1:]
}

// TC is O(n), SC is O(n)
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	res := make([]int, len(digits)+1)
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	for i := len(digits) - 1; i > 0; i-- {
		res[i+1] = digits[i] % 10
		digits[i-1] += digits[i] / 10
	}
	if digits[0] >= 10 {
		res[1] = digits[0] % 10
		res[0] = digits[0] / 10
	} else {
		res[1] = digits[0]
	}
	if res[0] != 0 {
		return res
	}
	return res[1:]
}

// looks like a little advanced and cool
func plusOne(digits []int) []int {
	var n int = len(digits)
	for i := n - 1; i >= 0; i-- {
		// because if there is one position is less than 9, it could return
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			// otherwise the position is set 0 and continue next loop
			digits[i] = 0
		}
	}
	// if there is no return and arrive here, means no position could be return
	// it must like a 10000 number with n+1 places
	var a = make([]int, n+1)
	a[0] = 1
	return a

}

