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
		// 除了head 后面的点肯定都和上一个点至少column或者row是相等的
		if move(this.snake, path, direction, this.width, this.height) == -1 {
			return -1
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

func move(snake [][]int, path map[string][]int, d string, width int, height int) int {
	r := snake[len(snake)-1][0]
	c := snake[len(snake)-1][1]
	snake[len(snake)-1][0] += path[d][0]
	snake[len(snake)-1][1] += path[d][1]
	if snake[len(snake)-1][0] > height || snake[len(snake)-1][1] > width || snake[len(snake)-1][0] < 0 || snake[len(snake)-1][1] < 0 {
		// fmt.Println(snake[len(snake)-1])
		return -1
	}
	if len(snake) > 1 {
		for i := len(snake) - 2; i >= 0; i-- {
			// same direction with next node
			c1 := snake[i][0]
			r1 := snake[i][1]
			if c == c1 && r == r1 {
				// 	fmt.Println("move before")
				// fmt.Println(snake[len(snake)-1])
				// 	fmt.Println(snake[i])
				return -1
			}

			// {"U": []int{-1, 0}, "L": []int{0, -1}, "R": []int{0, 1}, "D": []int{1, 0}}
			// in one row
			if snake[i][0] == c {
				if snake[i][1] < r {
					snake[i][0] += path["D"][0]
					snake[i][1] += path["D"][1]
				} else {
					snake[i][0] += path["U"][0]
					snake[i][1] += path["U"][1]
				}
			} else {
				// in one column
				if snake[i][0] < c {
					snake[i][0] += path["R"][0]
					snake[i][1] += path["R"][1]
				} else {
					snake[i][0] += path["L"][0]
					snake[i][1] += path["L"][1]
				}
			}
			c = c1
			r = r1
			// bites self
			if snake[i][0] == snake[len(snake)-1][0] && snake[i][1] == snake[len(snake)-1][1] {
				// fmt.Println("move after")
				// fmt.Println(c1)
				// fmt.Println(r1)
				// fmt.Println(c)
				// fmt.Println(r)
				// fmt.Println(snake[len(snake)-1])
				// fmt.Println(snake[i])
				return -1
			}

		}
	}
	return 0
}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
