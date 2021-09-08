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

	initialBoard := ReadBoardFromFile(initialBoardFile)
	//fmt.Println(initialBoard)  // 这里的strategy 和 value都可以显示出来

	updateOnce := PlaySpatialGames(initialBoard, 1, 3)
	fmt.Println(len(updateOnce))    //  n+1 次的循环个数
	fmt.Println(len(updateOnce[1])) // 第一次的循环board的情况
	//fmt.Println(updateOnce[1])
	for i := 0; i < 10; i++ {
		fmt.Println("===========第", i, "行================")
		fmt.Println(updateOnce[1][i])
	}

}
