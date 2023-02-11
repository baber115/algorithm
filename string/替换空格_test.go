package string

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/

func TestReplaceSpace(t *testing.T) {
	str := "We are happy."
	a := replaceSpace(str)
	fmt.Println(a)
}

/*
能提前算出最终结果数组的长度
从后往前替换，不需要考虑替换时把后面的元素后移
*/
func replaceSpace(s string) string {
	b := []byte(s)
	left := len(b) - 1
	spaceCount := 0
	for _, v := range s {
		if v == ' ' {
			spaceCount++
		}
	}
	// 原来数组里面有一个空格，替换成%20，所以*2，不是*3
	resizeCount := spaceCount * 2
	temp := make([]byte, resizeCount)
	b = append(b, temp...)
	right := len(b) - 1
	for left < right {
		if s[left] == ' ' {
			b[right] = '0'
			b[right-1] = '2'
			b[right-2] = '%'
			right -= 3
		} else {
			b[right] = b[left]
			right--
		}
		left--
	}

	return string(b)
}
