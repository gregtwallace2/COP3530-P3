package main

import (
	"fmt"
	"project3/src/hashmap"
	"project3/src/maxheap"
	"time"

	"github.com/loov/hrtime"
)

// topAllWords prints all words and their usage using 2 different methods;
func topAllWords(wordUseHashMap *hashmap.HashMap) {
	fmt.Print("Inserting word usage into max heap ... ")

	var totalHeap time.Duration
	var totalSort time.Duration

	// make heap
	tStart := hrtime.Now()
	mh := maxheap.NewMaxHeap(27000)
	it := wordUseHashMap.Begin()
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
	fmt.Printf("done in %d nanoseconds.\n", (tEnd - tStart).Nanoseconds())

	// Print top 50 using the max heap
	fmt.Printf("All Words And Usage (from max heap):\n")
	for i := range wordUseHashMap.Size() {
		k, v := mh.Pop()
		fmt.Printf("%d) %s: %d\n", i+1, k, v)
	}
	tEnd = hrtime.Now()
	totalHeap = (tEnd - tStart)
	fmt.Printf("Max heap took %d nanoseconds to pop all.\n\n", (tEnd - tStart).Nanoseconds())

	// second option - insertion sort of an array
	fmt.Print("Inserting word usage into an array ... ")

	tStart = hrtime.Now()
	arr := make([]pair, wordUseHashMap.Size())

	it = wordUseHashMap.Begin()
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
	fmt.Printf("done in %d nanoseconds.\n", (tEnd - tStart).Nanoseconds())

	// performing insertion sort
	fmt.Printf("Insertion sorting ... ")
	insertionSort(&arr)
	tEnd = hrtime.Now()
	totalSort = (tEnd - tStart)
	fmt.Printf("done in %d nanoseconds.\n\n", (tEnd - tStart).Nanoseconds())

	// Print top 50 using the array
	fmt.Printf("All Words And Usage (from insertion sort):\n")
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
