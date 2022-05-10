type MinStack struct {
	s, m []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if len(this.m) == 0 || val < this.GetMin() {
		this.m = append(this.m, val)
	} else {
		this.m = append(this.m, this.GetMin())
	}
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.m = this.m[:len(this.m)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.m[len(this.m)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
