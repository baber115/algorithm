package main

import (
	"testing"
)

// https://leetcode.cn/problems/valid-anagram/
func TestIsAnagram(test *testing.T) {
	s, t := "nagaram", "anagram"
	isAnagram(s, t)
}

func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, v := range s {
		record[v-'a']++
	}
	for _, v := range t {
		record[v-'a']--
	}
	return record == [26]int{}
}
