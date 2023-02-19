package string

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabc"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	slen := len(s)
	rk, res := -1, 0

	for i := 0; i < slen; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		for rk+1 < slen && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		if rk-i+1 > res {
			res = rk - i + 1
		}
	}

	return res
}
