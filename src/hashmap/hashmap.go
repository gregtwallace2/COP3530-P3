package hashmap

// WARNING: HashMap is not safe for concurrent access, modification, or iteration.

const (
	fnvOffsetBasis = uint64(14695981039346656037)
	fnvPrime       = uint64(1099511628211)
)

// hashFunc is the hashing function for the hash map (FNV-1a hash)
func hashFunc(key string, capacity uint64) uint64 {
	h := fnvOffsetBasis

	for _, byt := range []byte(key) {
		h = h ^ uint64(byt)
		h = h * fnvPrime
	}

	return h % capacity
}

// node is used to store data in the hash map
type node struct {
	key   string
	value uint64

	next *node
}

// hashMap is the actual struct for the hashMap
type HashMap struct {
	data []*node

	size     uint64
	capacity uint64
}

// NewHashMap creates a hashmap with the specified capacity
func NewHashMap(capacity uint64) *HashMap {
	hm := &HashMap{}

	hm.data = make([]*node, capacity)
	hm.capacity = capacity
	hm.size = 0

	return hm
}

// Increase increments the value associated with the key by 1; if
// the key isn't in the hashmap, it creates the key and sets the value
// to one.
func (hm *HashMap) Increase(key string) {
	indx := hashFunc(key, hm.capacity)

	// if indx exists, check if the correct node exists
	var n *node
	if hm.data[indx] != nil {
		n = hm.data[indx]
		for {
			// if this is the right node, increment & done
			if n.key == key {
				n.value++
				return
			}

			// end of nodes? break, not found
			if n.next == nil {
				break
			}

			// not at end, check next node
			n = n.next
		}
	}

	// will be adding a node
	hm.size++
	newNode := &node{
		key:   key,
		value: 1,
	}

	// if indx existed, but ran out of nodes, add new node
	if n != nil {
		n.next = newNode
	} else {
		// indx didn't exist
		hm.data[indx] = newNode
	}
}

// GetValue returns the value associated with key. If the key does not exist,
// 0 is returned.
func (hm *HashMap) GetValue(key string) uint64 {
	indx := hashFunc(key, hm.capacity)

	// check for indx
	if hm.data[indx] != nil {
		// check for node
		var n *node
		n = hm.data[indx]
		for {
			// if this is the right node, return its value
			if n.key == key {
				return n.value
			}

			// end of nodes? break, not found
			if n.next == nil {
				break
			}

			// not at end, check next node
			n = n.next
		}
	}

	// not found
	return 0
}

func (hm *HashMap) Size() uint64 {
	return hm.size
}
