package types

type Queue struct {
	list List
}

func NewQueue(data ...any) Queue {
	return Queue{
		list: *NewList(data...),
	}
}

// Returns true if the queue has no elements.
func (Q *Queue) IsEmpty() bool {
	return Q.list.IsEmpty()
}

// Returns the number of elements in the queue.
func (Q *Queue) Size() int {
	return Q.list.Size()
}

// Inserts a new element to the end of the queue.
func (Q *Queue) Enqueue(data any) {
	Q.list.Insert(&Node{Data: data})
}

// Removes and returns the element at the front of the queue.
func (Q *Queue) Dequeue() any {
	if Q.IsEmpty() {
		return nil
	}

	r := Q.list.Remove(func(n *Node) bool {
		// The head element
		return n.previous == nil
	})

	return r.Data
}

// Returns the element at the front of the queue.
func (Q *Queue) Front() any {
	if Q.IsEmpty() {
		return nil
	}

	return Q.list.GetHead().Data
}

// Returns the element at the back of the queue.
func (Q *Queue) Back() any {
	if Q.IsEmpty() {
		return nil
	}

	return Q.list.GetTail().Data
}
