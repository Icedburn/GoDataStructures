package dataStructures

import (
	"strconv"
	"strings"
)

type LinkedListNode struct {
	value string
	next *LinkedListNode
	prev *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
	Count int
}

func (oll *LinkedList) Init(value string) {
	n := LinkedListNode{
		value: value,
		next:  nil,
		prev:  nil,
	}

	oll.Head = &n
	oll.Tail = &n
	oll.Count = 1
}

func (oll *LinkedList) ToString() string {
	result := "["
	nextNode := oll.Head
	for nextNode != nil {
		result += nextNode.value + ", "
		nextNode = nextNode.next
	}
	return strings.TrimSuffix(result, ", ") + "]"
}

func (oll *LinkedList) Length() string {
	return strconv.Itoa(oll.Count)
}

func (oll *LinkedList) AddHead(value string) {
	if oll.Head == nil {
		oll.Init(value)
	} else {
		newNode := LinkedListNode{
			value: value,
			next:  oll.Head,
			prev:  nil,
		}

		oll.Head.prev = &newNode
		oll.Head = &newNode
		oll.Count++
	}
}

func (oll *LinkedList) AddTail(value string) {
	if oll.Head == nil {
		oll.Init(value)
	} else {
		newNode := LinkedListNode{
			value: value,
			next:  nil,
			prev:  oll.Tail,
		}

		oll.Tail.next = &newNode
		oll.Tail = &newNode
		oll.Count++
	}
}

func (oll *LinkedList) FindNode(value string) *LinkedListNode {
	nextNode := oll.Head
	for nextNode != nil {
		if nextNode.value == value {
			break
		}
		nextNode = nextNode.next
	}
	return nextNode
}

func (oll *LinkedList) HasValue(value string) bool {
	return oll.FindNode(value) != nil
}

func (oll *LinkedList) RemoveHead() bool {
	if oll.Count == 0 {
		return false
	} else if oll.Count == 1 {
		oll.Count = 0
		oll.Head = nil
		oll.Tail = nil
	} else {
		oll.Head.next.prev = nil
		oll.Head = oll.Head.next
		oll.Count--
	}
	return true
}

func (oll *LinkedList) RemoveTail() bool {
	if oll.Count == 0 {
		return false
	} else if oll.Count == 1 {
		oll.RemoveHead()
	} else {
		oll.Tail.prev.next = nil
		oll.Tail = oll.Tail.prev
		oll.Count--
	}
	return true
}

func (oll *LinkedList) Remove(value string) bool {
	nodeToRemove := oll.FindNode(value)
	if nodeToRemove == nil {
		return false
	} else {
		if oll.Count == 1 {
			//Only LinkedListNode
			oll.Head = nil
			oll.Tail = nil
			oll.Count = 0
		} else if nodeToRemove.prev == nil {
			//Is Head
			oll.Head = nodeToRemove.next
			oll.Head.prev = nil
			oll.Count--
		} else if nodeToRemove.next == nil {
			//Is Tail
			oll.Tail = nodeToRemove.prev
			oll.Tail.next = nil
			oll.Count--
		} else {
				//Is in the middle of 2 other nodes
				nodeToRemove.prev.next = nodeToRemove.next
				nodeToRemove.next.prev = nodeToRemove.prev
				oll.Count--
		}
		return true
	}
}