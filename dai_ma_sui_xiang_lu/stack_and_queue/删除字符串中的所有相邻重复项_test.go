package greedy_algorithm

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/

func TestRemoveDuplicates(t *testing.T) {
	str := "abbaca"
	result := removeDuplicates(str)
	fmt.Println(result)
}

func removeDuplicates(s string) string {
	var res []byte

	for i := 0; i < len(s); i++ {
		if len(res) > 0 && res[len(res)-1] == s[i] {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}
