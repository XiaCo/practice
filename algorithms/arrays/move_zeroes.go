package arrays

// https://leetcode-cn.com/problems/move-zeroes/
// 双指针
// 每回合快指针都将数字赋值给慢指针所在位置
// 快指针每步都向后移动一格，慢指针遇见非0数字才移动

func MoveZeroes(nums []int) {
	var indexSlow, indexFast int
	for indexFast < len(nums) {
		nums[indexSlow] = nums[indexFast]
		indexFast++
		if nums[indexSlow] != 0 {
			indexSlow++
		}
	}
	for indexSlow < len(nums) {
		nums[indexSlow] = 0
		indexSlow++
	}
}
