package main

import (
	"fmt"
	"gifhelper"
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
	initialBoardFile := "CellularAutomata/boards/smallfield.txt" // my starting GameBoard file name
	//initialBoardFile := "CellularAutomata/boards/f99.txt" // my starting GameBoard file name
	//initialBoardFile := "CellularAutomata/boards/rand200-10.txt" // my starting GameBoard file name
	outputFileDir := "CellularAutomata/output/b_1.65_numGen_30.gif"
	// set the weight b
	//b := 1.65
	//b := 2.0
	b := 1.86

	// set the number of generation
	numGen := 10
	//numGen := 200

	// set the cell width
	cellWidth := 5

	initialBoard := ReadBoardFromFile(initialBoardFile)
	//fmt.Println(initialBoard)  // 这里的strategy 和 value都可以显示出来

	boards := PlaySpatialGames(initialBoard, numGen, b)
	//fmt.Println(updateOnce)
	//fmt.Println(len(boards)) //  n+1 次的循环个数

	for i := 0; i <= numGen; i++ {
		fmt.Println("================第", i, "次循环===================") // 第一次的循环board的情况
		for j := 0; j < 10; j++ {
			fmt.Println(boards[i][j]) // 第一次的循环board的情况
		}
	}

	imglist := DrawGameBoards(boards, cellWidth)
	//fmt.Println(imglist)

	gifhelper.ImagesToGIF(imglist, outputFileDir)

}
