package day_three

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/description/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6&languageTags=golang

func TestFindTheDistanceValue(t *testing.T) {
	arr1 := []int{-803, 715, -224, 909, 121, -296, 872, 807, 715, 407, 94, -8, 572, 90, -520, -867, 485, -918, -827, -728, -653, -659, 865, 102, -564, -452, 554, -320, 229, 36, 722, -478, -247, -307, -304, -767, -404, -519, 776, 933, 236, 596, 954, 464}
	arr2 := []int{817, 1, -723, 187, 128, 577, -787, -344, -920, -168, -851, -222, 773, 614, -699, 696, -744, -302, -766, 259, 203, 601, 896, -226, -844, 168, 126, -542, 159, -833, 950, -454, -253, 824, -395, 155, 94, 894, -766, -63, 836, -433, -780, 611, -907, 695, -395, -975, 256, 373, -971, -813, -154, -765, 691, 812, 617, -919, -616, -510, 608, 201, -138, -669, -764, -77, -658, 394, -506, -675, 523, 730, -790, -109, 865, 975, -226, 651, 987, 111, 862, 675, -398, 126, -482, 457, -24, -356, -795, -575, 335, -350, -919, -945, -979, 611}
	d := 37
	ans := findTheDistanceValue(arr1, arr2, d)
	fmt.Println(ans)
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	ans := 0
	for _, v1 := range arr1 {
		low, high := v1-d, v1+d
		res := binarySearch(arr2, low, high)
		if res {
			ans++
		}
	}

	return ans
}

/*
在arr2里找y，且满足num-d <= y <= num+d，如果存在则arr1中的x就不满足题目要求
*/
func binarySearch(arr2 []int, low, high int) bool {
	left, right := 0, len(arr2)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr2[mid] < low {
			left = mid + 1
		} else if arr2[mid] > high {
			right = mid - 1
		} else {
			// arr2[mid] >= low && arr2[mid] <= high
			return false
		}
	}

	return true
}
