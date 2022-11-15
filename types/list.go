package types

type Node struct {
	Data     any
	previous *Node
	next     *Node
}

// A generic implementation of a LinkedList
type List struct {
	first *Node
	last  *Node
	size  int
}

func NewList(data ...any) *List {
	list := &List{}
	for _, d := range data {
		if v, ok := d.(*Node); ok {
			list.Insert(v)
		} else {
			list.Insert(&Node{Data: d})
		}
	}

	return list
}

// Returns the number of elements in the list.
func (L *List) Size() int {
	return L.size
}

// Returns true if the list has no elements.
func (L *List) IsEmpty() bool {
	return L.size == 0
}

// Returns the element at the start of the list.
func (L *List) GetHead() *Node {
	return L.first
}

// Returns the element at the end of the list.
func (L *List) GetTail() *Node {
	return L.last
}

// Insert a new element at the end of the list.
func (L *List) Insert(node *Node) {
	if L.first == nil {
		L.first = node
		L.last = node
	} else {
		node.previous = L.last
		L.last.next = node
		L.last = node
	}

	L.size += 1
}

// Excludes an element from the list.
func (L *List) Remove(matchingFunc func(*Node) bool) *Node {
	if L.IsEmpty() {
		return nil
	}

	// Handle single head node
	if L.Size() == 1 {
		if matchingFunc(L.first) {
			curr := L.first
			L.first = nil
			L.size -= 1
			return curr
		}

		return nil
	}

	curr := L.first
	for curr != nil {
		if matchingFunc(curr) {
			// Handle head node
			if curr.previous == nil {
				L.first = curr.next
				L.first.previous = nil
			} else {
				// Handle tail node.
				if curr.next == nil {
					L.last = curr.previous
					L.last.next = nil
				}

				// Handle middle node
				if curr.next != nil {
					curr.previous.next = curr.next
					curr.next.previous = curr.previous.next
				}
			}

			L.size -= 1
			return curr
		}

		curr = curr.next
	}

	return nil
}

// Traverses the list in search of the given element.
// If the node is found, the function returns a new list with this element at the end.
func (L *List) Find(matchingFunc func(*Node) bool) (result List, found bool) {
	curr := L.first

	if curr == nil {
		return result, found
	}

	for curr != nil {
		result.Insert(curr)

		if matchingFunc(curr) {
			found = true
			return result, found
		}

		curr = curr.next
	}

	return result, found
}

// Traverses the list and swaps the direction of the list such that:
//
// A -> B -> C becomes C -> B -> A
func (L *List) Reverse() {
	// Empty list.
	if L.first == nil {
		return
	}

	var previous *Node
	curr := L.first

	// Swap pointers
	for curr != nil {
		next := curr.next
		curr.next = previous
		previous = curr
		curr = next
	}

	L.last = L.first
	L.first = previous
}

// Returns all list elements as an array.
func (L *List) ToArray() []any {
	result := make([]any, 0)

	curr := L.first
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.next
	}

	return result
}
