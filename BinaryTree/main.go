package main

import (
	"fmt"

	"github.com/krish8learn/DataStructures/BinaryTree/src"
)

func main() {
	tree := src.BinaryTree{}

	fmt.Println("Insertion")
	tree.Insert(12)
	tree.Insert(67)
	tree.Insert(65)
	tree.Insert(77)
	tree.Insert(23)
	tree.Insert(44)
	tree.Insert(89)
	tree.Insert(55)
	tree.Insert(9)
	tree.Insert(4)

	fmt.Printf("PreOrder:")
	src.PreOrder(tree.Root)
	fmt.Println()
	fmt.Printf("PostOrder:")
	src.PostOrder(tree.Root)
	fmt.Println()
	fmt.Printf("InOrder:")
	src.InOrder(tree.Root)
	fmt.Println()

	fmt.Println("Finding MAX and MIN")
	fmt.Println("MAXIMUM", tree.FindMax().Data)
	fmt.Println("MINIMUM", tree.FindMin().Data)

	fmt.Println("Searching")
	tree.Search(4)

	fmt.Println("DELETION")
	tree.Delete(12)
	src.InOrder(tree.Root)
	fmt.Println()
}
