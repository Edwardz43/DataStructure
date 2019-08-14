package main

import "log"

// Node ...
type Node struct {
	// parent *Node
	left  *Node
	right *Node
	data  int
}

func insertRecursive(v int, n *Node) *Node {
	if n == nil {
		return &Node{
			// parent: n,
			data: v,
		}
	}
	if v < n.data {
		n.left = insertRecursive(v, n.left)
	} else if v > n.data {
		n.right = insertRecursive(v, n.right)
	} else {
		return n
	}
	return n
}

// InsertOne inserts new data to tree.
func (n *Node) InsertOne(value int) {
	insertRecursive(value, n)
}

// Insert inserts array of data to tree.
func (n *Node) Insert(values []int) {
	for _, v := range values {
		n.InsertOne(v)
	}
}

func (n *Node) preorderTriverse() {
	log.Println(n.data)
	if n.left != nil {
		n.left.preorderTriverse()
	}
	if n.right != nil {
		n.right.preorderTriverse()
	}
}

func (n *Node) inorderTriverse() {
	if n.left != nil {
		n.left.preorderTriverse()
	}
	log.Println(n.data)
	if n.right != nil {
		n.right.preorderTriverse()
	}
}

func (n *Node) postorderTriverse() {
	if n.left != nil {
		n.left.preorderTriverse()
	}
	if n.right != nil {
		n.right.preorderTriverse()
	}
	log.Println(n.data)
}

// Triverse triverses all data from tree.
func (n *Node) Triverse(t string) {
	switch t {
	case "preorder":
		n.preorderTriverse()
		break
	case "inorder":
		n.inorderTriverse()
		break
	case "postorder":
		n.postorderTriverse()
		break
	}
}

func main() {
	root := Node{data: 4}

	data := []int{5, 2, 3, 1, 6, 4, 8}

	root.Insert(data)

	root.Triverse("preorder")
}
