// a := []int{6,1,2,7,9,3,4,5,10,8}
// becuase the slice in go is pointer, so if the nums is not changed by capicity, the nums always point to the same nums

func quick_sort(nums []int, l int, h int) []int {
	pivot := nums[l]
	left := l
	right := h
	if left >= right {
		return nil
	}
	for l != h {
		// notice !! it must that the right go firstly
		for l < h && nums[h] >= pivot {
			h--
		}
		for l < h && nums[l] <= pivot {
			l++
		}
		if l < h {
			nums[l], nums[h] = nums[h], nums[l]
		}
	}

	nums[left] = nums[l]
	nums[l] = pivot
	quick_sort(nums, left, l-1)
	quick_sort(nums, l+1, right)
	return nums
}