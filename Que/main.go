package main

import (
	"fmt"

	"github.com/krish8learn/DataStructures/Que/src"
)

func main() {
	myqueue := &src.Queue{
		Fqueue: make([]string, 0),
	}

	fmt.Println("Performing Enqueue")
	myqueue.Enqueue("A")
	myqueue.Enqueue("B")
	myqueue.Enqueue("C")
	myqueue.Enqueue("D")
	fmt.Println(myqueue.Fqueue)
	myqueue.Dequeue()
	fmt.Println(myqueue.Fqueue)
	val, frontErr := myqueue.Front()
	if frontErr != nil {
		fmt.Println(frontErr)
	}
	fmt.Println(val)
	fmt.Println(myqueue.Empty(),myqueue.Length())
}
