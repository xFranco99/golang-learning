package main

import "fmt"

type Element struct {
	val int
	label string
	next *Element
}

func newElement(value int) Element {
	return Element{
		val: value,
		label: fmt.Sprintf("%d", value),
		next: nil,
	}
}