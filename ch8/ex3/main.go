package ex3

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
	if n != nil && i == index {

	}
}
