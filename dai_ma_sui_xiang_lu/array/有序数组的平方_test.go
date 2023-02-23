package array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/squares-of-a-sorted-array/
func TestSquaresOfASortedArray(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	a := sortedSquares(nums)
	fmt.Println(a)
}

func sortedSquares(nums []int) []int {
	count := len(nums)
	i, j := 0, count-1
	result := make([]int, count)

	for i < j {
		left, right := nums[i]*nums[i], nums[j]*nums[j]
		count--
		if left < right {
			result[count] = right
			j--
		} else {
			result[count] = left
			i++
		}
	}

	return result
}
