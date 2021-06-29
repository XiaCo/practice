package dynamic

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	return _rob(root, m)
}

func _rob(root *TreeNode, m map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	// 偷了当前
	var res1, res2 int
	res1 = root.Val
	if root.Left != nil {
		val1 := value(m, root.Left.Left)
		val2 := value(m, root.Left.Right)
		res1 += val1 + val2
	}
	if root.Right != nil {
		val1 := value(m, root.Right.Left)
		val2 := value(m, root.Right.Right)
		res1 += val1 + val2
	}

	// 没偷当前
	res2 = value(m, root.Left) + value(m, root.Right)
	if res1 > res2 {
		return res1
	}
	return res2
}

func value(m map[*TreeNode]int, node *TreeNode) int {
	if node == nil {
		return 0
	}
	if val, exist := m[node]; exist {
		return val
	}
	val := _rob(node, m)
	m[node] = val
	return val
}
