package sliding_window

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=jgwt6s2

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "1231231"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	fast := -1
	m := make(map[byte]bool, len(s))
	for slow := 0; slow < len(s); slow++ {
		if slow > 0 {
			delete(m, s[slow-1])
		}
		for fast+1 < len(s) && m[s[fast+1]] == false {
			m[s[fast+1]] = true
			fast++
		}
		fmt.Println(string(s[slow]))
		sum := fast - slow + 1
		if sum > res {
			res = sum
		}
	}

	return res
}
