package main

import (
	"fmt"
	"gifhelper"
	"os"
	"strconv"
)

// The data stored in a single cell of a field
type Cell struct {
	strategy string  //represents "C" or "D" corresponding to the type of prisoner in the cell
	score    float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The game board is a 2D slice of Cell objects
type GameBoard [][]Cell

func main() {
	fmt.Println("Prisoner paradox initialized")

	initialBoardFile := os.Args[1]
	//initialBoardFile := "CellularAutomata/boards/smallfield.txt" // my starting GameBoard file name
	outputFileDir := os.Args[2]
	//outputFileDir := "CellularAutomata/output/b_1.65_numGen_120.gif"

	// set the weight b
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		panic("Error: Problem converting cell width parameter to an integer.")
	}
	//b := 1.65

	// set the number of generation
	numGen, err := strconv.Atoi(os.Args[4])
	if err != nil {
		panic("Error: Problem converting cell width parameter to an integer.")
	}

	// set the cell width
	cellWidth := 5

	initialBoard := ReadBoardFromFile(initialBoardFile)

	boards := PlaySpatialGames(initialBoard, numGen, b)

	fmt.Println("Automaton played. Now, drawing images.")

	//for i := 0; i <= numGen; i++ {
	//	fmt.Println("================第", i, "次循环===================") // 第一次的循环board的情况
	//	for j := 0; j < 10; j++ {
	//		fmt.Println(boards[i][j]) // 第一次的循环board的情况
	//	}
	//}

	imglist := DrawGameBoards(boards, cellWidth)
	fmt.Println("Boards drawn to images! Now, convert to animated GIF.")

	gifhelper.ImagesToGIF(imglist, outputFileDir)
	fmt.Println("Success! GIF produced.")

}
