package day_three

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/valid-perfect-square/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=f2ex1x6

func TestIsPerfectSquare(t *testing.T) {
	num := 16
	res := isPerfectSquare(num)
	fmt.Println(res)
	num1 := 14
	res1 := isPerfectSquare(num1)
	fmt.Println(res1)
	num2 := 1
	res2 := isPerfectSquare(num2)
	fmt.Println(res2)
}

// 暴力法
func isPerfectSquareI(num int) bool {
	left := 1
	for left <= num {
		square := left * left
		if square == num {
			return true
		}
		if square > num {
			return false
		}
		left++
	}

	return false
}

// 二分法
func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square > num {
			right = mid - 1
		} else if square < num {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
