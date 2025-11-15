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
	// TODO: remove
	count := uint64(0)
	//

	it := wordUseHashMap.Begin()
	tStart := time.Now()
	for {
		// hit end, done adding
		if it == wordUseHashMap.End() {
			break
		}

		// add to max heap
		mh.Insert(it.KeyValue())

		count++

		// advance iterator
		it = it.Next()
	}
	tEnd := time.Now()
	fmt.Printf(" done in %d microseconds.\n", tEnd.Sub(tStart).Microseconds())

	// TODO: Remove
	fmt.Printf("TODO: REMOVE: iterator count: %d\n", count)
	//

	// TODO: another data structure for comparison
}
