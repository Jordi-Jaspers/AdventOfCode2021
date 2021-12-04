package main

import (
	"log"
	"math"

	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
)

func main() {

	// read the input file
	input := utils.ReadInput("../input.txt")

	// convert the input to a slice of digits
	inputInt := utils.ConvertStringToDigits(input)

	// calculate the power consumption
	binary := getFinalBit(inputInt)

	// calculate the power consumption
	output := getPowerConsmption(binary)

	// print the output
	log.Println(output)
}

// Get the power consumption by multiplying the binary numbers.
func getPowerConsmption(binary []int) int {
	length := len(binary) - 1
	gammaRate := 0
	epsilonRate := 0

	// Convert binary to decimal
	for i := length; i != -1; i-- {
		if binary[i] == 1 {
			gammaRate += int(math.Pow(2, float64(length-i)))
		} else {
			epsilonRate += int(math.Pow(2, float64(length-i)))
		}
	}
	return gammaRate * epsilonRate
}

// Get the final bit by adding all the bits in the input and count the number of 1s.
func getFinalBit(input [][]int) []int {
	rows := len(input)
	cols := len(input[0])
	output := make([]int, len(input[0]))

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
