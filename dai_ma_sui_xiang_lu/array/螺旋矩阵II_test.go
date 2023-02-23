package array

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/spiral-matrix-ii/
func TestSpiralMatrixII(t *testing.T) {
	result := generateMatrix(3)
	fmt.Println(result)
}

// 左闭右开原则
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	// 定义循环圈的开始位置
	startx, starty := 0, 0
	// 循环次数，如果n=3，则循环一次，n=5，循环2次
	loop := n / 2
	// 矩阵中间的位置，n=3，中间位置是(1,1)，n=5，中间位置是(2,2)
	mid := n / 2
	// 给矩阵每一个空格赋值
	count := 1
	// 控制每条边的遍历长度
	offset := 1

	for a := 0; a < n; a++ {
		matrix[a] = make([]int, n)
	}

	i, j := 0, 0
	for loop > 0 {
		i, j = startx, starty
		for j = starty; j < starty+n-offset; j++ {
			matrix[startx][j] = count
			count++
		}
		for i = startx; i < startx+n-offset; i++ {
			matrix[i][j] = count
			count++
		}

		for ; j > starty; j-- {
			matrix[i][j] = count
			count++
		}

		for ; i > startx; i-- {
			matrix[i][j] = count
			count++
		}
		// 第二圈开始的时候，起始位置各自加1
		// 第一圈是(0,0)，第二圈是(1,1)
		startx++
		starty++
		// offset用于控制每一圈中每一条边遍历的长度
		offset += 2
		loop--
	}

	if n%2 > 0 {
		matrix[mid][mid] = count
	}

	return matrix
}
