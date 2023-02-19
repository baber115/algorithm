package greedy_algorithm

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/valid-parentheses/

func TestIsValid(t *testing.T) {
	str := "()[]{}"
	result := isValid(str)
	fmt.Println(result)
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	template := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var result []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			result = append(result, s[i])
		} else if len(result) > 0 && template[s[i]] == result[len(result)-1] {
			result = result[:len(result)-1]
		} else {
			return false
		}
	}

	return len(result) == 0
}
