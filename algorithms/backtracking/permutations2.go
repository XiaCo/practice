package backtracking

// https://leetcode-cn.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var find func(now []int, remain []int)
	find = func(now []int, remain []int) {
		if len(remain) == 0 {
			res = append(res, now)
		}
		for index := range remain {
			if index != 0 && remain[index] == remain[index-1] {
				continue
			}
			re := append(append([]int{}, remain[:index]...), remain[index+1:]...)
			find(append(now, remain[index]), re)
		}
	}
	find(nil, nums)
	return res
}
