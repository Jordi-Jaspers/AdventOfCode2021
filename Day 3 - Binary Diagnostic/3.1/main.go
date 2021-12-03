// The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

// The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

// You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

// Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

// 00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010
// Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

// The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

// The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

// So, the gamma rate is the binary number 10110, or 22 in decimal.

// The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.
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
