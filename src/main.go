package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("\nShakespeare’s Undying Love (of Words)\n\n")
	fmt.Printf("Starting...\n")

	// load the word data into the hashmap
	wordUseHashMap, err := loadData()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nMain Menu:\n")
		fmt.Print("1) Shakespeare's Top 20 Most Used Words\n")
		fmt.Print("0) Exit\n\n")

		// user input
		fmt.Print("Enter digit selection: ")
		menuInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("User Input Error: %s", err)
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
			top20Words(wordUseHashMap)
		} else {
			fmt.Print("Invalid selection. Try again.\n")
		}
	}

	fmt.Print("Goodbye!\n\n")
}
