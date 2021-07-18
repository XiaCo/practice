package backtracking

// https://leetcode-cn.com/problems/permutations/

func permute(nums []int) [][]int {
	var res [][]int
	var find func(now []int, remain []int)
	find = func(now []int, remain []int) {
		if len(remain) == 0 {
			res = append(res, append([]int{}, now...))
			return
		}
		for index := range remain {
			var re []int
			re = append(append([]int{}, remain[:index]...), remain[index+1:]...)
			find(append(now, remain[index]), re)
		}
	}
	find(nil, nums)
	return res
}
