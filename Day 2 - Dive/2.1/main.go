// Now, you need to figure out how to pilot this thing.

// It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

// forward X increases the horizontal position by X units.
// down X increases the depth by X units.
// up X decreases the depth by X units.
// Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

// The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

// forward 5
// down 5
// forward 8
// up 3
// down 8
// forward 2
// Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

// forward 5 adds 5 to your horizontal position, a total of 5.
// down 5 adds 5 to your depth, resulting in a value of 5.
// forward 8 adds 8 to your horizontal position, a total of 13.
// up 3 decreases your depth by 3, resulting in a value of 2.
// down 8 adds 8 to your depth, resulting in a value of 10.
// forward 2 adds 2 to your horizontal position, a total of 15.
// After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

// ========================================================== CODE ==========================================================
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read the input file.
	input := readInput("input.txt")

	// Calculate the final position.
	output := getSubmarinePosition(input)

	// Print the output.
	fmt.Println(output)
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
			// Print an error message.
			fmt.Println("Invalid direction:", direction)
		}
	}
	return x * y
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
