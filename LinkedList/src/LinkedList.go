package src

import (
	"errors"
	"fmt"
)

//single linked list

type Node struct {
	Pointer *Node
	Value   int
}

type LinkedList struct {
	head   *Node
	length int
}

func CreateNode(data int) *Node {
	node := Node{
		Pointer: nil,
		Value:   data,
	}
	return &node
}

func (l *LinkedList) InsertOne(data int) {
	temp := l.head
	if l.head == nil {
		//for 1st node when temp is not pointing to anynode
		l.head = CreateNode(data)
		l.length++
	} else if temp != nil {
		//means linked list has nodes
		for temp.Pointer != nil {
			temp = temp.Pointer
		}
		temp.Pointer = CreateNode(data)
		l.length++
	}
}

func (l *LinkedList) ShowData() (map[int]int, error) {
	linkedlistslice := make(map[int]int)
	var position int
	temp := l.head
	if l.head == nil {
		return nil, errors.New("The LINKED LIST IS EMPTY")
	} else if temp != nil {
		for temp.Pointer != nil {
			linkedlistslice[position] = temp.Value
			position++
			temp = temp.Pointer
		}
		//THIS TAKES THE LAST VALUE IN LINKED LIST
		linkedlistslice[position] = temp.Value
	}
	return linkedlistslice, nil
}

func (l *LinkedList) GetLength() int {
	return l.length
}

func (l *LinkedList) GetAt(data int) (int, error) {
	//var condition bool = false
	temp := l.head
	for i := 0; i < l.length; i++ {
		if data == temp.Value {
			//condition = true
			return i, nil
		}
		temp = temp.Pointer
	}
	return 0, errors.New("NO DATA PRESENT")
}

func (l *LinkedList) InsertAll(slice []int) {
	temp := l.head
	for _, val := range slice {
		if temp == nil {
			temp = CreateNode(val)
			l.length++
		} else if temp != nil {
			for temp.Pointer != nil {
				temp = temp.Pointer
			}
			temp.Pointer = CreateNode(val)
			l.length++
		}
	}
}

func (l *LinkedList) Delete(data int) error {
	var prevnode *Node
	temp := l.head
	if l.length == 0 {
		return errors.New("EMPTY LINKED LIST")
	}
	for i := 0; i < l.length; i++ {
		if data == temp.Value {
			prevnode.Pointer = temp.Pointer
			temp.Pointer = nil
			break
		}
		prevnode = temp
		temp = temp.Pointer
	}
	fmt.Println("SUCCESSFUL DELETION")
	return nil
}
