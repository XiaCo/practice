package dynamic

// https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
// 此题提交到leetcode出问题，但是执行代码按钮是跑通的

var ropeMap = make(map[int]int)

func cuttingRope(n int) int {
	return cutting(n, true)
}

// cuttingRope(n) => max{cuttingRope(n-length)*length}
// length > 0
// n - length > 0
func cutting(n int, needCut bool) int {
	if value, exist := ropeMap[n]; exist {
		return value
	}
	var max int
	for length := 1; length < n; length++ {
		if num := cutting(n-length, false) * length; num > max {
			max = num
		}
	}
	if needCut {
		return max
	}
	if n > max {
		max = n
	}
	ropeMap[n] = max
	return max
}
