package main

import (
	"fmt"
	"project3/src/hashmap"
	"project3/src/maxheap"

	"github.com/loov/hrtime"
)

type pair struct {
	key   string
	value uint64
}

// topWords finds the `topX` most used words and prints them using 2 different methods;
func topWords(wordUseHashMap *hashmap.HashMap, topX uint64) {
	fmt.Print("Inserting word usage into max heap ... ")

	// make heap
	timeStart := hrtime.Now()
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
	timeEnd := hrtime.Now()
	timeHeap := timeEnd - timeStart
	fmt.Printf("done in %d nanoseconds.\n", (timeEnd - timeStart).Nanoseconds())

	// Print top topX using the max heap
	fmt.Printf("Most Used Words (from max heap):\n")
	timeStart = hrtime.Now()
	for i := range topX {
		k, v := mh.Pop()
		fmt.Printf("%d) %s: %d\n", i+1, k, v)
	}
	timeEnd = hrtime.Now()
	timeHeap += (timeEnd - timeStart)
	fmt.Printf("Max heap took %d nanoseconds to pop the top %d.\n\n", (timeEnd - timeStart).Nanoseconds(), topX)

	// second option - insertion sort of an array
	fmt.Print("Inserting word usage into an array ... ")

	timeStart = hrtime.Now()
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
	timeEnd = hrtime.Now()
	timeSort := (timeEnd - timeStart)
	fmt.Printf("done in %d nanoseconds.\n", (timeEnd - timeStart).Nanoseconds())

	// performing insertion sort
	fmt.Printf("Insertion sorting ... ")
	timeStart = hrtime.Now()
	insertionSort(&arr)
	timeEnd = hrtime.Now()
	timeSort += (timeEnd - timeStart)
	fmt.Printf("done in %d nanoseconds.\n\n", (timeEnd - timeStart).Nanoseconds())

	// Print top topX using the array
	fmt.Printf("Most Used Words (from insertion sort):\n")
	timeStart = hrtime.Now()
	for i := range topX {
		fmt.Printf("%d) %s: %d\n", i+1, arr[i].key, arr[i].value)
	}
	timeEnd = hrtime.Now()
	timeSort += (timeEnd - timeStart)
	fmt.Printf("Accessing sorted array took %d nanoseconds.\n\n", (timeEnd - timeStart).Nanoseconds())

	// show comparison faster/slower
	if timeHeap == timeSort {
		fmt.Printf("Max heap was roughly comparable to insertion sort (%d nanoseconds).\n\n", timeHeap)
	} else if timeHeap < timeSort {
		fmt.Printf("Max heap was %.2f times faster than insertion sort (heap: %d vs sort: %d nanoseconds).\n\n",
			float64(timeSort)/float64(timeHeap), timeHeap, timeSort)
	} else {
		fmt.Printf("Insertion sort was %.2f times faster than max heap (heap: %d vs sort: %d nanoseconds).\n\n",
			float64(timeHeap)/float64(timeSort), timeHeap, timeSort)
	}
}
