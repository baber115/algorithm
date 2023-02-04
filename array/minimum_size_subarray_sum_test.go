package array

import (
	"fmt"
	"testing"
)

func TestMinimumSizeSubarraySum(T *testing.T) {
	s := 15
	nums := []int{1, 2, 3, 4, 5}
	a := minSubArrayLen(s, nums)
	fmt.Println(a)
}

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	i := 0
	length := len(nums)
	result := length + 1
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if subLen < result {
				result = subLen
			}
			sum -= nums[i]
			i++
		}
	}
	if result == length+1 {
		return 0
	}
	return result
}
