package leetcode

type MinStack struct {
	elements  []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{
		elements:  []int{},
		minValues: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)

	if len(this.minValues) == 0 {
		this.minValues = append(this.minValues, val)
	} else {
		currentMin := this.minValues[len(this.minValues)-1]
		if val < currentMin {
			this.minValues = append(this.minValues, val)
		} else {
			this.minValues = append(this.minValues, currentMin)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.elements) > 0 {
		this.elements = this.elements[:len(this.elements)-1]
		this.minValues = this.minValues[:len(this.minValues)-1]
	}
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValues[len(this.minValues)-1]
}
