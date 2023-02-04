package string

import (
	"fmt"
	"testing"
)

func TestFF(t *testing.T) {
	s, k := "abcdefg", 2
	b := reverseLeftWords(s, k)
	fmt.Println(b)
}

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, len(b)-1)
	reverse(b, 0, len(b)-n-1)
	reverse(b, len(b)-n, len(b)-1)

	return string(b)
}
