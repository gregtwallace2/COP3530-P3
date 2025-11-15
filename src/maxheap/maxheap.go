package maxheap

// node is used to store data in the max heap array
type node struct {
	key    string
	weight uint64

	next *node
}

// MaxHeap is the struct for the MaxHeap object
type MaxHeap struct {
	data []node

	size     uint64
	capacity uint64
}

// NewMaxHeap creates a MaxHeap the specified capacity
func NewMaxHeap(capacity uint64) *MaxHeap {
	mh := &MaxHeap{}

	mh.data = make([]node, capacity)
	mh.capacity = capacity
	mh.size = 0

	return mh
}

// Insert inserts the key with specified weight into the max heap
func (mh *MaxHeap) Insert(key string, weight uint64) {
	// insert at end
	mh.data[mh.size] = node{
		key:    key,
		weight: weight,
	}
	mh.size++

	// shift up the tree as appropriate
	childIndx := mh.size - 1
	parentIndx := (childIndx - 1) / 2

	for {
		// if child is root, we're done
		if childIndx == 0 {
			break
		}

		// compare child to parent and if child is bigger, swap them
		if mh.data[childIndx].weight > mh.data[parentIndx].weight {
			temp := node{
				key:    mh.data[childIndx].key,
				weight: mh.data[childIndx].weight,
			}

			mh.data[childIndx].key = mh.data[parentIndx].key
			mh.data[childIndx].weight = mh.data[parentIndx].weight

			mh.data[parentIndx].key = temp.key
			mh.data[parentIndx].weight = temp.weight
		}

		// move up queue
		childIndx = parentIndx
		parentIndx = (childIndx - 1) / 2
	}
}

// Pop removes and returns the max key: value from the maxHeap
func (mh *MaxHeap) Pop() (string, uint64) {
	// get value to return
	retNode := mh.data[0]

	// remove value and place last value at root
	mh.data[0] = mh.data[mh.size-1]
	mh.size--

	// children *2 + 1 ; *2 + 2
	parentIndx := uint64(0)
	for {
		child1Indx := parentIndx*2 + 1
		child2Indx := parentIndx*2 + 2

		// if first (smaller index child) is out of bounds, done (no children)
		if child1Indx > mh.size {
			break
		}

		// if only one child, check if it is bigger than parent, if so, swap
		if child2Indx > mh.size {
			if mh.data[child1Indx].weight > mh.data[parentIndx].weight {
				tempKey := mh.data[child1Indx].key
				tempWeight := mh.data[child1Indx].weight

				mh.data[child1Indx].key = mh.data[parentIndx].key
				mh.data[child1Indx].weight = mh.data[parentIndx].weight

				mh.data[parentIndx].key = tempKey
				mh.data[parentIndx].weight = tempWeight

				// update parentIndx for next iteration
				parentIndx = child1Indx
			} else {
				// not bigger, done
				break
			}
		} else {
			// 2 children (compare using whichever child is bigger)
			compChildIndx := child1Indx
			if mh.data[child2Indx].weight > mh.data[child1Indx].weight {
				compChildIndx = child2Indx
			}

			if mh.data[compChildIndx].weight > mh.data[parentIndx].weight {
				tempKey := mh.data[compChildIndx].key
				tempWeight := mh.data[compChildIndx].weight

				mh.data[compChildIndx].key = mh.data[parentIndx].key
				mh.data[compChildIndx].weight = mh.data[parentIndx].weight

				mh.data[parentIndx].key = tempKey
				mh.data[parentIndx].weight = tempWeight

				// update parentIndx for next iteration
				parentIndx = compChildIndx
			} else {
				// not bigger, done
				break
			}
		}
	}

	// done
	return retNode.key, retNode.weight
}
