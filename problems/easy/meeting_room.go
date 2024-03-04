import "sort"

/*
 * @Author: hongwei.sun
 * @Date: 2024-03-03 17:19:22
 * @LastEditors: your name
 * @LastEditTime: 2024-03-03 17:24:52
 * @Description: file content
 */
// https://leetcode.com/problems/meeting-rooms/description/
// Given an array of meeting time intervals consisting of start and end times[[s1,e1],[s2,e2],...](si< ei), determine if a person could attend all meetings.
// Example 1:
// Input:
// [[0,30],[5,10],[15,20]]
// Output:
//  false
// Example 2:
// Input:
//  [[7,10],[2,4]]

// Output:
//  true
/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

/**
 * @param intervals: an array of meeting time intervals
 * @return: if a person could attend all meetings
 */
/*
1, sort intervals by start time
2, if meeting end time is later than next meeting start time, means can't attend next meeting
*/
func CanAttendMeetings(intervals []*Interval) bool {
	// Write your code here
	if len(intervals) == 0 || len(intervals) == 1 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		before := intervals[i]
		end := intervals[j]
		if before.Start < end.Start { //按照开始时间进行排序
			return true
		}
		return false
	})

	for i := 0; i < len(intervals)-1; i++ {
		j := i + 1
		if intervals[i].End > intervals[j].Start {
			return false
		}
	}
	return true
}