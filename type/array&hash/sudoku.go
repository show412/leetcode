/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-03 16:52:39
 * @Description: file content
 */
// https://leetcode.com/problems/valid-sudoku/
/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.
*/
/*
this question is about finding if there is duplicated 1-9 in row, column and 3*3 sub box
row and clomun we can use hashmap, the difficult is about how to represent subbox
可以把每一个 sub box 看成一个3*3二维数组，key 就是 [i/3, j/3]
*/

func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[int]int, 0)
	colMap := make(map[int]map[int]int, 0)
	subMap := make(map[[2]int]map[int]int, 0)
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[int]int, 0)
		colMap[i] = make(map[int]int, 0)
		
	}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            subMap[[2]int{i, j}] = make(map[int]int, 0)
        }
    }
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				rowMap[i][num]++
				colMap[j][num]++
				subMap[[2]int{i / 3, j / 3}][num]++
				if rowMap[i][num] > 1 || colMap[j][num] > 1 || subMap[[2]int{i / 3, j / 3}][num] > 1 {
					return false
				}
			}
		}
	}
	return true
}