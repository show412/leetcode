// https://leetcode.com/problems/design-snake-game/
/*
Design a Snake game that is played on a device with screen size = width x height. Play the game online if you are not familiar with the game.

The snake is initially positioned at the top left corner (0,0) with length = 1 unit.

You are given a list of food's positions in row-column order. When a snake eats the food, its length and the game's score both increase by 1.

Each food appears one by one on the screen. For example, the second food will not appear until the first food was eaten by the snake.

When a food does appear on the screen, it is guaranteed that it will not appear on a block occupied by the snake.

Example:

Given width = 3, height = 2, and food = [[1,2],[0,1]].

Snake snake = new Snake(width, height, food);

Initially the snake appears at position (0,0) and the food at (1,2).

|S| | |
| | |F|

snake.move("R"); -> Returns 0

| |S| |
| | |F|

snake.move("D"); -> Returns 0

| | | |
| |S|F|

snake.move("R"); -> Returns 1 (Snake eats the first food and right after that, the second food appears at (0,1) )

| |F| |
| |S|S|

snake.move("U"); -> Returns 1

| |F|S|
| | |S|

snake.move("L"); -> Returns 2 (Snake eats the second food)

| |S|S|
| | |S|

snake.move("U"); -> Returns -1 (Game over because snake collides with border)
*/
/*
test cases:
["SnakeGame","move","move","move","move","move","move"]
[[3,2,[[1,2],[0,1]]],["R"],["D"],["R"],["U"],["L"],["U"]]

["SnakeGame","move","move","move","move"]
[[2,2,[[1,1],[0,1]]],["D"],["R"],["U"],["L"]]

["SnakeGame","move","move","move","move","move","move","move","move","move","move","move","move","move","move","move","move","move","move"]
[[3,3,[[2,0],[0,0],[0,2],[0,1],[2,2],[2,1]]],["D"],["D"],["R"],["U"],["U"],["L"],["D"],["R"],["R"],["U"],["L"],["L"],["D"],["R"],["R"],["D"],["L"],["L"]]

["SnakeGame","move","move","move","move","move","move","move","move","move","move","move","move","move","move","move"]
[[3,3,[[2,0],[0,0],[0,2],[0,1],[2,2],[0,1]]],["D"],["D"],["R"],["U"],["U"],["L"],["D"],["R"],["R"],["U"],["L"],["L"],["D"],["R"],["U"]]


refer to https://leetcode.com/problems/design-snake-game/discuss/294312/Java-straightforward-solution-beats-99
*/

type SnakeGame struct {
	snake  [][]int
	width  int
	height int
	food   [][]int
}

/** Initialize your data structure here.
  @param width - screen width
  @param height - screen height
  @param food - A list of food positions
  E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */

func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		snake:  [][]int{[]int{0, 0}},
		width:  width - 1,
		height: height - 1,
		food:   food}
}

/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
	Game over when snake crosses the screen boundary or bites its body.
	食物一定是head吃的 因为说When a food does appear on the screen, it is guaranteed that it will not appear on a block occupied by the snake.
*/

func (this *SnakeGame) Move(direction string) int {
	path := map[string][]int{"U": []int{-1, 0}, "L": []int{0, -1}, "R": []int{0, 1}, "D": []int{1, 0}}
	head := len(this.snake) - 1
	if len(this.food) > 0 && eatFood(this.snake[head], path[direction], this.food[0]) == true {
		this.snake = append(this.snake, this.food[0])
		this.food = this.food[1:]
	} else {
		// 除了head 后面的点都是往前提一位
		r := this.snake[len(this.snake)-1][0]
		c := this.snake[len(this.snake)-1][1]
		r += path[direction][0]
		c += path[direction][1]
		if r > this.height || c > this.width || r < 0 || c < 0 {
			return -1
		}
		this.snake = this.snake[1:]
		this.snake = append(this.snake, []int{r, c})
		// check hit self
		for i := len(this.snake) - 2; i >= 0; i-- {
			if this.snake[i][0] == r && this.snake[i][1] == c {
				return -1
			}
		}
	}
	return len(this.snake) - 1
}

func eatFood(shankHead []int, d []int, food []int) bool {
	if (shankHead[0]+d[0]) == food[0] && (shankHead[1]+d[1]) == food[1] {
		return true
	}
	return false
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
