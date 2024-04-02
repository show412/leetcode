/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 10:44:11
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 10:44:27
 * @Description: file content
 */
// a := []int{6,1,2,7,9,3,4,5,10,8}
// becuase the slice in go is pointer, so if the nums is not changed by capicity, the nums always point to the same nums
/*
   基本思想： 选定一个pivot， 左边的数都比pivot小，右边的都比pivot大
   然后两边分治 直到只有一个元素的时候
   O（nlogn）
*/
func quick_sort(nums []int, l int, r int) {
	// start end代表这次quicksort的开始和结尾
	// l and r present two pointer to move from left and right
	start := l
	end := r
	// start 》= end的时候 说明只有一个元素
	if start >= end {
		return
	}
    // 选定pivot
	pivot := nums[l]
	for l != r {
		// notice !! it must that the right go firstly
		// 保证右边的数都大于等于pivot
		for l < r && nums[r] >= pivot {
			r--
		}
		// 保证左边的数都小于等于pivot
		for l < r && nums[l] <= pivot {
			l++
		}
		// 因为前面已经有判断nums[h] 和pivot 所以在这的肯定是nums[l] > nums[h],两方交换
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	// 最后 因为我们选的pivot是nums[start], so 把pivot 和 nums[l] switch 以达到左边都比pivot小 右边都比pivot大
	nums[start] = nums[l]
	nums[l] = pivot
	// 然后以l为分界，分别进行quick sort
	quick_sort(nums, start, l-1)
	quick_sort(nums, l+1, end)
	return
}