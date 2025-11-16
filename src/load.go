package main

import (
	"encoding/json"
	"fmt"
	"os"
	"project3/src/hashmap"
	"strings"
	"unicode"

	"github.com/loov/hrtime"
)

type shakespeareWork struct {
	Title    string `json:"title"`
	WorkType string `json:"type"`
	Text     string `json:"text"`
}

// fn is a function used to ignore punctuation when word finding
var fn = func(c rune) bool {
	return (unicode.IsSpace(c) ||
		c == '.' ||
		c == '!' ||
		c == '&' ||
		c == '(' ||
		c == ')' ||
		c == ',' ||
		// c == '-' ||
		c == ':' ||
		c == ';' ||
		c == '?' ||
		c == '[' ||
		c == ']' ||
		c == '_' ||
		c == '`' ||
		c == '|' ||
		c == '}')
}

// loadData loads the word list from shakespeare's works json file
func loadData() (*hashmap.HashMap, error) {
	// open file
	dataFile, err := os.Open("./resources/shakespeare.json")
	if err != nil {
		return nil, fmt.Errorf("load: failed to open dataset file shakespeare.json (%s)", err)
	}
	defer dataFile.Close()

	// parse json data
	data := []shakespeareWork{}
	decoder := json.NewDecoder(dataFile)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("load: failed to parse json data from shakespeare.json (%s)", err)
	}

	// count words by adding to hashmap
	fmt.Printf("Adding words to hash map ... ")
	tStart := hrtime.Now()
	hm := hashmap.NewHashMap(800000)
	// allWords := map[string]uint64{}

	// for each work
	for i := range data {
		// break up words
		textWords := strings.FieldsFunc(data[i].Text, fn)

		for j := range textWords {
			lowerWord := strings.ToLower(textWords[j])
			// allWords[lowerWord] += 1
			hm.Increase(lowerWord)
		}
	}
	tEnd := hrtime.Now()
	fmt.Printf("done in %d nanoseconds.\n", (tEnd - tStart).Nanoseconds())

	fmt.Printf("Total unique words loaded: %d\n", hm.Size())
	fmt.Print("(Words are case insensitive.)\n")

	return hm, nil
}
