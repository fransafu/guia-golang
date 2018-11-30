package main

import (
	"fmt"
)

// Node estructura basica de una lista
type Node struct {
	value int
	next  *Node
}

// List of node
type List struct {
	length int
	head   *Node
}

// Append node to list
func (l *List) Append(node *Node) {
	if l.length == 0 {
		l.head = node
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = node
	}
	l.length++
}

func (l *List) printInfo(index int, valueNode int, nextMemory Node) {
	fmt.Printf("Nodo: %d\n\tvalor: %d\n\tSiguiente espacio de memoria: %v\n", index, valueNode, &nextMemory.next)
}

// PrintList imprime la lista enlazada simple
func PrintList(l *List) {
	currentNode := l.head
	index := 1
	for currentNode.next != nil {
		l.printInfo(index, currentNode.value, *currentNode)
		currentNode = currentNode.next
		index++
	}
	l.printInfo(index, currentNode.value, *currentNode)
}

// Este ejercicio consiste en realizar una lista enlazada simple
func main() {
	fmt.Println("Lista enlazada simple")

	l := &List{}
	node1 := Node{value: 1}
	node2 := Node{value: 2}
	node3 := Node{value: 4}
	node4 := Node{value: 6}
	node5 := Node{value: 8}
	l.Append(&node1)
	l.Append(&node2)
	l.Append(&node3)
	l.Append(&node4)
	l.Append(&node5)

	PrintList(l)
}
