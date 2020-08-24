package main

import (
	datastructures "DataStructures/pkg/datastructures"
	"fmt"
)

func main() {
	testStack()
}

func testStack()  {
	st := new(datastructures.Stack)
	fmt.Println("Initial Stack")
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Push obj1")
	st.Push("obj1")
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Push obj2")
	st.Push("obj2")
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Push obj3")
	st.Push("obj3")
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("PEAK")
	fmt.Println("Peaking: " + st.Peak())
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Popping")
	fmt.Println("Popped: " + st.Pop())
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Popping")
	fmt.Println("Popped: " + st.Pop())
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Popping")
	fmt.Println("Popped: " + st.Pop())
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Popping")
	fmt.Println("Popped: " + st.Pop())
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())

	fmt.Println("Popping")
	fmt.Println("Popped: " + st.Pop())
	fmt.Println(st.ToString())
	fmt.Println("Count: " + st.Length())
}

func testQueue()  {
	qu := new(datastructures.Queue)
	fmt.Println("Initial Queue")
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Queue obj1")
	qu.Queue("obj1")
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Queue obj2")
	qu.Queue("obj2")
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Queue obj3")
	qu.Queue("obj3")
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("PEAK")
	fmt.Println("Peaking: " + qu.Peak())
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Dequeue")
	fmt.Println("Dequeue: " + qu.Dequeue())
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Dequeue")
	fmt.Println("Dequeue: " + qu.Dequeue())
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Dequeue")
	fmt.Println("Dequeue: " + qu.Dequeue())
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Dequeue")
	fmt.Println("Dequeue: " + qu.Dequeue())
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())

	fmt.Println("Dequeue")
	fmt.Println("Dequeue: " + qu.Dequeue())
	fmt.Println(qu.ToString())
	fmt.Println("Count: " + qu.Length())
}

func testDeque() {
	dq := new(datastructures.Deque)
	fmt.Println("Initial Deque")
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Add obj1 on HEAD")
	dq.QueueHead("obj1")
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Add obj2 on HEAD")
	dq.QueueHead("obj2")
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Add obj3 on TAIL")
	dq.QueueTail("obj3")
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Add obj4 on TAIL")
	dq.QueueTail("obj4")
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("PEAK Head")
	fmt.Println("Peaking: " + dq.PeakHead())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("PEAK Tail")
	fmt.Println("Peaking: " + dq.PeakTail())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Dequeue Head")
	fmt.Println("Dequeued: " + dq.DequeueHead())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Dequeue Tail")
	fmt.Println("Dequeued: " + dq.DequeueTail())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Dequeue Head")
	fmt.Println("Dequeued: " + dq.DequeueHead())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Dequeue Head")
	fmt.Println("Dequeued: " + dq.DequeueHead())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Dequeue Head")
	fmt.Println("Dequeued: " + dq.DequeueHead())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())

	fmt.Println("Dequeue Head")
	fmt.Println("Dequeued: " + dq.DequeueHead())
	fmt.Println(dq.ToString())
	fmt.Println("Count: " + dq.Length())
}

func testSortedList() {
	sl := new(datastructures.SortedList)
	fmt.Println("Initial List")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Add("obj2")
	fmt.Println("After add obj2")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Add("obj1")
	fmt.Println("After add obj1")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Add("obj4")
	fmt.Println("After add obj4")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Add("obj3")
	fmt.Println("After add obj3")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Add("obj5")
	fmt.Println("After add obj5")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	fmt.Println("Doest it have obj6?")
	fmt.Println(sl.HasValue("obj6"))
	fmt.Println("Count: " + sl.Length())

	fmt.Println("Doest it have obj5?")
	fmt.Println(sl.HasValue("obj5"))
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj6")
	fmt.Println("After remove obj6")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj1")
	fmt.Println("After remove obj1")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj3")
	fmt.Println("After remove obj3")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj5")
	fmt.Println("After remove obj5")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj2")
	fmt.Println("After remove obj2")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj3")
	fmt.Println("After remove obj3")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())

	sl.Remove("obj4")
	fmt.Println("After remove obj4")
	fmt.Println(sl.ToString())
	fmt.Println("Count: " + sl.Length())
}

func testLinkedList() {
	ll := new(datastructures.LinkedList)
	fmt.Println("Initial List")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.AddHead("obj1")
	fmt.Println("After addHead")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.AddHead("obj2")
	fmt.Println("After addHead2")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.AddTail("obj3")
	fmt.Println("After addTail3")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.AddHead("obj4")
	fmt.Println("After addHead4")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	fmt.Println("Doest it have obj5?")
	fmt.Println(ll.HasValue("obj5"))
	fmt.Println("Count: " + ll.Length())

	fmt.Println("Doest it have obj4?")
	fmt.Println(ll.HasValue("obj4"))
	fmt.Println("Count: " + ll.Length())

	ll.Remove("obj5")
	fmt.Println("After remove obj5")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.Remove("obj1")
	fmt.Println("After remove obj1")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.Remove("obj3")
	fmt.Println("After remove obj3")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.Remove("obj4")
	fmt.Println("After remove obj4")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.Remove("obj2")
	fmt.Println("After remove obj2")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())

	ll.Remove("obj2")
	fmt.Println("After remove obj2")
	fmt.Println(ll.ToString())
	fmt.Println("Count: " + ll.Length())
}
