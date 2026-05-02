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

func (n *node) removeItemAt(pos int) *item {
	removedItem := n.items[pos] // getting the value of the key at that position in the node.
	n.items[pos] = nil // updating the key value to none, i.e removing the key.
	
	 if lastPos := n.numItems - 1; pos < lastPos {
			/* rearranging the keys in the node */
			copy(n.items[pos:lastPos], n.items[pos+1:lastPos+1])
			n.items[lastPos] = nil
		}
	n.numItems--
	return removedItem
}

func (n *node) removeChildAt(pos int) *node {
	removedChild := n.children[pos]
	n.children[pos] = nil
	
	// fill the gap, if the position we are removing from is not the very last occupiedin the "children" array.
	if lastPos := n.numChildren - 1; pos < lastPos {
		copy(n.children[pos:lastPos], n.children[pos+1:lastPos+1])
		n.children[lastPos] = nil
	}
	n.numChildren--
	
	return removedChild
}

func (n *node) fillChildAt(pos int) {
	switch {
		// borrow the right-most item from the left sibling if the left
		// sibling exists and has more than the minimum number of items.
		case pos > 0 && n.children[pos-1].numItems > minItems:
			// establish our left and right nodes.
			left, right := n.children[pos-1], n.children[pos]
			// take the item from the parent and place it at the left-most position of the right node.
			copy(right.items[1:right.numItems+1], right.items[:right.numItems])
			right.items[0] = n.items[pos-1]
			right.numItems++
		
		// for non-leaf nodes, make the right-most child of the left node the new left-most child of the right node.
		if !right.isLeaf(){
			right.insertChildAt(0, left.removeChildAt(left.numChildren-1))
		}
		// borrow the right-most item from the left node to r eplace 
		n.items[pos-1] = left.removeItemAt(left.numItems - 1)
		
		// borrow the left-most item from the right sibling if the right 
		// sibling exists and has more than the minimum number of items.
		case pos < n.numChildren - 1 && n.children[pos+1].numItems > minItems:
			// establish our left and right nodes.
			left, right := n.children[pos], n.children[pos+1]
			
			// take the item from the parent and place it at the right-most position of the left node.
			left.items[left.numItems] = n.items[pos]
			left.numItems++
			
			// for non-leaf nodes, make the left-most child of the right node the new right child of the left node.
			if !left.isLeaf() {
				left.insertChildAt(left.numChildren, right.removeChildAt(0))
			}
			// borrow the left-most item from the right node to replace the parent item.
			n.items[pos] = right.removeItemAt(0)
		
		// there are no suitable items to borrow a node from, so perform a merge.
		default:
			// if we are the right-most child pointer, merge the node with its left siblings.
			// in all other cases, we prefer to merge the node with its right sibling for simplicity.
			if pos >= n.numItems{
				pos = n.numItems - 1
			}
			// establish our left and right nodes.
			left, right := n.children[pos], n.children[pos+1]
			// borrow an item from the parent node and place it at the right-most available position of the left node.
			left.items[left.numItems] = n.removeItemAt(pos)
			left.numItems++
			
			// migrate all items from the right node to the left node.
			copy(left.items[left.numItems:], right.items[:right.numItems])
			left.numItems += right.numItems
			
			// for non-left nodes, migrate all applicable children from the right node to the left node.
			if !left.isLeaf(){
				copy(left.children[left.numChildren:], right.children[:right.numChildren])
				left.numChildren += right.numChildren
			}	
			// remove the child pointer from the parent to the right node and discard the right node.
			n.removeChildAt(pos + 1)
			right = nil
	}
}

func (n *node) delete(key []byte, isSeekingSuccessor bool) *item {
	pos, found := n.search(key)
	
	var next *node
	
	// we have found a node holding an item matching the supplied key.
	if found {
		// this is a leaf node, so we can simply remove the item.
		if n.isLeaf() {
			return n.removeItemAt(pos)
		}
		next, isSeekingSuccessor = n.children[pos+1], true
	} else {
		next = n.children[pos]
	}
	
	// we have reached the leaf node containing the inorder successor, so remove the successor from the leaf.
	if n.isLeaf() && isSeekingSuccessor {
		return n.removeItemAt(0)
	}
	
	// we were unable to find the item matching the given key. dont' do anything.
	if next == nil {
		return nil
	}
	
	// continue traversing the tree to find an item matching the supplied key.
	deletedItem := next.delete(key, isSeekingSuccessor)
	
	// we found the inorder successor, and we are now back at the internal node containing the item
	// matching the supplied key. therefore, we replace the item with its inorder successor, effectively
	// deleting the item from the tree
	if found && isSeekingSuccessor{
		n.items[pos] = deletedItem
	}
	
	// check if an underflow occurred after we deleted an item down the tree.
	if next.numItems < minItems{
		// repair the underflow.
		if found && isSeekingSuccessor{
			n.fillChildAt(pos + 1)
		} else {
			n.fillChildAt(pos)
		}
	}
	
	// propagate the deleted item back to the previous stack frame.
	return deletedItem
}
