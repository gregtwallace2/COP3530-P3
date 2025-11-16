package main

import (
	"bufio"
	"fmt"
	"os"
	"project3/src/binarysearchtree"
	"project3/src/hashmap"
	"strings"

	"github.com/loov/hrtime"
)

// searchWord allows the user to specify a word to search for and then outputs
// the number of times Shakespeare used that word
func searchWord(wordUseHashMap *hashmap.HashMap) {
	// user input
	fmt.Print("\nWhat word do you want to search (case insensitive)? ")

	reader := bufio.NewReader(os.Stdin)
	wordInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("User Word Input Error: %s\n", err)
		return
	}
	// remove end line chars
	wordInput = strings.Trim(wordInput, string(byte(10))) // LF
	wordInput = strings.Trim(wordInput, string(byte(13))) // CR

	// make word lowercase
	wordInput = strings.ToLower(wordInput)

	// search word method #1
	fmt.Print("\nGetting word usage from hash map ... ")
	timeStart := hrtime.Now()
	wordCount := wordUseHashMap.GetValue(wordInput)
	timeEnd := hrtime.Now()
	timeHashMap := timeEnd - timeStart
	fmt.Printf("done in %d nanoseconds.\n", (timeEnd - timeStart).Nanoseconds())
	fmt.Printf("%s used %d times.\n", wordInput, wordCount)

	// search word method #2
	fmt.Print("\nInserting word usage into binary search tree ... ")

	timeStart = hrtime.Now()
	bst := binarysearchtree.NewBinarySearchTree()

	it := wordUseHashMap.Begin()

	for {
		// hit end, done adding
		if it == wordUseHashMap.End() {
			break
		}

		// add to max heap
		bst.Insert(it.KeyValue())

		// advance iterator
		it = it.Next()
	}
	timeEnd = hrtime.Now()
	fmt.Printf("done in %d nanoseconds.\n", (timeEnd - timeStart).Nanoseconds())

	// search method #2
	fmt.Print("\nGetting word usage from binary search tree ... ")
	timeStart = hrtime.Now()
	wordCount = bst.Search(wordInput)
	timeEnd = hrtime.Now()
	timeBST := timeEnd - timeStart
	fmt.Printf("done in %d nanoseconds.\n", (timeEnd - timeStart).Nanoseconds())
	fmt.Printf("%s used %d times.\n", wordInput, wordCount)

	fmt.Print("\n")

	// show comparison faster/slower
	if timeHashMap == timeBST {
		fmt.Printf("Hash map search was roughly comparable to binary tree search (%d nanoseconds).\n\n", timeHashMap)
	} else if timeHashMap < timeBST {
		fmt.Printf("Hash map search was %.2f times faster than binary tree search (map: %d vs bst: %d nanoseconds).\n\n",
			float64(timeBST)/float64(timeHashMap), timeHashMap, timeBST)
	} else {
		fmt.Printf("Binary tree search search was %.2f times faster than hash map search (map: %d vs bst: %d nanoseconds).\n\n",
			float64(timeHashMap)/float64(timeBST), timeHashMap, timeBST)
	}
}
