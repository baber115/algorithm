package string

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	reverseWords(s)
}

func reverseWords(s string) string {
	// 1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	// 删除头部冗余空格
	for len(b) > 1 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}

	// 删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}

	// 删除尾部空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}

	// 2.反转字符串
	reverse(b, 0, len(b)-1)

	// 3.反转字符串中的单词
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {

		}
		reverse(b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[right], s[left] = s[left], s[right]
		left++
		right--
	}
}
