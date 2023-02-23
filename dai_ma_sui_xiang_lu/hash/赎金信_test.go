package main

import (
	"testing"
)

// https://leetcode.cn/problems/ransom-note/solutions/
func TestCanConstruct(t *testing.T) {
	ransomNote := "aa"
	magazine := "aab"
	canConstruct(ransomNote, magazine)
}

func canConstruct(ransomNote string, magazine string) bool {
	// 只有小写字母，可以把数组缩小到26
	m := make([]int, 26)
	// 先循环后一个，获取所有字符
	for _, v := range magazine {
		m[v-'a']++
	}
	// 循环目标字符串，一旦有复数肯定不能组成目标字符串
	for _, v := range ransomNote {
		m[v-'a']--
		if m[v-'a'] < 0 {
			return false
		}
	}
	return true
}
