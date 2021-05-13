package main

import (
	"fmt"

	"github.com/krish8learn/DataStructures/Heap/src"
)

func main() {
	instance := src.Heap{}

	elements := []int{13, 20, 22, 17, 89, 65, 1, 23, 85, 45, 77}

	for _, val := range elements {
		instance.Insert(val)
		fmt.Println(instance)
	}

	fmt.Println(instance.Extract())

	fmt.Println(instance)
}
