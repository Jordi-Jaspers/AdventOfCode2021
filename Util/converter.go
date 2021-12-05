package utils

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// ConvertStringToInt Convert a list strings to a list of integers.
func ConvertStringToInt(input []string) []int {
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

// ConvertStringToDigits Convert the string to a slice of digits by splitting the string into a slice of characters.
func ConvertStringToDigits(input []string) [][]int {
	var output [][]int

	for _, line := range input {
		row := make([]int, 0)
		for _, digit := range line {
			digitInt, err := strconv.Atoi(string(digit))

			if err != nil {
				log.Fatal(err)
			}

			row = append(row, digitInt)
		}
		output = append(output, row)
	}
	return output
}

// SplitDigitsFromSeperatedString ConvertDigitsToString Convert the digits to a string.
func SplitDigitsFromSeperatedString(digitString string) []int {
	numbers := make([]int, 0)
	s := strings.FieldsFunc(digitString, split)
	for _, digit := range s {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, digitInt)
	}
	return numbers
}

func split(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == ',' || r == '-' || r == '>'
}

// ConvertBinaryToDecimal Convert Binary to Decimal.
func ConvertBinaryToDecimal(binary []int) int {
	length := len(binary) - 1
	output := 0

	// Convert binary to decimal
	for i := length; i != -1; i-- {
		if binary[i] == 1 {
			output += int(math.Pow(2, float64(length-i)))
		}
	}

	return output
}
