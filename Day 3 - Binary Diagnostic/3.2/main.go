package main

import (
	"log"

	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
)

func main() {

	// Read the input file
	log.Println("Reading input file...")
	input := utils.ReadInput("../input.txt")

	// Convert the input to a slice of digits
	log.Println("Converting input to digits...")
	inputInt := utils.ConvertStringToDigits(input)

	// Calculate the oxygen rating
	log.Println("Calculating oxygen rating...")
	oxygenRating := getRating(inputInt, true)
	log.Println("Oxygen rating:", oxygenRating)

	// Calculate the CO2 scrubber rating
	log.Println("Calculating CO2 scrubber rating...")
	scrubberRating := getRating(inputInt, false)
	log.Println("CO2 scrubber rating:", scrubberRating)

	// Calculate the life support rating
	lifeSupportRating := utils.ConvertBinaryToDecimal(oxygenRating) * utils.ConvertBinaryToDecimal(scrubberRating)

	// print the output
	log.Printf("The life support rating is: %d\n", lifeSupportRating)
}

// Get the rating for the oxygen or CO2 scrubber depending on the state of the switch.
func getRating(input [][]int, isOxygenRating bool) []int {
	columns := len(input[0])
	for i := 0; i < columns; i++ {
		// Initialize for a new itteration.
		listOfZeros := make([][]int, 0)
		listOfOnes := make([][]int, 0)

		rows := len(input)
		for j := 0; j < rows; j++ {
			if input[j][i] == 1 {
				listOfOnes = append(listOfOnes, input[j])
			} else {
				listOfZeros = append(listOfZeros, input[j])
			}
		}

		if isOxygenRating {
			if len(listOfOnes) >= len(listOfZeros) {
				input = listOfOnes
			} else {
				input = listOfZeros
			}
		} else {
			if len(listOfZeros) <= len(listOfOnes) {
				input = listOfZeros
			} else {
				input = listOfOnes
			}
		}

		if len(input) == 1 {
			break
		}
	}
	return input[0]
}
