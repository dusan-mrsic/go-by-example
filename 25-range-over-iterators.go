package main

import (
	"fmt"
	"iter"
)

type LinkedList[T any] struct {
	head, tail *elem[T]
}

type elem[T any] struct {
	next *elem[T]
	val  T
}

func (lst *LinkedList[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &elem[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &elem[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *LinkedList[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func main() {
	lst := LinkedList[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	for e := range lst.All() {
		fmt.Println(e)
	}
}
