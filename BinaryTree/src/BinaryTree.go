package src

import (
	"errors"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

var depth int

func CreateNode(value int) *Node {
	node := &Node{
		Data:  value,
		Left:  nil,
		Right: nil,
	}
	return node
}

//Insertion
func (b *BinaryTree) Insert(value int) {
	//Incase new node is the 1st node
	if b.Root == nil {
		b.Root = CreateNode(value)
	} else {
		b.Root.Insert(value)
	}
}

func (n *Node) Insert(value int) error {
	if n.Data == value {
		return errors.New("This node already exists")
	} else if n.Data != value {
		if value < n.Data {
			if n.Left == nil {
				n.Left = CreateNode(value)
				return nil
			} else {
				n.Left.Insert(value)
			}
		} else if value > n.Data {
			if n.Right == nil {
				n.Right = CreateNode(value)
				return nil
			} else {
				n.Right.Insert(value)
			}
		}
	}
	return nil
}

//Traversal--(Post,Pre,In)order
//PreOrder
func PreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Data)
		PreOrder(n.Left)
		PreOrder(n.Right)
	}
}

//Inorder
func InOrder(n *Node) {
	if n == nil {
		return
	} else {
		InOrder(n.Left)
		fmt.Printf("%d ", n.Data)
		InOrder(n.Right)
	}
}

//PostOrder
func PostOrder(n *Node) {
	if n == nil {
		return
	} else {
		PostOrder(n.Left)
		PostOrder(n.Right)
		fmt.Printf("%d ", n.Data)
	}
}

//Finding Min and Max
//MAXIMUM
func (b *BinaryTree) FindMax() *Node {
	//for root node
	if b.Root.Right == nil {
		return b.Root
	}
	return b.Root.Right.FindMax()
}

func (n *Node) FindMax() *Node {
	if n.Right == nil {
		return n
	}
	return n.Right.FindMax()
}

//MINIMUM
func (b *BinaryTree) FindMin() *Node {
	//for root node
	if b.Root.Left == nil {
		return b.Root
	}
	return b.Root.Left.FindMin()
}

func (n *Node) FindMin() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.FindMin()
}

//searching
func (b *BinaryTree) Search(value int) {
	if b == nil {
		fmt.Println("Empty Tree")
	}
	if b.Root.Data == value {
		fmt.Printf("Value present at depth, %d \n", depth)
	}
	depth++
	b.Root.Search(value)
}

func (n *Node) Search(value int) {
	if value > n.Data {
		if n.Right == nil {
			fmt.Println("The value not present in the tree")
		} else if n.Right != nil {
			if n.Right.Data == value {
				fmt.Printf("Value present at depth, %d \n", depth)
			} else {
				depth++
				n.Right.Search(value)
			}
		}
	} else if value < n.Data {
		if n.Left == nil {
			fmt.Println("The value not present in the tree")
		} else if n.Left != nil {
			if n.Left.Data == value {
				fmt.Printf("Value present at depth, %d \n", depth)
			} else {
				depth++
				n.Left.Search(value)
			}
		}
	}
}

func (b *BinaryTree) Delete(value int) {
	if b.Root == nil {
		fmt.Println("The Tree is Empty")
		//return nil
	} else if value < b.Root.Data {
		b.Root.Left = b.Root.Left.Delete(value)
	} else if value > b.Root.Data {
		b.Root.Right = b.Root.Right.Delete(value)
	} else if b.Root.Data == value {
		//root node has no child
		if b.Root.Left == nil && b.Root.Right == nil {
			b.Root = nil
			fmt.Println("The value was in the Root node, Deleted, now the tree is empty")
		} else if b.Root.Left == nil {
			//the left node of root is nil
			temp := b
			b.Root = b.Root.Right
			temp.Root = nil
		} else if b.Root.Right == nil {
			//the right node of root is nil
			temp := b
			b.Root = b.Root.Left
			temp.Root = nil
		} else if b.Root.Left != nil && b.Root.Right != nil {
			temp := b.Root.Right.FindMin()
			b.Root.Data = temp.Data
			b.Root.Right = b.Root.Right.Delete(value)
		}
	}
}

func (n *Node) Delete(value int) *Node {
	//after the root node we will perform same opeartion in the sub nodes/tress
	if n == nil {
		return n
	} else if n.Data == value {
		if n.Right == nil && n.Left == nil {
			//means it is a leaf node without any children
			n = nil
		} else if n.Left != nil && n.Right != nil {
			temp := n.Right.FindMin()
			n.Data = temp.Data
			n.Right = n.Right.Delete(value)
		} else if n.Left == nil {
			temp := n
			n = n.Right
			temp.Data = 0
			temp = nil
		} else if n.Right == nil {
			temp := n
			n = n.Left
			temp.Data = 0
			temp = nil
		}
	} else if value < n.Data {
		n.Left = n.Left.Delete(value)
	} else if value > n.Data {
		n.Right = n.Right.Delete(value)
	}
	return n
}
