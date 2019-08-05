import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/bulls-and-cows/
/*
You are playing the following Bulls and Cows game 
with your friend: You write down a number 
and ask your friend to guess what the number is. 
Each time your friend makes a guess, 
you provide a hint that indicates 
how many digits in said guess match your secret number exactly 
in both digit and position (called "bulls") 
and how many digits match the secret number 
but locate in the wrong position (called "cows"). 
Your friend will use successive guesses 
and hints to eventually derive the secret number.

Write a function to return a hint 
according to the secret number and friend's guess, 
use A to indicate the bulls and B to indicate the cows.

Please note that both secret number and friend's guess 
may contain duplicate digits.

Example 1:

Input: secret = "1807", guess = "7810"

Output: "1A3B"

Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
Example 2:

Input: secret = "1123", guess = "0111"

Output: "1A1B"

Explanation: The 1st 1 in friend's guess is a bull, 
the 2nd or 3rd 1 is a cow.
Note: You may assume that the secret number 
and your friend's guess only contain digits, 
and their lengths are always equal.
*/
/*
test cases:
	"1807", "7810"  "1A3B"
	"1123", "0111" "1A1B"
  "1122" "1222" "3A0B"
*/
func getHint(secret string, guess string) string {
	//the map is for index to value
	mPtov := make(map[int]int)
	//the map is for value to position
	mVtop := make(map[int][]int)
	A := make([]int, 0)
	B := make([]int, 0)
	cow := make([]int, 0)
	for i := 0; i < len(secret); i++ {
		digit, _ := strconv.Atoi(string(secret[i]))
		mPtov[i] = digit
		mVtop[digit] = append(mVtop[digit], i)
	}
	fmt.Println(mPtov)
	fmt.Println(mVtop)
	for j := 0; j < len(guess); j++ {
		digit, _ := strconv.Atoi(string(guess[j]))
		if mPtov[j] == digit {
			A = append(A, digit)
			Postion(mVtop, digit, j)
			if len(mVtop[digit]) == 0 {
				delete(mVtop, digit)
			}
		} else {
			if _, ok := mVtop[digit]; ok {
				// select the candidates cow, but don't caculate B. Because we should find out A
				cow = append(cow, digit)
			}
		}
	}
	for k := 0; k < len(cow); k++ {
		digit := cow[k]
		if _, ok := mVtop[digit]; ok {
			B = append(B, digit)
			Postion(mVtop, digit, mVtop[digit][0])
			if len(mVtop[digit]) == 0 {
				delete(mVtop, digit)
			}
		}
	}
	res := strconv.Itoa(len(A)) + "A" + strconv.Itoa(len(B)) + "B"
	return res
}

func Postion(m map[int][]int, key int, index int) {
	for k := 0; k < len(m[key]); k++ {
		if m[key][k] == index {
			if k < len(m[key])-1 {
				m[key] = append(m[key][:k], m[key][k+1:]...)
			} else {
				m[key] = m[key][:k]
			}
			break
		}
	}
}

// 这两个方法都是一个意思
// here is a very efficent solution
// https://code.dennyzhang.com/bulls-and-cows

// Basic Ideas: 1 arrays with 10 elements; 1 pass
//
//  array[10]: positive: happen only in secret
//             negative: happen only in guess
//
//  For one specific position, we found ch1 in secret and ch2 in guess
//    If ch1==ch2, we increase bulls
//    Otherwise:
//       If array[ch1]<0, it means ch1 has happened in guess before. We increase cows
//       If array[ch2]>0, it means ch2 has happened in secret before. We increase cows
//       We increase array[ch1] by 1, decrease array[ch2] by 1
//
// Complexity: Time O(n), Space O(1)
func getHint(secret string, guess string) string {
	array := [10]int{}
	bulls, cows := 0, 0
	for i, _ := range secret {
		ch1, ch2 := secret[i]-'0', guess[i]-'0'
		if ch1 == ch2 {
			bulls += 1
			continue
		}
		if array[ch1] < 0 {
			cows += 1
		}
		if array[ch2] > 0 {
			cows += 1
		}
		array[ch1] += 1
		array[ch2] -= 1
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

// Blog link: https://code.dennyzhang.com/bulls-and-cows
// Basic Ideas: 2 arrays with 10 elements; 2 pass
// Complexity: Time O(n), Space O(1)
func getHint(secret string, guess string) string {
    secret_list := [10]int{}
    guess_list := [10]int{}
    bulls, cows := 0, 0
    for i, _ := range secret {
        if secret[i] == guess[i] {
            bulls += 1
        } else {
            secret_list[int(secret[i]-'0')] += 1
            guess_list[int(guess[i]-'0')] += 1
        }
    }
    for i:= 0; i<10; i++ {
        if secret_list[i] < guess_list[i] {
            cows += secret_list[i]
        } else {
            cows += guess_list[i]
        }
    }
    return fmt.Sprintf("%dA%dB", bulls, cows)