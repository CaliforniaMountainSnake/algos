package linkedlist

// Универсальная нода для хранения чего угодно.
type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}
