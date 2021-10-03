package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/dananjayavr/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	var sortedNames []string
	for name, _ := range counts {
		sortedNames = append(sortedNames, name)
	}

	sort.Strings(sortedNames)

	for _, name := range sortedNames {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}
}
