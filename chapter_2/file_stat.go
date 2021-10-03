package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Size())
}

