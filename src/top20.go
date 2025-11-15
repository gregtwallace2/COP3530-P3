package main

import (
	"fmt"
	"project3/src/hashmap"
	"project3/src/maxheap"
	"time"
)

type pair struct {
	key   string
	value uint64
}

// top20Words finds the 20 most used words and prints them using 2 different methods;
// time elapsed for various steps is also printed
func top20Words(wordUseHashMap *hashmap.HashMap) {
	fmt.Print("Inserting word usage into max heap ... ")

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
	fmt.Printf("done in %d microseconds.\n", tEnd.Sub(tStart).Microseconds())

	// Print top 20 using the max heap
	fmt.Printf("Most Used Words (from max heap):\n")
	tStart = time.Now()
	for i := range 20 {
		k, v := mh.Pop()
		fmt.Printf("%d) %s: %d\n", i+1, k, v)
	}
	tEnd = time.Now()
	fmt.Printf("Max heap took %d microseconds to pop the top 20.\n\n", tEnd.Sub(tStart).Microseconds())

	// second option - insertion sort of an array
	fmt.Print("Inserting word usage into an array ... ")
	arr := make([]pair, wordUseHashMap.Size())

	it = wordUseHashMap.Begin()
	tStart = time.Now()
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
	tEnd = time.Now()
	fmt.Printf("done in %d microseconds.\n", tEnd.Sub(tStart).Microseconds())

	// performing insertion sort
	fmt.Printf("Insertion sorting ... ")
	tStart = time.Now()
	insertionSort(&arr)
	tEnd = time.Now()
	fmt.Printf("done in %d microseconds.\n\n", tEnd.Sub(tStart).Microseconds())

	// Print top 20 using the array
	fmt.Printf("Most Used Words (from insertion sort):\n")
	tStart = time.Now()
	for i := range 20 {
		fmt.Printf("%d) %s: %d\n", i+1, arr[i].key, arr[i].value)
	}
	tEnd = time.Now()
	fmt.Printf("Accessing sorted array took %d microseconds.\n\n", tEnd.Sub(tStart).Microseconds())
}
