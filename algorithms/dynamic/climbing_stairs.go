package dynamic

// https://leetcode.com/problems/climbing-stairs/
// 什么是动态规划：https://www.zhihu.com/question/23995189
// 核心：动态规划 F(n) = F(n-1) + F(n-2)

// 递归
func ClimbStairs(n int) int {
	if n <= 3 {
		return n
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

// 滚动数组
func ClimbStairs2(n int) int {
	var n1, n2, result int
	n1 = 0
	n2 = 1
	result = n1 + n2
	for i := 1; i < n; i++ {
		n1 = n2
		n2 = result
		result = n1 + n2
	}
	return result
}
