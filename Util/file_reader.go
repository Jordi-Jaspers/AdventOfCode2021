// Package utils Exporting functions from another package require that the identifier of the function begins with a capital letter.
package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadInput Migrate a text file to a slice of strings.
func ReadInput(fileName string) []string {
	input := make([]string, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
