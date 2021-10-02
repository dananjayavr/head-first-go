package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func reportPanic() {
	fmt.Println("Report panic called")
	p := recover()
	if p == nil {
		return // do nothing
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func main() {
	defer reportPanic()
	scanDirectory(os.Args[1])
	//scanDirectory("go")
}

func scanDirectory(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}
