package day_six

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	s1, s2 := "adc", "dcda"
	res := checkInclusion(s1, s2)
	fmt.Println(res)
}

func checkInclusion(s1 string, s2 string) bool {
	// 计算出s1的数组
	var m1 [26]int
	fast := 0
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a']++
		fast++
	}
	for i := 0; i < len(s2) && fast <= len(s2); i++ {
		var m2 [26]int
		fmt.Println(i, fast)
		// 滑动窗口判断m1和m2是否相等
		for j := i; j < fast; j++ {
			m2[s2[j]-'a']++
		}
		fmt.Println(m1, m2)
		if m1 == m2 {
			return true
		}
		fast++
	}

	return false
}
