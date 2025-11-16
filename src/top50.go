package main

import (
	"fmt"
	"project3/src/hashmap"
	"project3/src/maxheap"
	"time"

	"github.com/loov/hrtime"
)

type pair struct {
	key   string
	value uint64
}

// top50Words finds the 50 most used words and prints them using 2 different methods;
// time elapsed for various steps is also printed
func top50Words(wordUseHashMap *hashmap.HashMap) {
	fmt.Print("Inserting word usage into max heap ... ")

	var totalHeap time.Duration
	var totalSort time.Duration

	// make heap
	mh := maxheap.NewMaxHeap(27000)

	it := wordUseHashMap.Begin()
	tStart := hrtime.Now()
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
	tEnd := hrtime.Now()
	totalHeap = tEnd - tStart
	fmt.Printf("done in %d nanoseconds.\n", (tEnd - tStart).Nanoseconds())

	// Print top 50 using the max heap
	fmt.Printf("Most Used Words (from max heap):\n")
	tStart = hrtime.Now()
	for i := range 50 {
		k, v := mh.Pop()
		fmt.Printf("%d) %s: %d\n", i+1, k, v)
	}
	tEnd = hrtime.Now()
	totalHeap += (tEnd - tStart)
	fmt.Printf("Max heap took %d nanoseconds to pop the top 50.\n\n", (tEnd - tStart).Nanoseconds())

	// second option - insertion sort of an array
	fmt.Print("Inserting word usage into an array ... ")
	arr := make([]pair, wordUseHashMap.Size())

	it = wordUseHashMap.Begin()
	tStart = hrtime.Now()
	i := 0
	for {
		// hit end, done adding
		if it == wordUseHashMap.End() {
			break
		}

		// add
		k, v := it.KeyValue()
		p := pair{
			key:   k,
			value: v,
		}
		arr[i] = p

		// advance iterator
		it = it.Next()
		i++
	}
	tEnd = hrtime.Now()
	totalSort = (tEnd - tStart)
	fmt.Printf("done in %d nanoseconds.\n", (tEnd - tStart).Nanoseconds())

	// performing insertion sort
	fmt.Printf("Insertion sorting ... ")
	tStart = hrtime.Now()
	insertionSort(&arr)
	tEnd = hrtime.Now()
	totalSort += (tEnd - tStart)
	fmt.Printf("done in %d nanoseconds.\n\n", (tEnd - tStart).Nanoseconds())

	// Print top 50 using the array
	fmt.Printf("Most Used Words (from insertion sort):\n")
	tStart = hrtime.Now()
	for i := range 50 {
		fmt.Printf("%d) %s: %d\n", i+1, arr[i].key, arr[i].value)
	}
	tEnd = hrtime.Now()
	fmt.Printf("Accessing sorted array took %d nanoseconds.\n\n", (tEnd - tStart).Nanoseconds())

	// show comparison faster/slower
	if totalHeap < totalSort {
		fmt.Printf("Max heap was faster than sort (heap: %d vs sort: %d nanoseconds).\n\n", totalHeap, totalSort)
	} else {
		// should never be the case though
		fmt.Printf("Sort was faster than max heap (heap: %d vs sort: %d nanoseconds).\n\n", totalHeap, totalSort)
	}
}
