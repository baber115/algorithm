package string

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/

/*
KMP算法
最大前缀表
*/
func TestStrStr(t *testing.T) {
	//haystack, needle := "sadbutsad", "sad"
	haystack, needle := "aabaabaafa", "aabaaf"
	strStr(haystack, needle)
}

/*
KMP算法
*/
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	if m == 0 {
		return 0
	}
	// 求最大前缀表
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	fmt.Println(pi)

	// 判断相等字符串
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == m {
			return i - m + 1
		}
	}
	return -1
}

/*
暴力解法
我们可以让字符串 needle 与字符串 haystack 的所有长度为 m 的子串均匹配一次。
*/
func strStrI(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}

	return -1
}
