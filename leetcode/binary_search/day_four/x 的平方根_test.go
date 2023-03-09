package day_four

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/sqrtx/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6
func TestMySqrt(t *testing.T) {
	res := mySqrt(1)
	fmt.Println(res)
}

func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		tmp := mid * mid
		if tmp > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}

	return ans
}
