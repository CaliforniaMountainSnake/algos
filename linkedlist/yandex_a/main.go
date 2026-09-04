package main

import (
	"fmt"
	"io"
	"os"
)

// Универсальная нода для хранения чего угодно.
type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	list := LinkedList[int]{
		Head: nil,
	}

	// Fscan поочередно парсит "следующее значение" или несколько таковых из ридера.
	// Значения считаются разделенными пробельными символами.
	// Fscan позволяет указать тип соответствующего значения.
	var countOfQueries int
	fmt.Fscan(in, &countOfQueries)
	for i := 0; i < countOfQueries; i++ {
		var queryType int
		fmt.Fscan(in, &queryType)

		switch queryType {
		case 1:
			var x, y int
			fmt.Fscan(in, &x, &y)
			executeQueryOne(&list, x, y)
		case 2:
			var x int
			fmt.Fscan(in, &x)
			executeQueryTwo(&list, x, out)
		case 3:
			var x int
			fmt.Fscan(in, &x)
			executeQueryThree(&list, x)
		}
	}
}

// Запрос 1-ого типа: добавить число y после x-ого числа в списке.
// Если x=0 , то нужно сделать число y новым началом списка.
func executeQueryOne(list *LinkedList[int], x int, y int) {
	if x == 0 {
		newHead := Node[int]{
			Data: y,
			Next: list.Head,
		}
		list.Head = &newHead
		return
	}

	// Перемещаемся к нужной ноде за O(n).
	node := list.Head
	for i := range x {
		if i == x-1 {
			// добавить число y после x-ого числа в списке
			// у этой ноды next должна стать новая нода. А у новой ноды next должна стать next этой ноды.
			newNode := Node[int]{
				Data: y,
				Next: node.Next,
			}
			node.Next = &newNode
			return
		}
		node = node.Next
	}
}

// Запрос 2-ого типа: вывести число, которое находится на позиции x в списке.
// "Позиция" - это номер, натуральное число, не индекс.
func executeQueryTwo(list *LinkedList[int], x int, out io.Writer) {
	// Перемещаемся к нужной ноде за O(n).
	node := list.Head
	for i := range x {
		if i == x-1 {
			fmt.Fprintln(out, node.Data)
			return
		}
		node = node.Next
	}
}

// Запрос 3-его типа: удалить число, которое находится на позиции x в списке
func executeQueryThree(list *LinkedList[int], x int) {
	// Перемещаемся к нужной ноде за O(n).
	var prevNode *Node[int] = nil
	node := list.Head
	for i := range x {
		if i == x-1 {
			// удалить ноду
			if prevNode == nil {
				// надо удалить самый первый элемент списка
				list.Head = node.Next
				return
			}
			prevNode.Next = node.Next
		}
		prevNode = node
		node = node.Next
	}
}
