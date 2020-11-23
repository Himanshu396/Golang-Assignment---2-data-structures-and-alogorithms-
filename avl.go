package main

import (
	"fmt"
)

type Node struct {
	Data   int
	left   *Node
	right  *Node
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func height(n *Node) int {
	if n == nil {
		return 0
	} else {
		return n.height
	}
}

//newNode := Node{data, nil, nil, 1}

func rightRotate(y *Node) *Node {
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2
	// update height
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x

}
func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left
	y.left = x
	x.right = t2
	// update height
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y

}
func getBalance(n *Node) int {
	if n == nil {
		return 0
	} else {
		return height(n.left) - height(n.right)
	}
}

func insert(n *Node, data int) *Node {
	if n == nil {
		newNode := Node{data, nil, nil, 1}
		return &newNode
	}
	if data < n.Data {
		n.left = insert(n.left, data)
	} else if data > n.Data {
		n.right = insert(n.right, data)
	} else {
		return n //duplicate data  are not allow in bst
	}
	//node->height = 1 + max(height(node->left), height(node->right));
	n.height = 1 + max(height(n.left), height(n.right))
	balance := getBalance(n)
	//fmt.Println(balance)
	// LL
	if balance > 1 && n.left != nil && data < n.left.Data {

		fmt.Println("balance 1")
		return rightRotate(n)
	}

	//RR
	if balance < -1 && n.right != nil && data > n.right.Data {

		fmt.Println("balance 2")
		return leftRotate(n)
	}

	// LR
	if balance > 1 && n.left != nil && data > n.left.Data {
		n.left = leftRotate(n.left)

		fmt.Println("balance 3")
		return rightRotate(n)
	}

	//RL
	if balance < -1 && n.right != nil && data < n.right.Data {
		n.right = rightRotate(n.right)
		fmt.Println("balance 4")
		return leftRotate(n)
	}
	return n
}

func preOrder(n *Node) {
	if n != nil {
		fmt.Print(n.Data, " ")
		preOrder(n.left)
		preOrder(n.right)

	}
}

func main() {
	var root *Node = nil

	/* Constructing tree given in
	   the above figure */
	root = insert(root, 10)
	root = insert(root, 20)
	root = insert(root, 30)
	root = insert(root, 40)
	root = insert(root, 50)
	root = insert(root, 25)
	//fmt.Println(root)
	/* The constructed AVL Tree would be
	           30
	       / \
	       20 40
	       / \ \
	   10 25 50
	*/
	fmt.Println("Preorder traversal of the constructed AVL tree is \n")
	preOrder(root)

	//return 0;
}