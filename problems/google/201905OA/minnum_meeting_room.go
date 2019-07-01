// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

// Example
// Example1

// Input: intervals = [(0,30),(5,10),(15,20)]
// Output: 2
// Explanation:
// We need two meeting rooms
// room1: (0,30)
// room2: (5,10),(15,20)
// Example2

// Input: intervals = [(2,7)]
// Output: 1
// Explanation:
// Only need one meeting room
// https://www.lintcode.com/problem/meeting-rooms-ii/description

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */

// 这是一个通用的算法，就是遇到任何interval，将interval的start当成key存入，value+1，遇到end，存入。但是要-1；
// 用一个TreeMap存储，保证key的值是sort的。
// 时间复杂度是O(N)
// 该方法可以用于如何calendar以及meeting room的题解。
// 关键是hash里面的key是要能排序的，golang因为不能保证hash的顺序 所以可能不一定能做这个题

func minMeetingRooms(intervals []*Interval) int {
	// Write your code here
	if len(intervals) == 0 {
		return 0
	}
	m := make(map[int]int, len(intervals))
	ants := 0
	cnts := 0
	for i := 0; i < len(intervals); i++ {
		if _, ok := m[intervals[i].Start]; ok {
			m[intervals[i].Start] = m[intervals[i].Start] + 1
		} else {
			m[intervals[i].Start] = 1
		}
		if _, ok := m[intervals[i].End]; ok {
			m[intervals[i].End] = m[intervals[i].End] - 1
		} else {
			m[intervals[i].End] = -1
		}
	}
	for _, v := range m {
		cnts = cnts + v
		ants = max(ants, cnts)
	}
	return ants
}
func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

// type intervalSort []*Interval

// func (a intervalSort) Len() int           { return len(a) }
// func (a intervalSort) Less(i, j int) bool { return a[i].Start < a[j].Start }
// func (a intervalSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func quickSort(intervalArray []*Interval, s int, e int) []*Interval {
	pivot := intervalArray[s].Start
	left := s
	right := e
	for s != e {
		for s < e && intervalArray[e].Start >= pivot {
			e--
		}
		for s < e && intervalArray[s].Start <= pivot {
			s++
		}
		if s < e {
			intervalArray[s], intervalArray[e] = intervalArray[e], intervalArray[s]
		}
	}
	temp := intervalArray[left]
	intervalArray[left] = intervalArray[s]
	intervalArray[s] = temp
	quickSort(intervalArray, left, s-1)
	quickSort(intervalArray, s+1, right)
	return intervalArray
}

func minMeetingRooms(intervals []*Interval) int {
	// Write your code here
	if len(intervals) == 0 {
		return 0
	}
	// perhaps, we need to sort by the si firstly
	// sort.Sort(intervals)
	intervals = quickSort(intervals, 0, len(intervals)-1)
	// state, f[i] means the first i need the minnum meeting rooms
	f := make([]int, len(intervals))
	f[0] = 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= intervals[i-1].Start && intervals[i].Start <= intervals[i-1].End {
			f[i] = f[i-1] + 1
		} else {
			f[i] = f[i-1]
		}
	}
	return f[len(intervals)-1]
}

// public class Solution {
//     /**
//      * @param intervals: an array of meeting time intervals
//      * @return: the minimum number of conference rooms required
//      */
//     public int minMeetingRooms(List<Interval> intervals) {
//         // Write your code here
//         TreeMap<Integer, Integer> tmap = new TreeMap<>();
//         int ans = 0, cnt = 0;
//         for (Interval i : intervals) {
//             int start = i.start, end = i.end;
//             tmap.put(start, tmap.getOrDefault(start, 0) + 1);
//             tmap.put(end, tmap.getOrDefault(end, 0) - 1);
//         }

//         for (int k : tmap.keySet()) {
//             System.out.println(k);
//             System.out.println(tmap.get(k));
//             cnt += tmap.get(k);
//             ans = Math.max(ans, cnt);
//         }
//         return ans;
//     }
// }

