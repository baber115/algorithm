package string

import (
	"testing"
)

func TestReverseStringII(t *testing.T) {
	reverseStr("abcdefg", 2)
}

func reverseStr(s string, k int) string {
	str := []byte(s)
	length := len(str)
	for i := 0; i < length; i += (2 * k) {
		if i+k <= length {
			reverseString(str[i : i+k])
		} else {
			reverseString(str[i:length])
		}
	}

	return string(str)
}

func TestReverseString(t *testing.T) {
	str := []byte("Hello")
	reverseString(str)
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[right], s[left] = s[left], s[right]
		left++
		right--
	}
}
