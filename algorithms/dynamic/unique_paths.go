package dynamic

// f(x, y) 到达x，y时，有多少种路线
// f(x, y) => f(x, y-1) + f(x-1, y)
// 爬楼梯的思想
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

// 构建矩阵
// f[x][y] == f[x][y-1] + f[x-1][y]
func uniquePaths1(m int, n int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		matrix[i][0] = 1
	}
	for i := 0; i < n; i++ {
		matrix[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}
	return matrix[m-1][n-1]
}
