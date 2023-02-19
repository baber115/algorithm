package string

import (
	"fmt"
	"testing"
)

// TODO
// https://leetcode.cn/problems/longest-palindromic-substring/

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {

	return ""
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
