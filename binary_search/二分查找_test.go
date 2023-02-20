package binary_search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums, target := []int{9}, 9
	res := search(nums, target)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
