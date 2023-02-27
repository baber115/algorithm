package day_two

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/peak-index-in-a-mountain-array/submissions/406946478/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6

func TestPeakIndexInMountainArray(t *testing.T) {
	nums := []int{3, 4, 5, 1}
	res := peakIndexInMountainArray(nums)
	fmt.Println(res)
}

// 枚举法
func peakIndexInMountainArrayI(arr []int) int {
	ans := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			ans = i
			break
		}
	}

	return ans
}

// 二分查找
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
