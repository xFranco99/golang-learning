package main

import (
	"strings"
)

type MyLinkedList struct {
	head *Element
	tail *Element
}

func newMyLinkedList(elem Element) MyLinkedList {
	return MyLinkedList{
		head: &elem,
		tail: &elem,
	}
}

func (linkedList *MyLinkedList) addElement(elem Element) {
	if linkedList.head == nil || linkedList.tail == nil {
		tmpLL := newMyLinkedList(elem)
		linkedList = &tmpLL
	} else {
		linkedList.tail.next = &elem
		linkedList.tail = linkedList.tail.next
	}
}

func newMyLinkedListFomValues(vals ...int) MyLinkedList {

	elem := newElement(vals[0])
	ll := newMyLinkedList(elem)

	for i := 1; i < len(vals); i++ {
		ll.addElement(newElement(vals[i]))
	}

	return ll
}

func (linklinkedList *MyLinkedList) myLikedListToString() string{
	var llString strings.Builder

	llString.WriteString("[")

	var pointer *Element = linklinkedList.head
	for pointer != nil {
		llString.WriteString(pointer.label)

		if pointer.next != nil {
			llString.WriteString(", ")
		}

		pointer = pointer.next
	}

	llString.WriteString("]")
	return llString.String()
}