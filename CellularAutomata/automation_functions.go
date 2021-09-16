package main

import (
	"fmt"
	"time"
)

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("该函数执行完成耗时 = %v\n", tc)
	}
}

func PlaySpatialGames(initialBoard GameBoard, numGens int, b float64) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard
	for i := 1; i <= numGens; i++ {
		fmt.Println("Round", i)
		//start := time.Now()

		boards[i] = UpdateBoard(boards[i-1], b)
		//tc := time.Since(start)
		//fmt.Printf("该函数执行完成耗时 = %v\n", tc)
	}
	return boards
}

func UpdateBoard(currBoard GameBoard, b float64) GameBoard {
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)

	// surround with D
	currBoardWithD := SurroundWithD(currBoard, numRows, numCols)

	// 开始读取，读取到newBoard上 2ms 左右
	newBoard := SurroundWithD(currBoard, numRows, numCols)
	for r := 1; r <= numRows; r++ {
		for c := 1; c <= numCols; c++ {
			newBoard[r][c] = ObtainNeighbors(currBoardWithD, r, c, b)
		}
	}

	// 创建新表，接受酒标
	newStrategyBoard := SurroundWithD(currBoard, numRows, numCols)

	//=============================
	// 定位到bug
	//=============================
	start := time.Now()

	for r := 1; r <= numRows; r++ {
		for c := 1; c <= numCols; c++ {
			newStrategyBoard[r][c] = StrategyReplaceByNbrs(newBoard, r, c)
		}
	}

	tc := time.Since(start)
	fmt.Printf("该函数执行完成耗时 = %v\n", tc)

	// 去除外层的D
	finalStrategyBoard := InitializeBoard(numRows, numCols)
	for r := 1; r < numRows+1; r++ {
		for c := 1; c < numCols+1; c++ {
			finalStrategyBoard[r-1][c-1] = newStrategyBoard[r][c]
		}
	}

	return finalStrategyBoard
}

func SurroundWithD(currBoard GameBoard, numRows, numCols int) GameBoard {
	// 初始+2的棋盘
	currBoardWithD := InitializeBoard(numRows+2, numCols+2)

	for r := 0; r < numRows+2; r++ {
		currBoardWithD[r][0] = Cell{strategy: "D", score: 0.0}
		currBoardWithD[r][numCols+1] = Cell{strategy: "D", score: 0.0}
	}

	for c := 0; c < numCols+1; c++ {
		currBoardWithD[0][c] = Cell{strategy: "D", score: 0.0}
		currBoardWithD[numRows+1][c] = Cell{strategy: "D", score: 0.0}
	}

	// 将老棋盘打入带D的新棋盘

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			currBoardWithD[r+1][c+1] = currBoard[r][c]
		}
	}
	return currBoardWithD
}

func StrategyReplaceByNbrs(board GameBoard, i, j int) Cell {

	center := board[i][j]
	northwest := board[i-1][j-1]
	north := board[i-1][j]
	northeast := board[i-1][j+1]
	east := board[i][j+1]
	southeast := board[i+1][j+1]
	south := board[i+1][j]
	southwest := board[i+1][j-1]
	west := board[i][j-1]

	neighbors := []Cell{northwest, north, northeast, east, southeast, south, southwest, west, center}
	return FindMaxNbr(neighbors)
}

func FindMaxNbr(neighbors []Cell) Cell {
	tempMax := Cell{strategy: "", score: 0.0}
	for _, neighbor := range neighbors {
		if neighbor.score >= tempMax.score {
			tempMax = neighbor
		}
	}
	tempMax.score = 0.0
	return tempMax

}

func ObtainNeighbors(board GameBoard, i, j int, b float64) Cell {
	center := board[i][j]
	northwest := board[i-1][j-1]
	north := board[i-1][j]
	northeast := board[i-1][j+1]
	east := board[i][j+1]
	southeast := board[i+1][j+1]
	south := board[i+1][j]
	southwest := board[i+1][j-1]
	west := board[i][j-1]

	neighbors := []Cell{northwest, north, northeast, east, southeast, south, southwest, west}
	board[i][j] = ValueCalCell(center, neighbors, b)

	return board[i][j]
}

func ValueCalCell(center Cell, neighbors []Cell, b float64) Cell {
	var totalVal float64 = 0

	for _, neighbor := range neighbors {
		// strategy is cooperation

		centerState := center.strategy

		if centerState == "C" {
			if neighbor.strategy == centerState {
				totalVal = totalVal + 1
			} else {
				totalVal = totalVal + 0
			}

		} else {
			// centerState 为 D
			if neighbor.strategy == centerState {
				totalVal = totalVal + 0
			} else {
				totalVal = totalVal + b
			}
		}

	}
	center.score = totalVal
	return center

}

func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}

//InitializeBoard takes a number of rows and columns as inputs and
//returns a gameboard with appropriate number of rows and colums, where all values = 0.
func InitializeBoard(numRows, numCols int) GameBoard {
	// make a 2-D slice (default values = false)
	var board GameBoard
	board = make(GameBoard, numRows)
	// now we need to make the rows too
	for r := range board {
		board[r] = make([]Cell, numCols)
	}
	return board
}
