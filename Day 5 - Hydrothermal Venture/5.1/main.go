package main

import (
	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
	"log"
)

type Vector struct {
	start Coordinate
	end   Coordinate
}

type Coordinate struct {
	x int
	y int
}

type Space struct {
	width   int
	height  int
	overlap [][]int
}

const MINIMUM_OVERLAP = 2

func main() {
	// read the input of the file
	log.Println("Reading the input file...")
	input := utils.ReadInput("../input.txt")

	// Setup Environment
	log.Println("Setting up the environment...")
	space, vectors := setup(input)
	log.Printf("Created space-matrix of '%d' * '%d' with '%d' vectors.\n", space.width, space.height, len(vectors))

	// Check if vectors overlap in the space
	log.Println("Checking if vectors overlap...")
	space = checkOverlap(space, vectors)

	// Find the coordinate with the most amount of overlap
	coordinates := getCoordinatesWithMinimumOverlap(space, MINIMUM_OVERLAP)
	log.Printf("There are '%d' coordinates with at least '%d' overlap.\n", len(coordinates), MINIMUM_OVERLAP)
	//log.Printf("space %v", space.overlap)
}

func getCoordinatesWithMinimumOverlap(space Space, maxOverlap int) []Coordinate {
	coordinates := make([]Coordinate, 0)

	for x := 0; x < space.width; x++ {
		for y := 0; y < space.height; y++ {
			if space.overlap[x][y] >= maxOverlap {
				coordinates = append(coordinates, Coordinate{x, y})
			}
		}
	}

	return coordinates
}

func checkOverlap(space Space, vectors []Vector) Space {
	matrix := make([][]int, space.width)
	for i := range matrix {
		matrix[i] = make([]int, space.height)
	}

	for _, vector := range vectors {
		if vector.start.x == vector.end.x || vector.start.y == vector.end.y {
			if vector.start.x > vector.end.x {
				for x := vector.start.x; x >= vector.end.x; x-- {
					if vector.start.y > vector.end.y {
						for y := vector.start.y; y >= vector.end.y; y-- {
							matrix[x-1][y-1]++
						}
					} else {
						for y := vector.start.y; y <= vector.end.y; y++ {
							matrix[x-1][y-1]++
						}
					}
				}
			} else if vector.start.x <= vector.end.x{
				for x := vector.start.x; x <= vector.end.x; x++ {
					if vector.start.y > vector.end.y {
						for y := vector.start.y; y >= vector.end.y; y-- {
							matrix[x-1][y-1]++
						}
					} else {
						for y := vector.start.y; y <= vector.end.y; y++ {
							matrix[x-1][y-1]++
						}
					}
				}
			}
		}
	}
	space.overlap = matrix
	return space
}

func setup(input []string) (Space, []Vector) {
	maxX := 0
	maxY := 0
	vectors := make([]Vector, 0)
	for _, line := range input {
		// get the fuel required for the module
		coordinates := utils.SplitDigitsFromSeperatedString(line)

		// create the vector
		vector := Vector{
			start: Coordinate{
				x: coordinates[0],
				y: coordinates[1],
			},
			end: Coordinate{
				x: coordinates[2],
				y: coordinates[3],
			},
		}
		vectors = append(vectors, vector)

		// Find maximum X & Y coordinate.
		if vector.start.x > maxX {
			maxX = vector.start.x
		}
		if vector.end.x > maxX {
			maxX = vector.end.x
		}
		if vector.start.y > maxY {
			maxY = vector.start.y
		}
		if vector.end.y > maxY {
			maxY = vector.end.y
		}
	}

	return Space{
		width:  maxX,
		height: maxY,
	}, vectors
}
