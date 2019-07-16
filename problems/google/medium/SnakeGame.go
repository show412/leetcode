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

type SnakeGame struct {
	snake [][]int
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
		snake: [][]int{[]int{0,0}}
		width:  width-1,
		height: height-1,
		food:   food
}

/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
  Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
	move := map[string][]int{'U' : []int{-1, 0}, 'L' : []int{0, -1}, 'R' : []int{0, 1}, 'D' : []int{1, 0}}
	head := len(this.snake) - 1
	if len(this.food) > 0 && eatFood(this.snake[head], move[direction], this.food[0]) == true {
		this.snake = append(this.snake, this.food[0])
		this.food = this.food[1:]
	} else {

	}

}

func eatFood(shankHead []int, moveD []int, food []int) bool {

}

/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */
