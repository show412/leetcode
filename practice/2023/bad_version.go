/*
 * @Author: hongwei.sun
 * @Date: 2023-10-05 14:42:02
 * @LastEditors: your name
 * @LastEditTime: 2023-10-05 15:44:38
 * @Description: file content
 */
//  https://leetcode.com/problems/first-bad-version/
//  binary search
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
/*
It's one typical binary search writing
notice, left = mid + 1
*/
func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}