// Start by comparing the first and second three-measurement windows.
// The measurements in the first window are marked A (199, 200, 208);
// their sum is 199 + 200 + 208 = 607. The second window is marked B (200, 208, 210); its sum is 618.
// The sum of measurements in the second window is larger than the sum of the first, so this first comparison increased.

// Your goal now is to count the number of times the sum of measurements in this sliding window increases from the previous sum.
// So, compare A with B, then compare B with C, then C with D, and so on. Stop when there aren't enough measurements left to create a new three-measurement sum.

// In the above example, the sum of each three-measurement window is as follows:

// A: 607 (N/A - no previous sum)
// B: 618 (increased)
// C: 618 (no change)
// D: 617 (decreased)
// E: 647 (increased)
// F: 716 (increased)
// G: 769 (increased)
// H: 792 (increased)

// In this example, there are 5 sums that are larger than the previous sum.

// ================================================= CODE =================================================
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read the input file.
	content := readInput("input.txt")

	// Convert the input to integers.
	input := convertStringToInt(content)

	// Remove the noise.
	input = removeNoise(input)

	// Count the number of times the sum of the three numbers is larger than the previous sum.
	log.Println(increasingCounter(input))
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

// add 3 elements to the list and move the window forward.
func removeNoise(input []int) []int {
	length := len(input) - 1
	output := make([]int, 0)

	for i := 0; i < length; i++ {
		if i+1 < length || i+2 < length {
			output = append(output, input[i]+input[i+1]+input[i+2])
		}
	}
	return output
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
