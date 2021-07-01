package dynamic

// https://leetcode-cn.com/problems/dungeon-game/submissions/

// 6  4 1
// 5 10 4
// 0 0 -5
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	// matrix 存 当前格最低血量 = 下一格所需要的最低血量 - 当前格的增益效果（伤害效果）
	// 如果<=0 ，则为0 （暂时忽略血量为0时，会死亡的情况）
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	// 初始化终点、下边界、右边界
	if val := dungeon[m-1][n-1]; val >= 0 {
		matrix[m-1][n-1] = 0
	} else {
		matrix[m-1][n-1] = -val
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i != m-1 && j != n-1 {
				val := minNum(matrix[i+1][j], matrix[i][j+1]) - dungeon[i][j]
				if val <= 0 {
					val = 0
				}
				matrix[i][j] = val
			} else if i != m-1 { // 边界
				val := matrix[i+1][n-1] - dungeon[i][n-1]
				if val <= 0 {
					val = 0
				}
				matrix[i][n-1] = val
			} else if j != n-1 {
				val := matrix[m-1][j+1] - dungeon[m-1][j]
				if val <= 0 {
					val = 0
				}
				matrix[m-1][j] = val
			}
		}
	}

	return matrix[0][0] + 1
}
