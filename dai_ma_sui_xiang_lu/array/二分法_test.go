package array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/binary-search/
func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	result := search(nums, target)
	fmt.Println(result)
}

func search(nums []int, target int) int {
	slow := 0
	fast := len(nums) - 1
	for slow <= fast {
		mid := (fast-slow)/2 + slow
		if target < nums[mid] {
			fast = mid - 1
		} else if target > nums[mid] {
			slow = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
