package double_pointer

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/reverse-words-in-a-string-iii/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=jgwt6s2

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	res := reverseWords(s)
	fmt.Println(res)
}

func reverseWords(s string) string {
	b := []byte(s)
	slow := 0
	for fast := 0; fast <= len(b); fast++ {
		if fast == len(b) || b[fast] == ' ' {
			reverseString(b, slow, fast-1)
			slow = fast + 1
		}
	}
	return string(b)
}

func reverseString(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}
