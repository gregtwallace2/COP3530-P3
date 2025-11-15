package main

import (
	"fmt"
	"project3/src/hashmap"
	"project3/src/maxheap"
	"time"
)

func top10Words(wordUseHashMap *hashmap.HashMap) {
	fmt.Print("Inserting word usage into max heap ...")

	// make heap
	mh := maxheap.NewMaxHeap(27000)

	it := wordUseHashMap.Begin()
	tStart := time.Now()
	for {
		// hit end, done adding
		if it == wordUseHashMap.End() {
			break
		}

		// add to max heap
		mh.Insert(it.KeyValue())

		// advance iterator
		it = it.Next()
	}
	tEnd := time.Now()
	fmt.Printf(" done in %d microseconds.\n", tEnd.Sub(tStart).Microseconds())

	// Print top 10 using the max heap
	fmt.Printf("Most Used Words (from max heap):\n")
	tStart = time.Now()
	for i := range 10 {
		k, v := mh.Pop()
		fmt.Printf("%d) %s: %d\n", i+1, k, v)
	}
	tEnd = time.Now()
	fmt.Printf("Max heap took %d microseconds to pop the top 10.\n", tEnd.Sub(tStart).Microseconds())

	// TODO: remove
	count := uint64(0)
	//
	// TODO: Remove
	fmt.Printf("TODO: REMOVE: iterator count: %d\n", count)
	//

	// TODO: another data structure for comparison
}
