package main

import (
	"fmt"
	"log"

	"github.com/kartik-maguwala/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
