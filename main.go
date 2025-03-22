package main

import (
	"fmt"
)

type Node struct {
	value int
	Next  *Node
}

func main() {
	head := &Node{2, &Node{ //1 task
		7, &Node{
			11, &Node{
				78, nil}}}}

	printReversedLinkedList(head)
}

func printReversedLinkedList(head *Node) {
	var result []int

	for i := head; i != nil; i = i.Next {
		result = append(result, i.value)
	}

	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	fmt.Println(result)
}
