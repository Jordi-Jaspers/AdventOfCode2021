// Based on your calculations, the planned course doesn't seem to make any sense. You find the submarine manual and discover that the process is actually slightly more complicated.

// In addition to horizontal position and depth, you'll also need to track a third value, aim, which also starts at 0. The commands also mean something entirely different than you first thought:

// down X increases your aim by X units.
// up X decreases your aim by X units.
// forward X does two things:
// It increases your horizontal position by X units.
// It increases your depth by your aim multiplied by X.
// Again note that since you're on a submarine, down and up do the opposite of what you might expect: "down" means aiming in the positive direction.

// Now, the above example does something different:

// forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
// down 5 adds 5 to your aim, resulting in a value of 5.
// forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
// up 3 decreases your aim by 3, resulting in a value of 2.
// down 8 adds 8 to your aim, resulting in a value of 10.
// forward 2 adds 2 to your horizontal position, a total of 15. Because your aim is 10, your depth increases by 2*10=20 to a total of 60.
// After following these new instructions, you would have a horizontal position of 15 and a depth of 60. (Multiplying these produces 900.)

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

	// Initialize the aim.
	aim := 0

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
			// Increase the depth by the aim multiplied by the distance.
			y += aim * distance
		case "d":
			// Increase the aim by the distance.
			aim += distance
		case "u":
			// Decrease the aim by the distance.
			aim -= distance
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
