package day_two

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/search-insert-position/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6

func TestSearchInsert(t *testing.T) {
	nums := []int{
		1, 3, 5, 6,
	}
	target := 7
	res := searchInsert(nums, target)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
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
