package dynamic

// https://leetcode-cn.com/problems/unique-paths-ii/submissions/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	matrix := make([][]int, m)
	var flag = 1
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			flag = 0
		}
		matrix[i] = make([]int, n)
		matrix[i][0] = flag
	}
	flag = 1

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			flag = 0
		}
		matrix[0][i] = flag
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			matrix[i][j] = matrix[i][j-1] + matrix[i-1][j]
		}
	}

	return matrix[m-1][n-1]
}
