package main

import "testing"

// https://leetcode.cn/problems/happy-number/solutions/
/*
题中会会有死循环，所以sum会有重复，可以通过判断sum是否重复来结束循环
*/
func TestIsHappy(t *testing.T) {

}

// 用哈希集合检测循环
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}

	return n == 1
}

// 快慢指针法
func isHappyII(n int) bool {
	slow, fast := n, getSum(n)
	for fast != 1 && fast != slow {
		slow = getSum(n)
		fast = getSum(getSum(n))
	}

	return fast == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}
