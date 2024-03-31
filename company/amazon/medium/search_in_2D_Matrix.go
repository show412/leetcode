/*
 * @Author: hongwei.sun
 * @Date: 2024-03-08 16:32:02
 * @LastEditors: your name
 * @LastEditTime: 2024-03-08 16:32:02
 * @Description: file content
 */
// https://leetcode.com/problems/search-a-2d-matrix/
/*
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
*/
// TC is O(log(mn)), SC is O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m*n - 1
	for left <= right {
		index := (left + right) / 2
		// 除代表在第几行，模(余)代表在第几列
		if matrix[index/n][index%n] == target {
			return true
		}
		// here is plus 1 because maybe the left is equal to right
		if matrix[index/n][index%n] < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || target < matrix[0][0] {
		return false
	}
	line_count := len(matrix)
	row_count := len(matrix[0])
	//   search which row, find the last position which is smaller than target in row[0]
	var start = 0
	var end = line_count - 1
	var mid = 0
	var row_position = 0
	for {
		if start+1 >= end {
			break
		}
		mid = start + (end-start)/2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			start = mid
		} else if matrix[mid][0] > target {
			end = mid
		}
	}

	// find the last position where target is bigger that the row[0]
	// so it should put the end at the last to rewrite the row_position
	if matrix[start][0] < target {
		row_position = start
	}

	if matrix[end][0] < target {
		row_position = end
	}

	if matrix[start][0] == target {
		return true
	}
	if matrix[end][0] == target {
		return true
	}
	// fmt.Println(start)
	// fmt.Println(end)
	// fmt.Println(row_position)

	//   find which line the target
	start = 0
	end = row_count - 1
	mid = 0
	for {
		if start+1 >= end {
			break
		}
		mid = start + (end-start)/2
		if matrix[row_position][mid] == target {
			return true
		} else if matrix[row_position][mid] < target {
			start = mid
		} else if matrix[row_position][mid] > target {
			end = mid
		}
	}

	if matrix[row_position][end] == target {
		return true
	}
	if matrix[row_position][start] == target {
		return true
	}
	return false

}
