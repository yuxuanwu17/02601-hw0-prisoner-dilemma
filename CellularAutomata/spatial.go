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

	outputFileDir := os.Args[2]

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

	imglist := DrawGameBoards(boards, cellWidth)
	fmt.Println("Boards drawn to images! Now, convert to animated GIF.")

	finalImage := imglist[numGen]

	//fmt.Println(finalImage)
	gifhelper.ImageToPNG(finalImage, outputFileDir)
	fmt.Println("The final image had be drawn!")

	gifhelper.ImagesToGIF(imglist, outputFileDir)
	fmt.Println("Success! GIF produced.")

	//	./CellularAutomata boards/f99.txt output/f99_165_0.gif 1.65 0
}
