import "sort"

// https://leetcode.com/problems/queue-reconstruction-by-height/
/*
Suppose you have a random list of people standing in a queue.
Each person is described by a pair of integers (h, k),
where h is the height of the person
and k is the number of people in front of this person
who have a height greater than or equal to h.
Write an algorithm to reconstruct the queue.

Note:
The number of people is less than 1,100.


Example

Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/
/*
a = [h,k] if h is the shortest, k !=0,
the a should be on k index position of the left free index
if a and b have same h, bigger k should be first
because the if smaller k is first to book index
in next loop the index will go to next index, it's wrong
思路和这个很像 只是这道题是反着思考的
https://leetcode.com/problems/queue-reconstruction-by-height/solution/
根本思想是 “矮子插队无所谓，反正高个子看不到”
所以先处理最矮的k 肯定是正好的位置 剩下的就处理剩下的 把当前已经占了的位置不算在里面
*/

func reconstructQueue(people [][]int) [][]int {
	if len(people) == 0 {
		return people
	}
	res := make([][]int, len(people))
	// sort.Sort should use the defined type
	sort.Sort(peopleSlice(people))

	for i := 0; i < len(people); i++ {
		// start=0 就是把当前里面的不算在实际计算的index里
		start := 0
		for j := 0; j < len(res); j++ {
			if res[j] != nil {
				continue
			}
			if start == people[i][1] {
				res[j] = people[i]
			}
			start++
		}
	}
	return res
}

type peopleSlice [][]int

func (a peopleSlice) Len() int {
	return len(a)
}
func (a peopleSlice) Less(i, j int) bool {
	return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] > a[j][1])
}
func (a peopleSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
