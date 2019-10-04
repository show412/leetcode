//https://leetcode.com/problems/course-schedule-ii/
/*
There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites,
for example to take course 0 you have to first take course 1,
which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs,
return the ordering of courses you should take to finish all courses.

There may be multiple correct orders,
you just need to return one of them.
If it is impossible to finish all courses, return an empty array.

Example 1:

Input: 2, [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished
course 0. So the correct course order is [0,1] .
Example 2:

Input: 4, [[1,0],[2,0],[3,1],[3,2]]
Output: [0,1,2,3] or [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both
             courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
             So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3] .
Note:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
*/
// typical toplog sort
// 可以输出排序结果 也可以判断有没有环 有环的话 会有一些点没法实现入度为0 这个时候i就不可能等于num
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make(map[int][]int, 0)
	indgree := make([]int, numCourses)
	order := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		s := prerequisites[i][0]
		e := prerequisites[i][1]
		edges[e] = append(edges[e], s)
		indgree[s]++
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}
	i := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		order[i] = node
		i++
		for _, edge := range edges[node] {
			if indgree[edge] > 0 {
				indgree[edge]--
			}
			if indgree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}
	if i == numCourses {
		return order
	}
	return []int{}
}
