package string

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/length-of-last-word/

func TestLengthOfLastWord(t *testing.T) {
	s := "   fly me   to   the moon  "
	//s := "a"
	res := lengthOfLastWord(s)
	fmt.Println(res)
}

func lengthOfLastWord(s string) int {
	r := len(s) - 1
	for s[r] == ' ' {
		r--
	}
	res := 0
	for r >= 0 && s[r] != ' ' {
		r--
		res++
	}
	return res
}
