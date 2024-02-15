package main

import (
	"fmt"
	"testing"
)

/*
给定一个罗马数字，将其转换成整数。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/roman-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func TestRomanToInt(t *testing.T) {
	res := romanToInt("XLIX")
	fmt.Println(res)
}

func romanToInt(s string) int {
	temp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	slen := len(s)
	for i := range s {
		num := temp[s[i]]
		if i < slen-1 && num < temp[s[i+1]] {
			res -= num
		} else {

			res += num
		}
	}
	return res
}
