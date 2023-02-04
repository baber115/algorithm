package array

import (
	"fmt"
	"testing"
)

func TestFF(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	a := removeElement(nums, val)
	fmt.Println(a, nums)
}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
