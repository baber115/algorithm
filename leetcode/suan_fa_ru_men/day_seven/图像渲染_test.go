package day_seven

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/flood-fill/?envType=study-plan&id=suan-fa-ru-men&plan=algorithms&plan_progress=jgwt6s2

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr, sc := 1, 1
	newColor := 2
	res := floodFill(image, sr, sc, newColor)
	fmt.Println(res)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// 保存初始点颜色，判断是否要改变颜色
	currColor := image[sr][sc]
	if currColor != color {
		dfs(image, sr, sc, currColor, color)
	}
	return image
}

// 深度优先搜索
func dfs(image [][]int, x, y, curColor, color int) {
	// 超过边界，上下左右的边界
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) {
		return
	}
	// 色块与初始的色块不同
	if curColor != image[x][y] {
		return
	}
	image[x][y] = color
	// 基于当前块，遍历上下左右
	dfs(image, x-1, y, curColor, color)
	dfs(image, x+1, y, curColor, color)
	dfs(image, x, y-1, curColor, color)
	dfs(image, x, y+1, curColor, color)
}
