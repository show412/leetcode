/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-11 23:33:37
 * @Description: file content
 */
import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
/*
Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:

Division between two integers should truncate toward zero.
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
Example 1:

Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: ["4", "13", "5", "/", "+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
Output: 22
Explanation:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
/*
数组从前往外pop数到一个栈里， 当遇到符号的时候就从这个栈里pop出两个数计算
然后把结果再放到这个栈里
*/

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		operand := tokens[i]
		if operand != "+" && operand != "-" && operand != "*" && operand != "/" {
			stack = append(stack, operand)
		} else {
			operand2, _ := strconv.Atoi(stack[len(stack)-1])
			operand1, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			result := 0
			if operand == "+" {
				result = operand1 + operand2
			}
			if operand == "-" {
				result = operand1 - operand2
			}
			if operand == "*" {
				result = operand1 * operand2
			}
			if operand == "/" {
				result = operand1 / operand2
			}
			stack = append(stack, strconv.Itoa(result))
		}
	}
	res, _ := strconv.Atoi(stack[len(stack)-1])
	return res
}

/*
below is very clear and beaitiful and advanced code using struct
*/
type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	x := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return x
}

func (s *Stack) Top() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) getOperands() (int, int) {
	num1 := s.Pop()
	num2 := s.Pop()
	return num1, num2
}

func evalRPN(tokens []string) int {
	stack := NewStack()
	for _, val := range tokens {
		switch val {
		case "+":
			num1, num2 := stack.getOperands()
			stack.Push(num2 + num1)
		case "-":
			num1, num2 := stack.getOperands()
			stack.Push(num2 - num1)
		case "*":
			num1, num2 := stack.getOperands()
			stack.Push(num2 * num1)
		case "/":
			num1, num2 := stack.getOperands()
			stack.Push(num2 / num1)
		default:
			num, err := strconv.Atoi(val)
			if err != nil {
				panic("unsupported value")
			}
			stack.Push(num)
		}
	}
	return stack.Top()
}