package day_one

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/search-insert-position

func TestSearchInsert(t *testing.T) {
	nums, target := []int{1, 3, 5, 6}, 7
	res := searchInsertI(nums, target)
	fmt.Println(res)
}

func searchInsertI(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchInsertII(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
