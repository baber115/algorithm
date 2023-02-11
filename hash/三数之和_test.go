package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	//nums := []int{
	//	-1, 0, 1, 2, -1, -4,
	//}
	nums := []int{
		-1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1,
	}
	a := threeSum(nums)
	fmt.Println(a)
}

/*
三指针方法
*/
func threeSum(nums []int) [][]int {
	// 1. 先排序
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		// 如果第一个数字大于0，3个数肯定没法加成0
		if n1 > 0 {
			break
		}
		// 如果不加i>0来判断，会误判，-1，-1，2
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		// 2.定义快慢指针
		// 如果和大于0,right--，如果小于0，left++，否则就是答案
		left := i + 1
		right := len(nums) - 1
		for left < right {
			n2, n3 := nums[left], nums[right]
			sum := n1 + n2 + n3
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{
					n1, n2, n3,
				})
				// 题目要求不能重复，所以left和right需要去重
				/*
					假设，[-1,0,0,0,0,0,0,0,1,1,1,1,1,1]
					应该只返回[[-1,0,1],[0,0,0]]
				*/
				for left < right && nums[left+1] == n2 {
					left++
				}
				for left < right && nums[right-1] == n3 {
					right--
				}
				right--
				left++
			}
		}
	}

	return result
}
