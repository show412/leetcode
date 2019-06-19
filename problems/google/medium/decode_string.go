import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/decode-string/
/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string],
where the encoded_string inside the square brackets is being repeated exactly k times.
Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid;
No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
For example, there won't be input like 3a or 2[4].

Examples:

s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
*/
/*
Try to use a regex to solve the problem is difficult, becuase there is a confilct
between [ and into '[' with [. It's diffcult to write the regex
func decodeString(s string) string {
	reg := regexp.MustCompile("([0-9]+)(\\[)(.*)(\\])")
	data := reg.FindStringSubmatch(str)
	if len(data) == 0 {
		return data
	}
	k := data[1]
	subStr := data[3]
	res := ""
	for i := 0; i < k; i++ {
		res += decodeString(subStr)
	}
	return res
}
*/

func is_digit(c byte) bool {
	return c >= '0' && c <= '9'
}

type State struct {
	input []byte
	index int
}

func (s *State) curr() byte {
	return s.input[s.index]
}

func (s *State) curr_k() int {
	var num []byte
	num = append(num, s.curr())
	for s.next() {
		if !is_digit(s.curr()) {
			s.undo()
			break
		}
		num = append(num, s.curr())
	}

	i, err := strconv.Atoi(string(num))

	if err != nil {
		return 0
	} else {
		return i
	}
}

func (s *State) undo() {
	s.index--
}

func (s *State) next() bool {
	s.index++
	if s.index >= len(s.input) {
		return false
	} else {
		return true
	}
}

func decode_l(s *State) []byte {
	b := []byte{}
	b = append(b, s.curr())

	for {
		if !s.next() {
			break
		}

		if is_digit(s.curr()) {
			s.undo()
			break
		}

		b = append(b, s.curr())
	}
	return b
}

func decode_k(s *State) []byte {
	k := s.curr_k()
	fmt.Println(k)
	b := []byte{}
	s.next() // '['
	for {
		if !s.next() {
			break
		}
		c := s.curr()
		if c == ']' {
			// End
			o := b
			b = []byte{}
			for i := 0; i < k; i++ {
				b = append(b, o...)
			}
			return b
		} else if is_digit(c) {
			// Sub
			sb := decode_k(s)
			for i := 0; i < len(sb); i++ {
				b = append(b, sb[i])
			}
		} else {
			b = append(b, c)
		}
	}
	return b
}

func decodeString(s string) string {
	state := &State{
		input: []byte(s),
		index: -1,
	}

	var r []byte

	for state.next() {
		if is_digit(state.curr()) {
			// k string
			r = append(r, decode_k(state)...)
		} else {
			// literal string
			r = append(r, decode_l(state)...)
		}
	}

	return string(r)
}
