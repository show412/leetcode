/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 11:37:34
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 11:38:08
 * @Description: file content
 */
//  https://leetcode.com/problems/sort-an-array/description/
 //  Given an array of integers nums, sort the array in ascending order and return it.

// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

 

// Example 1:

// Input: nums = [5,2,3,1]
// Output: [1,2,3,5]
// Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).
// Example 2:

// Input: nums = [5,1,1,2,0,0]
// Output: [0,0,1,1,2,5]
// Explanation: Note that the values of nums are not necessairly unique.
 

// Constraints:

// 1 <= nums.length <= 5 * 104
// -5 * 104 <= nums[i] <= 5 * 104
// 思路: 数组切分两块，分别排序，最后merge
// TC 平均是 O(nlogn) SC O(n)
func sortArray(nums []int) []int {
    mergeSort(0, len(nums) - 1, nums)
    return nums
}

func mergeSort(start, end int, nums []int) {
    if (start >= end) {
        return
    }
    middle := start + (end - start) / 2;
    mergeSort(start, middle, nums);
    mergeSort(middle+1, end, nums);
    merge(start, middle, middle+1, end, nums);
}

func merge(start1, end1, start2, end2 int, nums []int) {
    result := make([]int, 0)
    i1, i2 := start1, start2

    for i1 <= end1 && i2 <= end2 {
        if nums[i1] <= nums[i2] {
            result = append(result, nums[i1])
            i1++
        } else {
            result = append(result, nums[i2])
            i2++
        }
    }
    for i1 <= end1 {
        result = append(result, nums[i1])
        i1++
    }
    for i2 <= end2 {
        result = append(result, nums[i2])
        i2++
    }
    
    for _, num := range result {
        nums[start1] = num
        start1++
    }
}
