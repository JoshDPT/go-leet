package main

import "errors"

// using linked list to make an efficient queue
type Node struct {
    Val byte
    Next *Node
}

// defining the queue data structure
type ByteQueue struct {
	// head & tails are pointers to nodes
	head *Node
	tail *Node
	size int
}

func (q * ByteQueue) Enqueue (val byte) {
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

func (q * ByteQueue) Dequeue () (byte, error) {
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

func (q * ByteQueue) IsEmpty () bool {
	return q.size == 0
}

func (q * ByteQueue) Size () int {
	return q.size
}