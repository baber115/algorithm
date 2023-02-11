package main

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.cn/problems/4sum/

/*
此题和三数之和类似
*/
func TestFourSum(t *testing.T) {
	nums := []int{
		2, 2, 2, 2, 2,
	}
	target := 8
	result := fourSum(nums, target)
	fmt.Println(result)
}

/*
四指针
i j left right
*/
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		// 剪枝处理
		if n1 > target && n1 >= 0 {
			break
		}
		/*
			不能这样写，可能是负数
			nums =[-4,-3,-2,-1],target=-10
			if n1 > target{
				break
			}
		*/
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			// 剪枝处理
			if (n1+n2) > target && (n1+n2) >= 0 {
				break
			}
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				n3, n4 := nums[left], nums[right]
				sum := n1 + n2 + n3 + n4
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					result = append(result, []int{
						n1, n2, n3, n4,
					})
					for left < right && nums[left+1] == n3 { // 去重
						left++
					}
					for left < right && nums[right-1] == n4 { // 去重
						right--
					}
					// 找到答案时，双指针同时靠近
					right--
					left++
				}
			}
		}
	}

	return result
}
