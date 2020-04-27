package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func main() {
	var q deque.Deque
	q.PushBack(1)
	q.PushBack(2)
	q.PushFront(3)

	fmt.Printf("Size: %d\n", q.Len()) // Prints: 3
	// front to back
	for j := 0; j < q.Len(); j++ {
		fmt.Printf("%d, ", q.At(j))
	}
	fmt.Printf("\n")
	q.PopFront() // remove 3
	q.PopBack()  // remove 2

	fmt.Printf("Size: %d\n", q.Len()) // Prints: 1
	for j := 0; j < q.Len(); j++ {
		fmt.Printf("%d, ", q.At(j))
	}
	fmt.Printf("\n")
}
