package ordmap

// Node represents a node in the linked list
type node struct {
	data string
	next *node
}

// linkedList represents a linked list
type linkedList struct {
	head *node
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

// Empty returns true if the linked list does not contain any elements
func (ll *linkedList) Empty() bool {
	return ll.head == nil
}

// Appends adds a new a node with the given data to the end of the linked list
func (ll *linkedList) Append(data string) {
	n := &node{data: data}

	if ll.head == nil {
		ll.head = n
		return
	}

	last := ll.head
	for last.next != nil {
		last = last.next
	}

	last.next = n
}

// ToSlice returns all the items in the linked list in the order they were added
func (ll *linkedList) ToSlice() []string {
	var result []string

	current := ll.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}

	return result
}

// Delete deletes the first occurrence of a node with the given data from the linked list
func (ll *linkedList) Delete(data string) {
	if ll.head == nil {
		return
	}

	// If the node to be deleted is the head node
	if ll.head.data == data {
		ll.head.data = ""
		ll.head = ll.head.next
		return
	}

	// Find the previous node of the node to be deleted
	prev := ll.head
	for prev.next != nil && prev.next.data != data {
		prev = prev.next
	}

	// If the node with the given data is not found
	if prev.next == nil {
		return
	}

	// Remove the node from the linked list
	prev.next.data = ""
	prev.next = prev.next.next
}

// Swap moves the node with the given data to the end of the linked list
func (ll *linkedList) Swap(data string) {
	if ll.head == nil || ll.head.next == nil {
		return // If the list is empty or has only one node, no need to swap
	}

	// Handle the case where the node to be swapped is the head of the linked list
	if ll.head.data == data {
		// Find the last node
		last := ll.head
		for last.next != nil {
			last = last.next
		}
		// Move the head node to the end
		last.next = ll.head
		ll.head = ll.head.next
		last.next.next = nil
		return
	}

	// Find the node with the given data and its previous node
	current := ll.head
	var prev *node
	for current != nil && current.data != data {
		prev = current
		current = current.next
	}

	// If the node with the given data is not found or it is already at the end, do nothing
	if current == nil || current.next == nil {
		return
	}

	// Move the node to the end of the linked list
	prev.next = current.next
	current.next = nil

	// Append the moved node to the end of the linked list
	last := ll.head
	for last.next != nil {
		last = last.next
	}
	last.next = current
}
