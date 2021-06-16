package dynamic

// https://leetcode-cn.com/problems/minimum-path-sum/

// f[m][n] => min{f[m-1][n], f[m][n-1]} + grid[m][n]
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 == m && j+1 == n {
				continue
			}
			if i+1 == m {
				grid[i][j] += grid[i][j+1]
				continue
			}
			if j+1 == n {
				grid[i][j] += grid[i+1][j]
				continue
			}
			grid[i][j] = minNum(grid[i+1][j], grid[i][j+1]) + grid[i][j]
		}
	}
	return grid[0][0]
}

func minNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
