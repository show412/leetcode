/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-01-25 16:07:51
 * @Description: file content
 */
import "strings"

// Given two words word1 and word2,
// find the minimum number of steps required to convert word1 to word2.
// (each operation is counted as 1 step.)

// You have the following 3 operations permitted on a word:

// Insert a character
// Delete a character
// Replace a character
// Example
// Example 1:

// Input:
// "horse"
// "ros"
// Output: 3
// Explanation:
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
// Example 2:

// Input:
// "intention"
// "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')
// https://www.lintcode.com/problem/edit-distance/description

/**
 * @param word1: A string
 * @param word2: A string
 * @return: The minimum number of steps.
 */
func minDistance(word1 string, word2 string) int {
	// write your code here
	// DP to implement it
	// state
	word1Array := strings.Split(word1, "")
	word2Array := strings.Split(word2, "")
	if len(word1Array) == 0 {
		return len(word2Array)
	}
	if len(word2Array) == 0 {
		return len(word1Array)
	}
	// f[i][j] means how many times the first i in word1 match to the first j in word2 by adjusting at the least
	f := make([][]int, len(word1Array)+1)
	for i := 0; i < len(word1Array)+1; i++ {
		f[i] = make([]int, len(word2Array)+1)
	}
	// initialze
	f[0][0] = 0
	for i := 0; i < len(word1Array)+1; i++ {
		f[i][0] = i
	}
	for j := 0; j < len(word2Array)+1; j++ {
		f[0][j] = j
	}
	// function
	for i := 1; i < len(word1Array)+1; i++ {
		for j := 1; j < len(word2Array)+1; j++ {
			if word1Array[i-1] == word2Array[j-1] {
				f[i][j] = min(f[i-1][j-1], f[i-1][j]+1, f[i][j-1]+1)
			} else {
				f[i][j] = min(f[i-1][j-1], f[i-1][j], f[i][j-1]) + 1
			}
		}
	}
	return f[len(word1Array)][len(word2Array)]
}
func min(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}
