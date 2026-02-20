package main

import "errors"

// // import "fmt"

// type Node struct {
// 	keys []string
// 	children []*Node
// 	leaf bool
// }

// type BTree struct {
// 	root *Node
// 	t int32
// }

// func search(key string, b *BTree, node *Node) (*Node, int32, error) {
// 	if (node == nil){
// 		node = b.root
// 	}

// 	i := 0
// 	for i < len(node.keys) && key > node.keys[i] {
// 		i +=1
// 	}
// 	if i < len(node.keys) && key == node.keys[i] {
// 		return node, int32(i), nil
// 	} else if node.leaf {
// 		return node, int32(0), nil
// 	} else {
// 		return search(key, b, node.children[i])
// 	}
// }

// func splitChild(b *BTree, x *Node, i int32) {
// 	t := b.t

// 	y := x.children[i]
// 	z := Node(y.leaf)
// }

// // func NewNode(n *Node) *Node{
// // 	return &Node{
// // 		key,
// // 		children,
// // 		false,
// // 	}
// // }
//

const (
	degree = 3
	maxChildren = 2 * degree // 6
	maxItems = maxChildren - 1
	minItems = degree - 1
)

type BTree struct {
	root *node
}

func NewBTree() *BTree{
	return &BTree{}
}

func (t *BTree) Find(key []byte) ([]byte, error) {
	// while next is not None essentially.
	for next := t.root; next != nil; {
		
		// get the current position and whether it's been found or not.
		pos, found := next.search(key)
		
		// if found, return the postion of the item in the node.
		if found {
			return next.items[pos].val, nil
		}
		
		// else, update the node with the next child node with the pointer to that position
		next = next.children[pos]
	}
	return nil, errors.New("key not found.")
}

func (t *BTree) splitRoot() {
	newRoot := &node{}
	midItem, newNode := t.root.split()
	newRoot.insertItemAt(0, midItem) // newRoot Node
	newRoot.insertChildAt(0, t.root) // newRoot left child
	newRoot.insertChildAt(1, newNode) // newRoot right child
	
	t.root = newRoot
}

func (t *BTree) Insert(key, val []byte) {
	i := &item{key, val}
	
	// if the root is empty, create a new node.
	if t.root == nil {
		t.root = &node{}
	}
	
	// if the root is full, split the root.
	if t.root.numItems >= maxItems {
		t.splitRoot()
	}
	
	// start insertion.
	t.root.insert(i)
	
}

