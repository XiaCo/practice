package arrays

import (
	"github.com/XiaCo/practice/algorithms/listNode"
)

// https://leetcode-cn.com/problems/sliding-window-maximum/

func getMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// 暴力
func MaxSlidingWindow(nums []int, k int) []int {
	step := len(nums) - k
	var result []int
	for i := 0; i <= step; i++ {
		result = append(result, getMax(nums[i:i+k]))
	}
	return result
}

// 单调栈
func MaxSlidingWindow2(nums []int, k int) []int {
	var result []int
	var dq = listNode.Constructor(len(nums))
	for index, num := range nums {
		// 构建单调递减栈
		for !dq.IsEmpty() && nums[dq.GetRear()] <= num {
			dq.DeleteLast()
		}
		dq.InsertLast(index)

		// 校验最大值是否是最windows边界外的值
		if dq.GetFront() <= index-k {
			dq.DeleteFront()
		}

		// 将最大值，放入结果队列
		if index+1 >= k {
			result = append(result, nums[dq.GetFront()])
		}
	}
	return result
}
