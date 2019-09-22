package dll

import (
	"errors"
)

// DLItem DoubleLinkedList Item
type DLItem struct {
	data interface{}
	prev *DLItem
	next *DLItem
}

// Prev function returns previous item
func (item *DLItem) Prev() (*DLItem, error) {
	if item == nil {
		return nil, errors.New("Can't get previous item, because item is nil")
	}
	return item.prev, nil
}

// Next function returns next item
func (item *DLItem) Next() (*DLItem, error) {
	if item == nil {
		return nil, errors.New("Can't get next item, because item is nil")
	}

	return item.next, nil
}

// Value function returns Data
func (item *DLItem) Value() interface{} {
	return item.data
}

// DLList DoubleLinkedList container
type DLList struct {
	first *DLItem
	last  *DLItem
}

// Len function returns length of list
func (list DLList) Len() int {

	if list.First() == nil {
		return 0
	}

	length := 0
	item := list.First()
	for item != nil {
		item, _ = item.Next()
		length++
	}

	return length
}

// First returns first element of list
func (list DLList) First() *DLItem {
	return list.first
}

// Last returns last element of list
func (list DLList) Last() *DLItem {
	return list.last
}

// PushFront function adds item to the front of list
func (list *DLList) PushFront(v interface{}) {
	var item DLItem
	item.data = v

	if list.first == nil {
		list.first = &item
		list.last = &item
	} else {
		item.next = list.first
		list.first.prev = &item
		list.first = &item
	}
}

// PushBack function adds item to the back of list
func (list *DLList) PushBack(v interface{}) {
	var item DLItem
	item.data = v

	if list.first == nil {
		list.first = &item
		list.last = &item
	} else {
		item.prev = list.last
		list.last.next = &item
		list.last = &item
	}
}

// Remove function removes item from list
func (list *DLList) Remove(i *DLItem) {

	if list.first == i {
		list.first = i.next
		i.next.prev = nil
	} else {
		if list.last == i {
			list.last = i.prev
			i.prev.next = nil
		} else {
			i.prev.next = i.next
			i.next.prev = i.prev
		}
	}
}
