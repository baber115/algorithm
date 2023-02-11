package main

import (
	"testing"
)

// https://leetcode.cn/problems/4sum-ii/

func TestFourSumCount(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fourSumCount(nums1, nums2, nums3, nums4)
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++ // 注意这里可能会有sum值相等的情况
		}
	}

	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[0-v3-v4]
		}
	}

	return count
}
