package main

import (
	"bufio"
	"fmt"
	"os"
	"project3/src/binarysearchtree"
	"project3/src/hashmap"
	"strings"
	"time"
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

	var totalHashMap time.Duration
	var totalBST time.Duration

	// search word method #1
	fmt.Print("\nGetting word usage from hash map ... ")
	tStart := time.Now()
	count := wordUseHashMap.GetValue(wordInput)
	tEnd := time.Now()
	totalHashMap = tEnd.Sub(tStart)
	fmt.Printf("done in %d nanoseconds.\n", tEnd.Sub(tStart).Nanoseconds())
	fmt.Printf("%s used %d times.\n", wordInput, count)

	// search word method #2
	fmt.Print("\nInserting word usage into binary search tree ... ")
	bst := binarysearchtree.NewBinarySearchTree()

	it := wordUseHashMap.Begin()
	tStart = time.Now()
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
	tEnd = time.Now()
	fmt.Printf("done in %d nanoseconds.\n", tEnd.Sub(tStart).Nanoseconds())

	// search method #2
	fmt.Print("\nGetting word usage from binary search tree ... ")
	tStart = time.Now()
	count = bst.Search(wordInput)
	tEnd = time.Now()
	totalBST = tEnd.Sub(tStart)
	fmt.Printf("done in %d nanoseconds.\n", tEnd.Sub(tStart).Nanoseconds())
	fmt.Printf("%s used %d times.\n", wordInput, count)

	fmt.Print("\n")

	// Note: Time adding to the BST is deliberately excluded from below to compare search times only

	// show comparison faster/slower
	if totalHashMap < totalBST {
		fmt.Printf("Hash map was faster searching than binary search tree (map: %d vs bst: %d nanoseconds).\n\n", totalHashMap, totalBST)
	} else {
		fmt.Printf("Binary search tree was faster searchingthan hash map (map: %d vs bst: %d nanoseconds).\n\n", totalHashMap, totalBST)
	}
}
