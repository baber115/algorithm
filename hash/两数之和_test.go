package main

import "testing"

// https://leetcode.cn/problems/two-sum/

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{
				k,
				index,
			}
		} else {
			m[v] = k
		}
	}

	return []int{}
}
