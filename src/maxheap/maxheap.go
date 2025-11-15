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
