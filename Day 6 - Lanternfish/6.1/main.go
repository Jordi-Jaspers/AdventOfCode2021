package main

import (
	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
	"log"
)

// Each Lantern fish spawns a new lantern fish every 7 days
// The days they spawn a new lantern fish are not synchronized.
// A new Lantern fish needs 2 days Longer to spawn a new lantern fish.

const OBSERVATION_DAYS = 80
const SPAWN_CYCLE = 7
const SPAWN_CYCLE_NEWBORN = SPAWN_CYCLE + 2

func main() {
	// Read the input file.
	log.Println("Reading input file...")
	input := utils.ReadInput("../input.txt")

	// Parse the input.
	log.Println("Parsing input...")
	initialDays := utils.SplitDigitsFromSeperatedString(input[0])

	// Sort the input by using merge sort.
	log.Println("Calculating the amount of laters fish after X days...")
	log.Println("Initial days: ", len(initialDays))
	initialDays = utils.MergeSort(initialDays)

	// Calculate the amount of lantern fish after X days. --> 375482
	calculate(initialDays)
}

func calculate(fishes []int) {
	log.Println("After 0 days: ", fishes)
	//prevCounter := 0

	for i := 1; i <= OBSERVATION_DAYS-1; i++ {
		parents := make([]int, 0)
		babies := make([]int, 0)
		counter := 0


		for j := 0; j < len(fishes); j++ {
			fishes[j]--
			if fishes[j] == 0 {
				parents = append (parents, SPAWN_CYCLE)
				babies = append (babies, SPAWN_CYCLE_NEWBORN)
				counter++
			}
		}

		//log.Println("parents: ", parents)
		//fishes = fishes[counter:]
		//insertPosition := len(fishes) - prevCounter
		//tempList := fishes[insertPosition:]
		//log.Println("fish with no-zeros: ", fishes)
		//
		//
		//parents = append(fishes[:insertPosition], parents...)
		//log.Println("fishes[:insertPosition]: ", fishes[:insertPosition])
		//log.Println("fish + parents: ", parents)
		//
		//for _, v := range tempList {
		//	parents = append(parents, v)
		//}
		//time.Sleep(10*time.Second)
		//log.Println("precounter: ", prevCounter)
		//log.Println("fish start + parents + fish end: ", parents)
		//fishes = append(parents, babies...)
		//prevCounter = counter

		parents = append (parents, babies...)
		fishes = append (fishes[counter:], parents...)
		fishes = utils.MergeSort(fishes)
		log.Println("After ", i, " day(s): ", len(fishes))
	}
}
