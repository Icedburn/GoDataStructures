package dataStructures

type Queue struct {
	dq Deque
}

func (oldQueue *Queue) ToString() string {
	return oldQueue.dq.ToString()
}

func (oldQueue *Queue) Length() string {
	return oldQueue.dq.Length()
}

func (oldQueue *Queue) Queue(value string) {
	oldQueue.dq.QueueTail(value)
}

func (oldQueue *Queue) Peak() string {
	return oldQueue.dq.PeakHead()
}

func (oldQueue *Queue) Dequeue() string {
	return oldQueue.dq.DequeueHead()
}

