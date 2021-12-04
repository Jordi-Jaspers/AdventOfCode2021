package main

import (
	"log"

	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
)

func main() {

	// Read input
	input := utils.ReadInput("../input.txt")

	// Calculate the output
	output := increasingCounter(utils.ConvertStringToInt(input))

	// Printing the output to the console
	log.Println(output)
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
