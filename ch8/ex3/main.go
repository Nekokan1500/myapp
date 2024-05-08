package main

import "fmt"

type LinkedList[T comparable] struct {
	head *Node[T]
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func (l *LinkedList[T]) Add(t T) {
	if l.head == nil {
		l.head = &Node[T]{value: t, next: nil}
	} else {
		n := l.head
		for n.next != nil {
			n = n.next
		}
		n.next = &Node[T]{value: t, next: nil}
	}
}

func (l *LinkedList[T]) Insert(t T, i int) {
	index := 0
	n := l.head
	prev := l.head
	if i < 0 {
		fmt.Println("Invalid index")
	}
	if n == nil && i > 0 {
		fmt.Println("Index out of bound")
	}
	if i == 0 {
		l.head = &Node[T]{value: t, next: l.head}
	}
	for n != nil {
		prev = n
		n = n.next
		index += 1
		if index == i {
			prev.next = &Node[T]{value: t, next: n}
			break
		}
	}
	if index < i {
		fmt.Println("Index out of bound")
	}
}

func (l LinkedList[T]) Index(t T) int {
	i := 0
	for n := l.head; n != nil; n = n.next {
		if n.value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := LinkedList[int]{head: &Node[int]{value: 25, next: nil}}
	l.Add(30)
	l.Add(40)
	l.Insert(27, 1)
	n := l.head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
	fmt.Println(l.Index(30))
}
