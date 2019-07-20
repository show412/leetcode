package main

import (
	"fmt"
	"strconv"
	// "math"
)

/*
test case:
"kkkkzrkatkwpkkkktrq"
"bbbbaaaaababaababab"
*/
func main() {
	res := getHint("1123", "0111")
	fmt.Println(res)
}

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
				// B = append(B, digit)
				// Postion(mVtop, digit, mVtop[digit][0])
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
