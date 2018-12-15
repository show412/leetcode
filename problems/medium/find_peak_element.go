// https://leetcode.com/problems/find-peak-element/
func findPeakElement(nums []int) int {
  start := 0
  end := len(nums) - 1
  mid := 0
  for start + 1 < end {
    mid = start + (end - start)/2
    if nums[mid]> nums[mid-1] && nums[mid]>nums[mid+1] {
      return mid
    } else if nums[mid]> nums[mid-1] && nums[mid]<nums[mid+1] {
      start = mid 
    } else if nums[mid]< nums[mid-1] && nums[mid]>nums[mid+1] {
      end = mid
    } else if nums[mid]< nums[mid-1] && nums[mid]<nums[mid+1] {
      start = mid 
    }
  }
  
  if nums[start] > nums[end] {
    return start
  } else {
    return end
  }
  
}