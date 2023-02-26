package day_three

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{10, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast <= len(nums)-1 {
		if nums[fast] != 0 {
			temp := nums[slow]
			nums[slow] = nums[fast]
			nums[fast] = temp
			slow++
		}
		fast++
	}
}
