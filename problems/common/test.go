package main

import (
	"fmt"
)

func main() {
	// reg := regexp.MustCompile("([a-z]*)([0-9]+)(\\[\1)(.)*(\\]\1)")
	// str := "aaaa3[a2[cc]]2[bc]"
	// data := reg.FindAllStringSubmatch(str, -1)
	// fmt.Println(data)
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(res)
}

// Input:  [0,1,2,4,5,7]
// Output: ["0->2","4->5","7"]
func findDisappearedNumbers(nums []int) []int {
	nlen := len(nums)
	for i := 0; i < nlen; {
		if nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			continue
		}
		i++
	}
	var res []int
	for i := 0; i < nlen; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
