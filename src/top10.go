package main

import (
	"fmt"
	"log"
	"project3/src/hashmap"
)

func top10Words(wordUseHashMap *hashmap.HashMap) {
	fmt.Print("Inserting word usage into max heap...")

	// TODO: remove
	count := uint64(0)

	it := wordUseHashMap.Begin()
	for {
		// hit end, done adding
		if it == wordUseHashMap.End() {
			break
		}

		count++

		// advance iterator
		it = it.Next()
	}

	log.Println(count)
}
