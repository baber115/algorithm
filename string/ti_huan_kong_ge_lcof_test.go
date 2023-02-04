package string

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	str := "We are happy."
	a := replaceSpace(str)
	fmt.Println(a)
}

func replaceSpace(s string) string {
	str := []byte(s)
	strLength := len(str)
	spaceCount := 0
	for _, v := range str {
		if v == ' ' {
			spaceCount++
		}
	}
	tmp := make([]byte, spaceCount*2)
	str = append(str, tmp...)

	i := strLength - 1
	j := len(str) - 1

	for i >= 0 {
		if str[i] == ' ' {
			str[j] = '0'
			str[j-1] = '2'
			str[j-2] = '%'
			j = j - 3
		} else {
			str[j] = str[i]
			j--
		}
		i--
	}

	return string(str)
}
