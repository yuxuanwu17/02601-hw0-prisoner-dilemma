package main

import (
	"fmt"
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
	// set the weight b
	b := 1.5

	// set the number of generation
	numGen := 3

	initialBoard := ReadBoardFromFile(initialBoardFile)
	//fmt.Println(initialBoard)  // 这里的strategy 和 value都可以显示出来

	boards := PlaySpatialGames(initialBoard, numGen, b)
	//fmt.Println(updateOnce)
	fmt.Println(len(boards)) //  n+1 次的循环个数

	//for i := 1; i <= numGen; i++ {
	//	fmt.Println(boards[i])
	//	fmt.Println("==============")
	//}

	for i := 0; i <= numGen; i++ {
		fmt.Println("================第", i+1, "次循环===================") // 第一次的循环board的情况
		for j := 0; j < 10; j++ {
			fmt.Println(boards[i][j]) // 第一次的循环board的情况
		}
	}

}
