package main

import (
	"fmt"
	"strconv"

	"github.com/krish8learn/DataStructures/Stack/src"
)

func main() {
	//creating stack
	s := src.Stack{}

	s.CreateStack(5)
	s.Push(19)
	s.Push(17)
	s.Push(22)
	s.Push(26)
	s.Push(8)
	//s.Push(9)
	fmt.Println("Last Item:", s.LastItem())

	ans := ""
	for _, value := range s.Print() {
		ans = ans + strconv.Itoa(value) + "--->"
	}
	fmt.Println(ans)

	s.Pop()
	ans1 := ""
	for _, value := range s.Print() {
		ans1 = ans1 + strconv.Itoa(value) + "--->"
	}
	fmt.Println(ans1)

}
