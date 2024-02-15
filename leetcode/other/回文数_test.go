package main

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	res := isPalindrome(132)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	slow, fast := 0, len(s)-1
	for slow < fast {
		if s[slow] != s[fast] {
			return false
		}
		slow++
		fast--
	}

	return true
}
