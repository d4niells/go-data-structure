package queue

import "fmt"

// Queues are a special list, just like stacks. Unlike
// stacks, however, queues are known for being FIFO data structures.
// This means that whatever item that goes in first comes out first.
// You can imagine a line for the bus, where the first person in the line
// will get on the bus first, while the rest have to wait until their turn.
// A queue is simply a list where you can only insert at the tail (back) and
// pop at the head (front).

// Popular operations that can be done on queues include:
// - Enqueuing, aka Pushing or Inserting
// - Dequeuing, aka Popping or Deleting
// - Front
// - IsEmpty or IsFull

type Queue struct {
	items []int
}

// We are inserting new elements to the tail end of q.items
func (q *Queue) Enqueuing(data int) {
	q.items = append(q.items, data)
}

// We are just taking a slice of q.items from the 1st element to the last
func (q *Queue) Dequeuing() {
	q.items = q.items[1:]
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}

	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Print() {
	for _, item := range q.items {
		fmt.Print(item, " ")
	}
	
	fmt.Print()
}