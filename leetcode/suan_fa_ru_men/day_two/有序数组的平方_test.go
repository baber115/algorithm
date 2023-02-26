package day_two

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/squares-of-a-sorted-array/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=jgwt6s2

func TestSortedSquares(t *testing.T) {
	n := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(n)
	fmt.Println(res)
}

func sortedSquares(nums []int) []int {
	count := len(nums)
	left, right := 0, count-1
	res := make([]int, count)
	for left <= right {
		ln, rn := nums[left]*nums[left], nums[right]*nums[right]
		count--
		if ln < rn {
			res[count] = rn
			right--
		} else {
			res[count] = ln
			left++
		}
	}

	return res
}
