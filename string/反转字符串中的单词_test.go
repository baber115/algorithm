package string

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	a := reverseWords(s)
	fmt.Println(a)
}

func reverseWords(s string) string {
	b := []byte(s)
	ReverseString(b)

	slow := 0
	for fast := 0; fast <= len(b); fast++ {
		if fast == len(b) || b[fast] == ' ' {
			ReverseString(b[slow:fast])
			slow = fast + 1
		}
	}

	return string(b)
}
