// Next, you should verify the life support rating, which can be determined by multiplying the oxygen generator rating by the CO2 scrubber rating.

// Both the oxygen generator rating and the CO2 scrubber rating are values that can be found in your diagnostic report - finding them is the tricky part. Both values are located using a similar process that involves filtering out values until only one remains. Before searching for either rating value, start with the full list of binary numbers from your diagnostic report and consider just the first bit of those numbers. Then:

// Keep only numbers selected by the bit criteria for the type of rating value for which you are searching. Discard numbers which do not match the bit criteria.
// If you only have one number left, stop; this is the rating value for which you are searching.
// Otherwise, repeat the process, considering the next bit to the right.
// The bit criteria depends on which type of rating value you want to find:

// To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 1 in the position being considered.
// To find CO2 scrubber rating, determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 0 in the position being considered.
// For example, to determine the oxygen generator rating value using the same example diagnostic report from above:

// Start with all 12 numbers and consider only the first bit of each number. There are more 1 bits (7) than 0 bits (5), so keep only the 7 numbers with a 1 in the first position: 11110, 10110, 10111, 10101, 11100, 10000, and 11001.
// Then, consider the second bit of the 7 remaining numbers: there are more 0 bits (4) than 1 bits (3), so keep only the 4 numbers with a 0 in the second position: 10110, 10111, 10101, and 10000.
// In the third position, three of the four numbers have a 1, so keep those three: 10110, 10111, and 10101.
// In the fourth position, two of the three numbers have a 1, so keep those two: 10110 and 10111.
// In the fifth position, there are an equal number of 0 bits and 1 bits (one each). So, to find the oxygen generator rating, keep the number with a 1 in that position: 10111.
// As there is only one number left, stop; the oxygen generator rating is 10111, or 23 in decimal.
// Then, to determine the CO2 scrubber rating value from the same example above:

// Start again with all 12 numbers and consider only the first bit of each number. There are fewer 0 bits (5) than 1 bits (7), so keep only the 5 numbers with a 0 in the first position: 00100, 01111, 00111, 00010, and 01010.
// Then, consider the second bit of the 5 remaining numbers: there are fewer 1 bits (2) than 0 bits (3), so keep only the 2 numbers with a 1 in the second position: 01111 and 01010.
// In the third position, there are an equal number of 0 bits and 1 bits (one each). So, to find the CO2 scrubber rating, keep the number with a 0 in that position: 01010.
// As there is only one number left, stop; the CO2 scrubber rating is 01010, or 10 in decimal.
// Finally, to find the life support rating, multiply the oxygen generator rating (23) by the CO2 scrubber rating (10) to get 230.he epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.
// ========================================================== CODE ==========================================================
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// read the input file
	input := readInput("input.txt")

	// convert the input to a slice of digits
	inputInt := convertStringToDigits(input)

	// calculate the power consumption
	binary := getFinalBit(inputInt)

	// calculate the power consumption
	output := getPowerConsmption(binary)

	// print the output
	fmt.Println(output)
}

func getPowerConsmption(binary []int) int {
	length := len(binary) - 1
	gammaRate := 0
	epsilonRate := 0

	// Convert binary to decimal: 1001
	for i := length; i != -1; i-- {
		if binary[i] == 1 {
			gammaRate += int(math.Pow(2, float64(length-i)))
		} else {
			epsilonRate += int(math.Pow(2, float64(length-i)))
		}
	}
	return gammaRate * epsilonRate
}

// get the final bit by adding all the bits in the input and count the number of 1s.
func getFinalBit(input [][]int) []int {
	rows := len(input)
	cols := len(input[0])
	fmt.Println(rows, cols)
	output := make([]int, len(input[0]))

	// iterate over the rows -->  [1 0 1 0 1 0 1]
	for i := 0; i < cols; i++ {
		counter := 0
		for j := 0; j < rows; j++ {
			if input[j][i] == 1 {
				counter++
			}
		}

		if counter >= len(input)/2 {
			output[i] = 1
		} else {
			output[i] = 0
		}
	}
	return output
}

// by splitting the string into a slice of characters, we can convert the string to a slice of digits.
func convertStringToDigits(input []string) [][]int {
	var output [][]int

	for _, line := range input {

		row := make([]int, 0)
		for _, digit := range line {

			// convert rune to int
			digitInt, err := strconv.Atoi(string(digit))

			// check for errors
			if err != nil {
				log.Fatal(err)
			}

			// append digit to output
			row = append(row, digitInt)
		}
		output = append(output, row)
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
