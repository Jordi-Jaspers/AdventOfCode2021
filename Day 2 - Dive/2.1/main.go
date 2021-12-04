package main

import (
	"log"
	"regexp"
	"strconv"

	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
)

func main() {
	// Read the input file.
	input := utils.ReadInput("../input.txt")

	// Calculate the final position.
	output := getSubmarinePosition(input)

	// Print the output.
	log.Println(output)
}

// Calculate the final position by multiplying the horizontal position by the depth.
func getSubmarinePosition(input []string) int {
	// Initialize the horizontal position (X-axis).
	x := 0

	// Initialize the vertical position (Y-axis).
	y := 0

	// Iterate over the input.
	for _, instruction := range input {
		// Get the direction by getting the firt letter of the instruction.
		direction := instruction[0:1]

		// Get the distance by parsing the integer after the space using regex.
		re := regexp.MustCompile("[0-9]+")
		value := re.FindString(instruction)

		// Convert the distance to an integer.
		distance, err := strconv.Atoi(value)

		// Check for errors.
		if err != nil {
			log.Fatal(err)
		}

		// Check the direction.
		switch direction {
		case "f":
			// Increase the horizontal position by the distance.
			x += distance
		case "d":
			// Increase the vertical position by the distance.
			y += distance
		case "u":
			// Decrease the vertical position by the distance.
			y -= distance
		default:
			log.Fatal("Invalid direction:", direction)
		}
	}
	return x * y
}
