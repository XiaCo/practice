package arrays

// https://leetcode-cn.com/problems/container-with-most-water/
// 双指针
// 一个从头向尾出发，一个从尾向头出发
// 移动哪一个指针根据两个指针所在位置的墙高度来决定
// 移动矮的面积可能变大或变小，移动高的面积肯定变小

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxArea(height []int) int {
	var indexHead, indexTail int
	indexHead = 0
	indexTail = len(height) - 1
	var max int
	for indexHead < indexTail {
		minNum := minInt(height[indexHead], height[indexTail])
		if area := (indexTail - indexHead) * minNum; area > max {
			max = area
		}
		if minNum == height[indexHead] {
			indexHead++
		} else {
			indexTail--
		}
	}
	return max
}
