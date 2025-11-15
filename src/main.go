package main

import (
	"log"
	"os"
)

func main() {
	// load the word data into the hashmap
	_, err := loadData()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

}
