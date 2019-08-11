package main

import (
	"fmt"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	l := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[l] = nums1[i]
			l--
			i--
		} else {
			nums1[l] = nums2[j]
			l--
			j--
		}
	}
	for i >= 0 {
		nums1[l] = nums1[i]
		l--
		i--
	}
	for j >= 0 {
		nums1[l] = nums2[j]
		l--
		j--
	}
	return
}
