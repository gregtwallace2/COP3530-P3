package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// hashMapCap sets the hash map's bucket size; can tweak this const to
// evaluate performance impact
const hashMapCap = 800000

// main app & app loop
func main() {
	fmt.Print("\nShakespeare’s Undying Love (of Words)\n\n")

	// load the word data into the hashmap
	wordUseHashMap, err := loadData()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nMain Menu:\n")
		fmt.Print("1) Top 50 Most Used Words\n")
		fmt.Print("2) All Words In Order\n")
		fmt.Print("3) Search Word for Usage Count\n")
		fmt.Print("0) Exit\n\n")

		// user input
		fmt.Print("Enter digit selection: ")
		menuInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("User Input Error: %s\n", err)
			continue
		}
		// remove end line chars
		menuInput = strings.Trim(menuInput, string(byte(10))) // LF
		menuInput = strings.Trim(menuInput, string(byte(13))) // CR

		fmt.Print("\n")

		// do user request
		if menuInput == "0" {
			break
		} else if menuInput == "1" {
			topWords(wordUseHashMap, 50)
		} else if menuInput == "2" {
			topWords(wordUseHashMap, wordUseHashMap.Size())
		} else if menuInput == "3" {
			searchWord(wordUseHashMap)
		} else {
			fmt.Print("Invalid selection. Try again.\n")
		}
	}

	fmt.Print("Goodbye!\n\n")
}
