func searchInsert(nums []int, target int) int {
  if len(nums) == 0 || target < nums[0] {
    return 0
  }
  var start = 0
  var end = len(nums) - 1
  var mid = 0
  if target > nums[end] {
    return end + 1
  }
  for {
    if start + 1 >= end {
      break
    }
    mid = start + (end - start)/2
    if nums[mid] == target {
      start = mid
    } else if nums[mid] < target {
      start = mid
    } else if nums[mid] > target {
      end = mid
    } 
  }
  
  if nums[start] == target {
    return start
  }
  if nums[end] == target {
    return end
  }
  
  if nums[mid] < target {
    return mid + 1
  } else {
    return mid
  }
  
  return mid
  
}