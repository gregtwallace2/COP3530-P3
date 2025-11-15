package binarysearchtree

// 3rd data struct; not one of the two required

// node is used to store data in the BST
type node struct {
	key   string
	value uint64

	left  *node
	right *node
}

// BinarySearchTree is the actual BST struct. It is sorted by nodes' keys
type BinarySearchTree struct {
	head *node
}

// NewBinarySearchTree returns an empty BST
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert inserts a key: value pair into the tree.
func (bst *BinarySearchTree) Insert(key string, value uint64) {
	// first node
	if bst.head == nil {
		bst.head = &node{
			key:   key,
			value: value,
		}

		return
	}

	currentNode := bst.head
	for {
		if key < currentNode.key {
			if currentNode.left == nil {
				currentNode.left = &node{
					key:   key,
					value: value,
				}

				return
			}

			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = &node{
					key:   key,
					value: value,
				}

				return
			}

			currentNode = currentNode.right
		}
	}
}

// Search finds the specified key in the bst and returns the value associated
// with it
func (bst *BinarySearchTree) Search(key string) uint64 {
	// empty tree
	if bst.head == nil {
		return 0
	}

	// search nodes
	currentNode := bst.head
	for {
		// found the right node
		if currentNode.key == key {
			return currentNode.value
		}

		// navigate to next node
		if key < currentNode.key {
			if currentNode.left == nil {
				// not found
				break
			}

			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				// not found
				break
			}

			currentNode = currentNode.right
		}
	}

	// not found
	return 0
}
