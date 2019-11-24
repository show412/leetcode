import (
	"fmt"
	"math/bits"
)

// https://leetcode.	com/problems/binary-watch/
/*
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.


For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".
*/
func readBinaryWatch(num int) []string {
	res := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			// onesCount代表这个数最少用几个二级制位组成
			if bits.OnesCount(uint(i))+bits.OnesCount(uint(j)) == num {
				// 注意这里要用Sprintf 因为printf返回的是int err.
				// format要用%02 这样当j不是2位的时候左边用0补齐 这是printf在各种语言中的通用格式
				res = append(res, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return res
}

