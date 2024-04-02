/*
 * @Author: hongwei.sun
 * @Date: 2024-04-02 14:19:00
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-02 14:39:02
 * @Description: file content
 */
//  https://leetcode.com/problems/surrounded-regions/description/
/*
Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

 

Example 1:


Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation: Notice that an 'O' should not be flipped if:
- It is on the border, or
- It is adjacent to an 'O' that should not be flipped.
The bottom 'O' is on the border, so it is not flipped.
The other three 'O' form a surrounded region, so they are flipped.
Example 2:

Input: board = [["X"]]
Output: [["X"]]
 

Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is 'X' or 'O'.
*/
/*
遍历第一行和最后一行，遍历第一列和最后一列，标记与'O'相连的区域为“#”
然后再全体遍历，把'O'变为'X'， ‘#’ 变为‘O’
*/
func solve(board [][]byte) {
    if len(board) == 0 {
        return
    }

    rows, cols := len(board), len(board[0])

    // 定义DFS函数
    var dfs func(row, col int)
    dfs = func(row, col int) {
        if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != 'O' {
            return
        }
        board[row][col] = '#'
        dfs(row+1, col)
        dfs(row-1, col)
        dfs(row, col+1)
        dfs(row, col-1)
    }

    // 遍历第一行和最后一行，标记与'O'相连的区域
    for i := 0; i < rows; i++ {
        dfs(i, 0)
        dfs(i, cols-1)
    }

    // 遍历第一列和最后一列，标记与'O'相连的区域
    for j := 0; j < cols; j++ {
        dfs(0, j)
        dfs(rows-1, j)
    }

    // 将标记为'#'的区域恢复为'O'，其余的'O'标记为'X'
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == '#' {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }
}