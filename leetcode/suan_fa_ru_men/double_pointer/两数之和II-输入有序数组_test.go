package double_pointer

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 3, 4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sub := numbers[left] + numbers[right]
		if sub > target {
			right--
		} else if sub < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}
