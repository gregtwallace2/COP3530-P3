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
