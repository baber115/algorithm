package greedy_algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	res := evalRPN(tokens)
	fmt.Println(res)
}

func evalRPN(tokens []string) int {
	var res []int
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err == nil {
			res = append(res, val)
		} else {
			num1, num2 := res[len(res)-2], res[len(res)-1]
			res = res[:len(res)-2]
			switch v {
			case "+":
				res = append(res, num1+num2)
				break
			case "-":
				res = append(res, num1-num2)
				break
			case "*":
				res = append(res, num1*num2)
				break
			case "/":
				res = append(res, num1/num2)
				break
			}
		}
	}

	return res[0]
}
