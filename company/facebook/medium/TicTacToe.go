// https://leetcode.com/problems/design-tic-tac-toe/
/*
Design a Tic-tac-toe game that is played between two players on a n x n grid.

You may assume the following rules:

A move is guaranteed to be valid and is placed on an empty block.
Once a winning condition is reached, no more moves is allowed.
A player who succeeds in placing n of their marks
in a horizontal, vertical, or diagonal row wins the game.
Example:
Given n = 3, assume that player 1 is "X" and player 2 is "O" in the board.

TicTacToe toe = new TicTacToe(3);

toe.move(0, 0, 1); -> Returns 0 (no one wins)
|X| | |
| | | |    // Player 1 makes a move at (0, 0).
| | | |

toe.move(0, 2, 2); -> Returns 0 (no one wins)
|X| |O|
| | | |    // Player 2 makes a move at (0, 2).
| | | |

toe.move(2, 2, 1); -> Returns 0 (no one wins)
|X| |O|
| | | |    // Player 1 makes a move at (2, 2).
| | |X|

toe.move(1, 1, 2); -> Returns 0 (no one wins)
|X| |O|
| |O| |    // Player 2 makes a move at (1, 1).
| | |X|

toe.move(2, 0, 1); -> Returns 0 (no one wins)
|X| |O|
| |O| |    // Player 1 makes a move at (2, 0).
|X| |X|

toe.move(1, 0, 2); -> Returns 0 (no one wins)
|X| |O|
|O|O| |    // Player 2 makes a move at (1, 0).
|X| |X|

toe.move(2, 1, 1); -> Returns 1 (player 1 wins)
|X| |O|
|O|O| |    // Player 1 makes a move at (2, 1).
|X|X|X|
Follow up:
Could you do better than O(n^2) per move() operation?
*/

type TicTacToe struct {
	rows         []int
	cols         []int
	diagonal     int
	antiDiagonal int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows:         make([]int, n),
		cols:         make([]int, n),
		diagonal:     0,
		antiDiagonal: 0}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	play := player
	// -1 stands for player 2
	if player == 2 {
		play = -1
	}
	size := len(this.rows)
	this.rows[row] += play
	this.cols[col] += play
	// in diagonal
	if row == col {
		this.diagonal += play
	}
	// in anti diagonal
	if col == len(this.cols)-1-row {
		this.antiDiagonal += play
	}
	if abs(this.rows[row]) == size ||
		abs(this.cols[col]) == size ||
		abs(this.diagonal) == size ||
		abs(this.antiDiagonal) == size {
		return player
	}
	return 0
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
