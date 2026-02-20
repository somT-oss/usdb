package main

import "bytes"

type item struct {
	key []byte
	val []byte
}

type node struct {
	items [maxItems] *item
	children [maxChildren] *node
	numItems int
	numChildren int
}

// checks if a node is a leaf node
/*
 * B trees have a rule where leaf nodes have no child node and 
 * contain items[key, value]. 
 */ 
func (n *node) isLeaf() bool {
	return n.numChildren == 0
}

func (n *node) search(key []byte) (int, bool) {
	
	// for a particular node, the left idx is 0
	// the right most idx is the max number of items, numItems.
	low, high := 0, n.numItems 
	var mid int
	for low < high {
		mid = (low + high) / 2
		
		// check to see if the key at mid is equal to the search key
		compare := bytes.Compare(key, n.items[mid].key)
		
		// read the function of the Compare method to understand more about this logic!
		switch {
			case compare > 0:
				low = mid + 1
			case compare < 0:
				high = mid
			case compare == 0:
				return mid, true
		}
	}
	return low, false
} 

func (n *node) insertItemAt(pos int, i *item) {
	if pos < n.numItems{
		// increase the length of the current item arr of this node, if the position we want to insert in is greater than 
		// the length of the item array of this node.
		copy(n.items[pos+1: n.numItems+1], n.items[pos:n.numItems])
	}
	n.items[pos] = i
	n.numItems++
}

func (n *node) insertChildAt(pos int, c *node) {
	if pos < n.numChildren {
		/* increase the length of the current children pointer array of this node, if the position we want to insert in is greater
	 	   than the length of the children pointer array of this node.
		*/
		copy(n.children[pos+1:n.numChildren+1], n.children[pos:n.numChildren])
	}
	n.children[pos] = c
	n.numChildren++
}

func (n *node) split() (*item, *node) {
	//get the middle item
	mid := minItems
	midItem := n.items[mid]
	
	// create a new node and copy half of the items from the current node to the new node
	newNode := &node{}
	copy(newNode.items[:], n.items[mid+1:])
	newNode.numItems = minItems
	
	// if necessary, copy half of the child pointers from the current node to the new node
	if !n.isLeaf() {
		copy(newNode.children[:], n.children[mid+1:])
		newNode.numChildren = minItems + 1
	}
	
	// remove data items and child pointers from the current node that have been added to the new node.
	for i, l := mid, n.numItems; i < l; i ++ {
		n.items[i] = nil
		n.numItems--
		
		if !n.isLeaf() {
			n.children[i+1] = nil
			n.numChildren--
		}
	}
	
	return midItem, newNode
}

// TODO: READ THIS FUNCTION AND UNDERSTAND PROPERLY
func (n *node) insert(item *item) bool {
	pos, found := n.search(item.key)
	
	// the data item already exists, so just update the value
	if found {
		n.items[pos] = item
		return false
	}
	
	// if the node is a leaf node, insert the item at the specified index.
	if n.isLeaf() {
		n.insertItemAt(pos, item)
		return true // successful insertion.
	}
	
	// check if the array of items of the node at n.children[pos] is greater than the number of maxItems, split.
	// after splitting, insert the midItem at the position where the search item[key:value] was found
	// insert the child at +1 of the position the searched item[key:value] was found. This will give the idea of
	/*
	 * 
		* [   item, item,  item    ]
	        /	   |      |     \
		   /       |      |      \
		  /        |      |       \
	    [node,     node,  node,    node]
	*/ 
	if n.children[pos].numItems >= maxItems {
		midItem, newNode := n.children[pos].split()
		n.insertItemAt(pos, midItem)
		n.insertChildAt(pos+1, newNode)
		
		switch cmp := bytes.Compare(item.key, n.items[pos].key); {
			case cmp < 0:
			
			case cmp > 0:
				pos ++
			
			default:
				n.items[pos] = item
				return true
		}
		
	}
	return n.children[pos].insert(item)
}