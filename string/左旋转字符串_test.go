package string

import (
	"testing"
)

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

func TestReverseLeftWords(t *testing.T) {
	s, k := "abcdefg", 2
	reverseLeftWords(s, k)
}

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	ReverseString(b[:n])
	ReverseString(b[n:])
	ReverseString(b)
	return string(b)
}
