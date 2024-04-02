/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 11:22:04
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 11:36:24
 * @Description: file content
 */
  // can practice with this problem https://leetcode.com/problems/sort-an-array/solutions/
/*
思路: 数组切分两块，分别排序，最后merge
TC 平均是 O(nlogn) SC O(n)
*/
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
