package main

import (
	"log"

	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
)

const BINGO_COLUMN_SIZE = 6

func main() {

	// Read input
	input := utils.ReadInput("../input.txt")

	// Split the Bingo boards and the called numbers of the input file.
	boards, bingoNumbers := getBingoInput(input)

	log.Println("Bingo numbers:", bingoNumbers)
	log.Println("Bingo boards:", boards)

}

func getBingoInput(input []string) ([][][]int, []int) {

	// Get the called bingo numbers.
	bingoNumbers := utils.SplitDigitsFromSeperatedString(input[0])

	// Remove the first line of the input file.
	input = input[1:]

	// Get the bingo boards.
	counter := 0
	boards := make([][][]int, 0)
	board := make([][]int, 0)

	for i := 0; i < len(input); i++ {
		if i%BINGO_COLUMN_SIZE != 0 {
			board = append(board, utils.SplitDigitsFromSeperatedString(input[i]))
			counter++

			if counter == BINGO_COLUMN_SIZE {
				boards = append(boards, board)
				board = make([][]int, 0)
				counter = 0
			}
		}
	}

	return boards, bingoNumbers
}
