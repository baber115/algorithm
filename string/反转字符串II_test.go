package string

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/reverse-string-ii/
func TestReverseStr(t *testing.T) {
	s, k := "abcdefg", 2
	a := reverseStr(s, k)
	fmt.Println(a)
}

func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(ss); i += 2 * k {
		if i+k > len(ss) {
			ReverseString(ss[i : i+k])
		} else {
			ReverseString(ss[i:])
		}
	}

	return string(ss)
}
