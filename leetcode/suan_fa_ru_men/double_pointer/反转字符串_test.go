package double_pointer

import (
	"fmt"
	"testing"
)

func TestReverseStringII(t *testing.T) {
	s := []byte{
		'h', 'e', 'l', 'l', 'o',
	}
	fmt.Println(s)
	ReverseString(s)
	fmt.Println(s)
}

func ReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
