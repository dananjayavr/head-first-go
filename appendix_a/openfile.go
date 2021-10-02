package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read(name string) {
	file, err := os.OpenFile(name, os.O_RDONLY, os.FileMode(0600))
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())
}

func write(name string)  {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("amazing!\n"))
	check(err)
	err = file.Close()
	check(err)
}

func check(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	//read("aardvark.txt")
	write("aardvark.txt")
	read("aardvark.txt")
}


