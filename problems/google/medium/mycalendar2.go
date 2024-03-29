// https://leetcode.com/problems/my-calendar-ii/

/*
Implement a MyCalendarTwo class to store your events. A new event can be added if adding the event will not cause a triple booking.

Your class will have one method, book(int start, int end). Formally, this represents a booking on the half open interval [start, end), the range of real numbers x such that start <= x < end.

A triple booking happens when three events have some non-empty intersection (ie., there is some time that is common to all 3 events.)

For each call to the method MyCalendar.book, return true if the event can be added to the calendar successfully without causing a triple booking. Otherwise, return false and do not add the event to the calendar.

Your class will be called like this: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
Example 1:

MyCalendar();
MyCalendar.book(10, 20); // returns true
MyCalendar.book(50, 60); // returns true
MyCalendar.book(10, 40); // returns true
MyCalendar.book(5, 15); // returns false
MyCalendar.book(5, 10); // returns true
MyCalendar.book(25, 55); // returns true
Explanation:
The first two events can be booked.  The third event can be double booked.
The fourth event (5, 15) can't be booked, because it would result in a triple booking.
The fifth event (5, 10) can be booked, as it does not use time 10 which is already double booked.
The sixth event (25, 55) can be booked, as the time in [25, 40) will be double booked with the third event;
the time [40, 50) will be single booked, and the time [50, 55) will be double booked with the second event.


Note:

The number of calls to MyCalendar.book per test case will be at most 1000.
In calls to MyCalendar.book(start, end), start and end are integers in the range [0, 10^9].
*/

type MyCalendarTwo struct {
	// it should be defined into [][]int, because we need to know one range
	// if it's defined into []int, it's only one point by one point
	calendarOne [][2]int
	calendarTwo [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{calendarOne: make([][2]int, 0), calendarTwo: make([][2]int, 0)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, v := range this.calendarTwo {
		// it should have the overlap not the cover
		// 这是这道题的关键 如何判断有overlap
		if end > v[0] && start < v[1] {
			return false
		}
	}
	for _, v := range this.calendarOne {
		// it should have the overlap not the cover
		if end > v[0] && start < v[1] {
			this.calendarTwo = append(this.calendarTwo, [2]int{max(start, v[0]), min(end, v[1])})
		}
	}
	// 这也要记录 因为start end可能会为后续的overlap做判断使用
	this.calendarOne = append(this.calendarOne, [2]int{start, end})
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
