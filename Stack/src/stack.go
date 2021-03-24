package src

import "fmt"

type Stack struct {
	slice     []int
	condition int
}

func (s *Stack) CreateStack(length int) {
	s.slice = make([]int, length)
}

func (s *Stack) Push(data int) {
	if s.condition == len(s.slice) {
		fmt.Println("Stack Overflow")
	} else if s.condition < len(s.slice) {
		s.slice[s.condition] = data
		s.condition++
		//fmt.Println(s.slice)
	}
}

func (s *Stack) Pop() {
	if s.condition == 0 {
		fmt.Println("Stack is empty, cannot remove element")
	} else if s.condition > 0 {
		//fmt.Println(s.condition)
		s.slice[s.condition-1] = 0
		s.condition--
	}
}

func (s *Stack) LastItem() int {
	return s.slice[s.condition-1]
}

func (s *Stack) Print() []int {
	var ret []int
	for i := 0; i < s.condition; i++ {
		ret = append(ret, s.slice[i])
	}
	return ret
}
