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
	b := 2.0

	// set the number of generation
	numGen := 10

	initialBoard := ReadBoardFromFile(initialBoardFile)
	//fmt.Println(initialBoard)  // 这里的strategy 和 value都可以显示出来

	boards := PlaySpatialGames(initialBoard, numGen, b)
	//fmt.Println("===================0======================")
	//fmt.Println(len(boards[0]))
	////fmt.Println(len(boards))
	//for i := 0; i < 10; i++ {
	//	fmt.Println(boards[0][i])
	//}
	//fmt.Println("===================1======================")
	////fmt.Println(len(boards[1]))
	//for i := 0; i < 12; i++ {
	//	fmt.Println(boards[1][i])
	//}
	//fmt.Println("===================2======================")
	////fmt.Println(boards[2])
	//for i := 0; i < 14; i++ {
	//	fmt.Println(boards[2][i])
	//}

	//for i := 1; i <= numGen; i++ {
	//	fmt.Println(boards[i])
	//	fmt.Println("==============")
	//}

	for i := 0; i <= numGen; i++ {
		fmt.Println("================第", i, "次循环===================") // 第一次的循环board的情况
		for j := 0; j < 10; j++ {
			fmt.Println(boards[i][j]) // 第一次的循环board的情况
		}
	}

}
