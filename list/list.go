package list

import (
	"fmt"
)

// List represents a singly linked list.
// https://en.wikipedia.org/wiki/Linked_list
type List struct {
	head, tail *Element
	len        uint
}

func (l *List) Copy(src List) {
	l.head, l.tail = src.head, src.tail
	l.len = src.len
}

// Len returns the number of elements of list l. The complexity is O(1).
func (l *List) Len() uint {
	return l.len
}

// Front returns the first element of list or nil if the list is empty.
func (l *List) Front() *Element {
	return l.head
}

// Back returns the last element of list or nil if the list is empty.
func (l *List) Back() *Element {
	return l.tail
}

// Print prints all nodes of the linked list.
func Print(l *List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
		if e.Next() != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

// PushFront inserts a new element at the front of list.
func (l *List) PushFront(v int) {
	e := NewElement(v)
	if l.len == 0 {
		l.head, l.tail = e, e
	} else {
		e.next = l.head
		l.head = e
	}
	l.len++
}

// PushBack inserts a new element at the back of list.
func (l *List) PushBack(v int) {
	e := NewElement(v)
	if l.len == 0 {
		l.head, l.tail = e, e
	} else {
		l.tail.next = e
		l.tail = e
	}
	l.len++
}

// Find returns an element with a given value from the list or nil,
// if the value is not found.
func (l *List) Find(v int) *Element {
	for n := l.Front(); n != nil; n = n.Next() {
		if n.Value == v {
			return n
		}
	}
	return nil
}

// InsertAfter inserts a new element with a given value after
// a particular element.
func (l *List) InsertAfter(v int, prev *Element) *Element {
	for n := l.Front(); n != nil; n = n.Next() {
		if prev == n {
			e := NewElement(v)
			e.next = prev.next
			prev.next = e
			return e
		}
	}
	return nil
}
