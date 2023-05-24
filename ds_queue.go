package main

import (
	"errors"
)

// using linked list to make an efficient queue
type Node struct {
	// I personally don't love doing this, but it allows for many data types
    Val interface{}
    Next *Node
}

// defining the queue data structure
type Queue struct {
	// head & tails are pointers to nodes
	head *Node
	tail *Node
	size int
}

func (q * Queue) Enqueue (val interface{}) {
	// ampersand is vital to maintain pointers for next
	node := &Node{val, nil}
	// this means it is empty, could use isEmpty()
	if q.tail == nil {
		q.head = node
		q.tail = node
		// add it to the end  
	} else {
		q.tail.Next = node
		q.tail = q.tail.Next
	}
	q.size+=1
}

func (q * Queue) Dequeue () (interface{}, error) {
	if q.head == nil {
		return 0, errors.New("No element to dequeue")
	}
	data := q.head.Val
	q.head = q.head.Next

	if q.head == nil {
		q.tail = nil
	}
	q.size -= 1
	return data, nil
}

func (q * Queue) IsEmpty () bool {
	return q.size == 0
}

func (q * Queue) Size () int {
	return q.size
}