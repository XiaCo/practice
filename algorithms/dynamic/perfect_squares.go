package dynamic

import "math"

// https://leetcode-cn.com/problems/perfect-squares

// 官方方法
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 我的方法。
// 思路其实都是 numSquares(n) => min{numSquares(n-square)} + 1
// 官方写法表明了某些情况，可以使用数组缓存结果，而不必使用map
var m = make(map[int]int)

func numSquares1(n int) int {
	if n == 1 {
		return 1
	}
	if value, exist := m[n]; exist {
		return value
	}
	var min = n
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			m[n] = 1
			return 1
		}
		num := numSquares(n - i*i)
		if num < min {
			min = num
		}
	}

	m[n] = min + 1
	return min + 1
}
