package dataStructures

type Deque struct {
	ll    LinkedList
	Head  **LinkedListNode
	Tail  **LinkedListNode
	Count *int
}

func (odq *Deque) Init(value string) {
	odq.ll.Init(value)
	odq.Head = &odq.ll.Head
	odq.Tail = &odq.ll.Tail
	odq.Count = &odq.ll.Count
}

func (odq *Deque) ToString() string {
	return odq.ll.ToString()
}

func (odq *Deque) Length() string {
	return odq.ll.Length()
}

func (odq *Deque) QueueHead(value string) {

	if odq.Count == nil || *odq.Count == 0 {
		odq.Init(value)
		return
	}
	odq.ll.AddHead(value)
}

func (odq *Deque) QueueTail(value string) {
	if odq.Count == nil || *odq.Count == 0 {
		odq.Init(value)
		return
	}
	odq.ll.AddTail(value)
}

func (odq *Deque) PeakHead() string {
	return (*odq.Head).value
}

func (odq *Deque) PeakTail() string {
	return (*odq.Tail).value
}

func (odq *Deque) DequeueHead() string {
	if odq.Count == nil || *odq.Count == 0 {
		return ""
	} else {
		result := (*odq.Head).value
		odq.ll.RemoveHead()
		return result
	}
}

func (odq *Deque) DequeueTail() string {
	if odq.Count == nil || *odq.Count == 0 {
		return ""
	} else {
		result := (*odq.Tail).value
		odq.ll.RemoveTail()
		return result
	}
}
