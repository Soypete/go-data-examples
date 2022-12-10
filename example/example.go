package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// list all the files in a directory
	files, err := os.ReadDir("./data/player")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	// import the player stats file
	f, err := os.Open("data/player/player_stats.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)

	// data imported here is of type string[][]
	// when data is imported in an ambiguous type it is typical to use
	// generics, and empty interface{}, or a string.
	// when when reading from a file it is common to use strings.
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range data {
		// first line is column names
		if i == 0 {
			fmt.Printf("There are %d columns in the file %s\n here are the csv column names: %s \n", len(line), "player_stats.csv", line)
			continue
		}
		if i > 9 {
			break
		}
		fmt.Println(line)
	}

}
