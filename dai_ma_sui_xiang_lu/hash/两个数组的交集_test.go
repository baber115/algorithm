package main

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/intersection-of-two-arrays/
func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1, 2}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	var res []int
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}

	return res
}
