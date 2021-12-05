package main

import (
	utils "github.com/Jordi-Jaspers/AdventOfCode2021/Util"
	"log"
	"sync"
)

const BOARD_COLUMN_SIZE = 5
const BOARD_ROW_SIZE = 5

type BingoBoard struct {
	id 				int
	bingoBoard 		[][]int
	stamps     		int
	winningNumber 	int
}

func main() {

	// Read input
	input := utils.ReadInput("../input.txt")

	// Split the Bingo boards and the called numbers of the input file.
	boards, bingoNumbers := getBingoInput(input)

	// Play the Bingo game.
	board := playBingo(boards, bingoNumbers)

	// sum all the numbers of the board.
	sum := 0
	for _, row := range board.bingoBoard {
        for _, number := range row {
            sum += number
        }
    }

	// Print the result.
	log.Println("The result is:", sum * board.winningNumber)
}

// Check if the bingo number is present in on of the boards.
func playBingo(boards [][][]int, bingoNumbers []int) BingoBoard {
	var wg sync.WaitGroup

	channel := make(chan BingoBoard, len(boards))
	log.Printf("Processing '%d' boards", len(boards))

	// Start the goroutines.
	for _, board := range boards {
        wg.Add(1)
        go func(board [][]int, bingoNumbers []int) {
            defer wg.Done()
            channel <- stampBingoNumber(board, bingoNumbers)
        }(board, bingoNumbers)
    }

	log.Println("Waiting for the goroutines to finish...")
	wg.Wait()
	log.Println("All goroutines finished.")

	output := make([]BingoBoard, len(boards))
	for i := 0; i < len(boards); i++ {
		output[i] = <-channel
	}

	close(channel)

	log.Printf("Received '%d' results", len(output))
	winner := BingoBoard{}
	for _, board := range output {
		if winner.stamps == 0 || board.stamps < winner.stamps {
			winner = board
		}
    }

	log.Printf("The winner is board '%d' with '%d' stamps.", winner.id, winner.stamps)
	return winner
}

// Change all matching bingo numbers on the board to '0'
func stampBingoNumber(board [][]int, bingoNumbers []int) BingoBoard {
	counter := 0
	for _, bingoNumber := range bingoNumbers {
		counter++
		for i := 0; i < BOARD_ROW_SIZE; i++ {
			for j := 0; j < BOARD_COLUMN_SIZE; j++ {
				if board[i][j] == bingoNumber {
					board[i][j] = 0
					if counter > 5 {
						if hasBingo(board, counter) {
                            return BingoBoard{
								id: counter,
								bingoBoard: board,
								stamps: counter,
								winningNumber: bingoNumber,
							}
                        }
					}
				}
			}
		}
	}
	return BingoBoard{}
}

// Check if the board has at lease 1 row or column complete with the bingo numbers.
func hasBingo(board [][]int, counter int) bool {
	//Check horizontal lines
	for i := 0; i < BOARD_ROW_SIZE; i++ {
		for j := 0; j < BOARD_COLUMN_SIZE; j++ {
			if board[i][j] != 0 {
				break
			} else if j == BOARD_COLUMN_SIZE - 1 {
				return true
			}
		}
	}

	//Check vertical lines
	for i := 0; i < BOARD_COLUMN_SIZE; i++ {
		for j := 0; j < BOARD_ROW_SIZE; j++ {
			if board[j][i] != 0 {
				break
			} else if j == BOARD_ROW_SIZE - 1 {
				return true
			}
		}
	}
	return false
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
		if i%(BOARD_COLUMN_SIZE+1) != 0 {
			board = append(board, utils.SplitDigitsFromSeperatedString(input[i]))
			counter++

			if counter == BOARD_COLUMN_SIZE {
				boards = append(boards, board)
				board = make([][]int, 0)
				counter = 0
			}
		}
	}

	return boards, bingoNumbers
}
