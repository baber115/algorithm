package day_three

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/description/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6&languageTags=golang

func TestFindTheDistanceValue(t *testing.T) {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	findTheDistanceValue(arr1, arr2, d)
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	ans := 0
	for i := 0; i < len(arr1); i++ {
		left, right := 0, len(arr2)
		for left <= right {
			mid := left + (right-left)/2
			square := int(math.Abs(float64(arr1[i] - arr2[mid])))
			if square {
				ans--
				break
			}
		}
		ans++
	}

	return ans
}
