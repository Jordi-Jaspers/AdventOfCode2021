// As the submarine drops below the surface of the ocean, it automatically performs a sonar sweep of the nearby sea floor. On a small screen, the sonar sweep report (your puzzle input) appears: each line is a measurement of the sea floor depth as the sweep looks further and further away from the submarine.

// This report indicates that, scanning outward from the submarine, the sonar sweep found depths of 199, 200, 208, 210, and so on.

// The first order of business is to figure out how quickly the depth increases, just so you know what you're dealing with - you never know if the keys will get carried into deeper water by an ocean current or a fish or something.

// To do this, count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.) In the example above, the changes are as follows:

// 199 (N/A - no previous measurement)
// 200 (increased)
// 208 (increased)
// 210 (increased)
// 200 (decreased)
// 207 (increased)
// 240 (increased)
// 269 (increased)
// 260 (decreased)
// 263 (increased)
// In this example, there are 7 measurements that are larger than the previous measurement.

// How many measurements are larger than the previous measurement?

// ================================================= CODE =================================================

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Read input
	input := readInput("input.txt")

	// Calculate the output
	output := increasingCounter(convertStringToInt(input))

	// Printing the output to the console
	fmt.Println(output)
}

// Check if the next number is larger than the previous one. Add to the counter.
func increasingCounter(input []int) int {
	counter := 0

	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			counter++
		}
	}

	return counter
}

// Convert a list strings to a list of integers.
func convertStringToInt(input []string) []int {
	var output []int
	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, i)
	}
	return output
}

// Migrate a text file to a slice of of strings.
func readInput(fileName string) []string {
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
