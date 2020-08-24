package dataStructures

type Stack struct {
	dq Deque
}

func (oldStack *Stack) ToString() string {
	return oldStack.dq.ToString()
}

func (oldStack *Stack) Length() string {
	return oldStack.dq.Length()
}

func (oldStack *Stack) Push(value string) {
	oldStack.dq.QueueHead(value)
}

func (oldStack *Stack) Peak() string {
	return oldStack.dq.PeakHead()
}

func (oldStack *Stack) Pop() string {
	return oldStack.dq.DequeueHead()
}

