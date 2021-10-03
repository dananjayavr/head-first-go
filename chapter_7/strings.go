// Package datafile allows reading data samples from files.
package main

import (
	"bufio"
	"os"
)

// GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var strings []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string := scanner.Text()
		strings = append(strings, string)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return strings, err
}
