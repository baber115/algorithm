package double_pointer

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	j := 0
	for i, v := range nums {
		j = (i + k) % len(nums)
		fmt.Println(j)
		newNums[j] = v
	}
	copy(nums, newNums)
}
