/*
 * @Author: hongwei.sun
 * @Date: 2024-03-27 21:44:46
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-27 22:13:41
 * @Description: file content
 */
import "math"

// https://leetcode.com/problems/median-of-two-sorted-arrays/
/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
/*
it requires TC O(log (m+n)), heap TC is O(nlogn)
so heap is not a good solution, we know binary search TC is O(logn)
it's a common solution for find kth elements in sorted array
一步步的递归下去，两个sorted arrary nums1 and nums2 要找kth
compare nums1 or nums2中的 前1/2 k个元素 把较小的舍去 这样就不用一个个的用双指针去比较了
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	// this step should be thought out
	if (len1+len2)%2 == 1 {
		return float64(findKth(nums1, nums2, (len1+len2)/2+1))
	} else {
		return float64(findKth(nums1, nums2, (len1+len2)/2)+findKth(nums1, nums2, (len1+len2)/2+1)) / 2
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	// the condition is very important
	// it's the condition of out of the rescusive
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	/*
		比较第k/2个 如果nums1的k/2比 nums2的k/2小
		说明整体的第k个肯定不在nums1的前k/2个里 所以可以舍掉nums1的前k/2个元素
		对nums2同理
		一些极限情况
		如果nums1剩下的没有k/2个元素了 这些元素不能舍去
		这时候可以说去来回比较nums1[0] nums1[end]和nums2中的情况 但是没有必要
		因为这时候中位数肯定不在nums1的前k/2个里面（因为nums1不够k/2个了）
		而且不可能同时有nums1和nums2都不满足剩余k/2个元素的情况 否则的话第k个是怎么来的呢
		就没有意义了
	*/
	nums1HalfK := math.MaxInt64
	nums2HalfK := math.MaxInt64
	if len(nums1) >= (k / 2) {
		nums1HalfK = nums1[(k/2)-1]
	}
	if len(nums2) >= (k / 2) {
		nums2HalfK = nums2[(k/2)-1]
	}

	if nums1HalfK <= nums2HalfK {
		return findKth(nums1[k/2:], nums2, k-k/2)
	} else {
		return findKth(nums1, nums2[k/2:], k-k/2)
	}
}

