package stack

// https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	elements []int
	index    int // 指向当前有元素的栈顶
	minIndex int // 存储最小元素的下标
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nil, -1, -1}
}

func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if this.minIndex == -1 {
		this.minIndex = 0
	}
	this.index++
	if this.elements[this.index] < this.elements[this.minIndex] {
		this.minIndex = this.index
	}
}

func (this *MinStack) Pop() {
	if this.index == -1 {
		return
	}

	this.elements = this.elements[:this.index]
	if this.minIndex == this.index { // 重新计算最小值的下标
		if this.minIndex == 0 {
			this.minIndex = -1
		} else {
			var newMinIndex int
			for index, num := range this.elements {
				if num < this.elements[newMinIndex] {
					newMinIndex = index
				}
			}
			this.minIndex = newMinIndex
		}
	}
	this.index--
}

func (this *MinStack) Top() int {
	if this.index == -1 {
		return 0
	}
	return this.elements[this.index]
}

func (this *MinStack) GetMin() int {
	if this.minIndex != -1 {
		return this.elements[this.minIndex]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
