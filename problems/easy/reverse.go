func rotate(nums []int, k int)  {
  if k == 0 || k==len(nums) {
    return
  }
  if k > len(nums) {
    k = k%len(nums)
	}
	// The nums is slice, slice is a pointer which point to one array,
	// so it doesn't need to return nums.
  reverse(nums, 0, len(nums)-k-1)
  reverse(nums, len(nums)-k,len(nums)-1)
  reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
  i := start
  j := end
  for ; i<j; {
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
    i++
    j--
  }
}