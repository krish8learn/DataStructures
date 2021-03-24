package main

import (
	"fmt"

	"github.com/krish8learn/DataStructures/LinkedList/src"
)

func main() {
	l := src.LinkedList{}
	fmt.Println("INSERTION")
	l.InsertOne(12)
	l.InsertOne(89)
	l.InsertOne(76)
	l.InsertOne(98)
	l.InsertOne(67)
	l.InsertOne(88)
	l.InsertOne(54)
	fmt.Println("PRINTING DATA")
	maps, err := l.ShowData()
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, val := range maps {
		fmt.Println(key, ",", val)
	}
	fmt.Println("GETTING LENGTH")
	fmt.Println(l.GetLength())

	fmt.Println("Finding Position of Given element")
	position, position_error := l.GetAt(89)
	if position_error != nil {
		fmt.Println(position_error)
		return
	}
	fmt.Println(position)
	//fmt.Println("INSERT ALL ELEMENTS AT ONCE")
	/*var slice = []int{1, 3, 5, 2, 9}
	l.InsertAll(slice)

	maps2, err2 := l.ShowData()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	for key2, val2 := range maps2 {
		fmt.Println(key2, ",", val2)
	}*/
	fmt.Println("DELETING ELEMENT")
	deleteErr := l.Delete(88)
	if deleteErr != nil {
		fmt.Print(deleteErr)
	}
	maps2, err2 := l.ShowData()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	for key2, val2 := range maps2 {
		fmt.Println(key2, ",", val2)
	}
}
