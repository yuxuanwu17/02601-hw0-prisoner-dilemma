package main

import (

)

// The data stored in a single cell of a field
type Cell struct {
	strategy  string //represents "C" or "D" corresponding to the type of prisoner in the cell
	score float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The game board is a 2D slice of Cell objects
type GameBoard [][]Cell

func main() {
	
}
