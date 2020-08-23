package dataStructures

import (
	"strconv"
	"strings"
)

type SortedListNode struct {
	value string
	next *SortedListNode
	prev *SortedListNode
}

type SortedList struct {
	Head *SortedListNode
	Tail *SortedListNode
	Count int
}

func (osl *SortedList) Init(value string) {
	n := SortedListNode{
		value: value,
		next:  nil,
		prev:  nil,
	}

	osl.Head = &n
	osl.Tail = &n
	osl.Count = 1
}

func (osl *SortedList) ToString() string {
	result := "["
	nextNode := osl.Head
	for nextNode != nil {
		result += nextNode.value + ", "
		nextNode = nextNode.next
	}
	return strings.TrimSuffix(result, ", ") + "]"
}

func (osl *SortedList) Length() string {
	return strconv.Itoa(osl.Count)
}

func (osl *SortedList) Add(value string) {
	newNode := SortedListNode{
		value: value,
	}

	if osl.Count == 0 {
		osl.Head = &newNode
		osl.Tail = &newNode
		osl.Count++
		return
	}

	nextNode := osl.Head
	for nextNode != nil {
		// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
		if strings.Compare(nextNode.value, value) > 0 {
			break
		}
		nextNode = nextNode.next
	}

	if nextNode == nil {
		//Add in the Tail
		newNode.prev = osl.Tail
		osl.Tail.next = &newNode
		osl.Tail = &newNode
	} else if nextNode.prev == nil {
		//Add in the Head
		newNode.next = osl.Head
		osl.Head.prev = &newNode
		osl.Head = &newNode
	} else {
		newNode.next = nextNode
		newNode.prev = nextNode.prev
		nextNode.prev.next = &newNode
		nextNode.prev = &newNode
	}
	osl.Count++
}

func (osl *SortedList) FindNode(value string) *SortedListNode {
	found := false
	nextNode := osl.Head
	for nextNode != nil && !found {
		if nextNode.value == value {
			found = true
			break
		} else if strings.Compare(nextNode.value, value)  >  0 {
			found = false
			break
		}
		nextNode = nextNode.next
	}

	if found {
		return nextNode
	} else {
		return nil
	}
}

func (osl *SortedList) HasValue(value string) bool {
	return osl.FindNode(value) != nil
}

func (osl *SortedList) Remove(value string) bool {
	nodeToRemove := osl.FindNode(value)
	if nodeToRemove == nil {
		return false
	} else {
		if osl.Count == 1 {
			//Only DequeNode
			osl.Head = nil
			osl.Tail = nil
			osl.Count = 0
		} else if nodeToRemove.prev == nil {
			//Is Head
			osl.Head = nodeToRemove.next
			osl.Head.prev = nil
			osl.Count--
		} else if nodeToRemove.next == nil {
			//Is Tail
			osl.Tail = nodeToRemove.prev
			osl.Tail.next = nil
			osl.Count--
		} else {
				//Is in the middle of 2 other nodes
				nodeToRemove.prev.next = nodeToRemove.next
				nodeToRemove.next.prev = nodeToRemove.prev
				osl.Count--
		}
		return true
	}
}