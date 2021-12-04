package main

import (
	"log"

	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
)

func main() {
	// Read the input file.
	content := utils.ReadInput("../input.txt")

	// Convert the input to integers.
	input := utils.ConvertStringToInt(content)

	// Remove the noise.
	input = removeNoise(input)

	// Count the number of times the sum of the three numbers is larger than the previous sum.
	log.Println(increasingCounter(input))
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
