package day_one

import "testing"

// https://leetcode.cn/problems/binary-search/submissions/406371864/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	search(nums, target)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
