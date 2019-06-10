// https://leetcode.com/problems/game-of-life/
func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	if len(board) == 1 && len(board[0]) == 1 {
		board[0][0] = 0
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			value := getValue(board, i-1, j-1) + getValue(board, i-1, j) +
				getValue(board, i-1, j+1) + getValue(board, i, j-1) +
				getValue(board, i, j+1) + getValue(board, i+1, j-1) +
				getValue(board, i+1, j) + getValue(board, i+1, j+1)
			if board[i][j]%2 == 1 && (value < 2 || value > 3) {
				board[i][j] = 3
			}
			if board[i][j]%2 == 0 && value == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 3 {
				board[i][j] = 0
			}
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}

}

func getValue(board [][]int, i int, j int) int {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return 0
	}
	return board[i][j] % 2
}
