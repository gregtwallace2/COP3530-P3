package hashmap

// Iterator allows iteration through the HashMap.
type Iterator struct {
	thisNode *node
	parentHm *HashMap
}

// Begin returns the first Iterator of the HashMap
func (hm *HashMap) Begin() *Iterator {
	// shorcut for empty HM
	if hm.size == 0 {
		return nil
	}

	// start at the beginning of the map
	node := hm.data[0]

	// if no node there, move forward until we find one
	if node == nil {
		indx := uint64(1)
		for {
			node = hm.data[indx]

			// if this indx has a node, break done
			if node != nil {
				break
			}

			// still no node, keep advancing until find one or run out of nodes
			indx++

			// out of bounds and still nothing, break (it is still nil, which is fine)
			if indx == hm.capacity {
				break
			}
		}
	}

	// still no node?
	if node == nil {
		return nil
	}

	// create iterator to return
	it := &Iterator{
		thisNode: node,
		parentHm: hm,
	}

	return it
}

// End returns a nil Iterator which can be used for iteration
func (hm *HashMap) End() *Iterator { return nil }

// Next returns the next node in the HashMap or nil if there are no more nodes
func (it *Iterator) Next() *Iterator {
	// if node has a next, that's the next node
	if it.thisNode.next != nil {
		return &Iterator{
			thisNode: it.thisNode.next,
			parentHm: it.parentHm,
		}
	}

	// if no next, search forward for next node
	indx := hashFunc(it.thisNode.key, it.parentHm.capacity) + 1 // add one for next index (already confirmed this index is done)

	for {
		// out of bounds, there is no next
		if indx >= it.parentHm.capacity {
			break
		}

		// check this indx
		if it.parentHm.data[indx] != nil {
			return &Iterator{
				thisNode: it.parentHm.data[indx],
				parentHm: it.parentHm,
			}
		}

		indx++
	}

	// never found, exhausted
	return nil
}

// KeyValue returns the key: value pair for the current iterator node
func (it *Iterator) KeyValue() (key string, value uint64) {
	return it.thisNode.key, it.thisNode.value
}
