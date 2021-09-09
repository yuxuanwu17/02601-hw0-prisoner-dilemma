package main

import "fmt"

func PlaySpatialGames(initialBoard GameBoard, numGens int, b float64) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	// ======================================
	// 问题的核心就是为什么我的initialBoard会变
	// 因为当第一个update完成后，他的board[0]就会变成一个固定的值，以及对应的initialBoard,可以在main函数里发现
	// 然后他每一次update前的时间都是固定的
	// ======================================

	for i := 1; i <= numGens; i++ {

		fmt.Println("在update 函数之前")
		fmt.Println("=====之前一个的情况boards[", i-1, "]的情况，是UpdateBoard的输入========= ")
		//for j := 0; j < 10; j++ {
		//	fmt.Println(boards[i-1][j]) // 第一次的循环board的情况
		//}
		boards[i] = UpdateBoard(boards[i-1], b)
		//for i := 0; i < numGen; i++ {
		//	fmt.Println("================第", i+1, "次循环===================") // 第一次的循环board的情况

		//}
		fmt.Println("在update 函数之后")
		//for j := 0; j < 10; j++ {
		//	fmt.Println(boards[i-1][j]) // 第一次的循环board的情况
		//}
	}
	return boards
}

func UpdateBoard(currBoard GameBoard, b float64) GameBoard {
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	newBoard := InitializeBoard(numRows, numCols)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			//注意这里返回的是key-val形式的{C 0}
			newBoard[r][c] = ObtainNeighbors(currBoard, r, c, numRows, numCols, b)
			if r == 4 && c == 3 {
				newBoard[r][c] = ObtainNeighbors(currBoard, r, c, numRows, numCols, b)
			}
		}
	}

	// 获得的是neighbor的值，还没有作出决策
	//fmt.Println()
	//fmt.Println("获得的是neighbor的值，还没有作出决策~~~~~~~~~~~~~~~~")
	//fmt.Println(newBoard)

	// 这里需要发生替换的操作
	newStrategyBoard := InitializeBoard(numRows, numCols)
	// 遍历整个数组
	//fmt.Println(newStrategyBoard)

	// 现在的核心问题就是如何将其固定住，不要update在新的棋盘上
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			newStrategyBoard[r][c] = StrateyReplaceByNbrs(newBoard, r, c, numRows, numCols, b)
		}
	}

	//fmt.Println(newStrategyBoard)
	//fmt.Println()
	return newStrategyBoard
}

func StrateyReplaceByNbrs(board GameBoard, i, j, numRow, numCol int, b float64) Cell {
	numRows := CountRows(board)
	numCols := CountCols(board)
	newBoard := InitializeBoard(numRows, numCols)

	// 左上角
	if i == 0 && j == 0 {
		// 这里的center还是右这样的问题
		center := board[i][j]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]
		neighbors := []Cell{east, southeast, south, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 上边行 i=-1, j 属于 【0,numCol]
	if i == 0 && j > 0 && j < numCol-1 {
		center := board[i][j]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]
		southwest := board[i+1][j-1]
		west := board[i][j-1]
		neighbors := []Cell{east, southeast, south, southwest, west, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 下边行
	if i == numRow-1 && j > 0 && j < numCol-1 {

		center := board[i][j]
		northwest := board[i-1][j-1]
		north := board[i-1][j]
		northeast := board[i-1][j+1]
		east := board[i][j+1]
		west := board[i][j-1]

		neighbors := []Cell{northwest, north, northeast, east, west, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 左边行 j <0 是固定的
	if i > 0 && i < numRow-1 && j == 0 {
		center := board[i][j]
		north := board[i-1][j]
		northeast := board[i-1][j+1]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]

		neighbors := []Cell{north, northeast, east, southeast, south, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 右边行
	if i > 0 && i < numRow-1 && j == numRow-1 {
		center := board[i][j]
		northwest := board[i-1][j-1]
		north := board[i-1][j]
		south := board[i+1][j]
		southwest := board[i+1][j-1]
		west := board[i][j-1]

		neighbors := []Cell{northwest, north, south, southwest, west, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 右下角
	if i == numRow-1 && j == numCol-1 {
		center := board[i][j]
		northwest := board[i-1][j-1]
		north := board[i-1][j]
		west := board[i][j-1]
		neighbors := []Cell{northwest, north, west, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 右上角
	if i == 0 && j == numCol-1 {
		center := board[i][j]
		south := board[i+1][j]
		southwest := board[i+1][j-1]
		west := board[i][j-1]
		neighbors := []Cell{south, southwest, west, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 左下角
	if i == numRow-1 && j == 0 {
		center := board[i][j]
		north := board[i-1][j]
		northeast := board[i-1][j+1]
		east := board[i][j+1]
		neighbors := []Cell{north, northeast, east, center}
		newBoard[i][j] = FindMaxNbr(neighbors)

	}

	// 中心neighbor
	if i > 0 && i < numRow-1 && j > 0 && j < numCol-1 {
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
		updateCell := FindMaxNbr(neighbors)
		newBoard[i][j] = updateCell
	}

	return newBoard[i][j]
}

func ObtainNeighbors(board GameBoard, i, j, numRow, numCol int, b float64) Cell {
	// 左上角
	if i == 0 && j == 0 {
		// 这里的center还是右这样的问题
		center := board[i][j]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]
		neighbors := []Cell{east, southeast, south}

		// ========================================================
		// 问题出现在这里，这里每次update 都会让initialBoard也被更新
		// ========================================================

		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 上边行 i=-1, j 属于 【0,numCol]
	if i == 0 && j > 0 && j < numCol-1 {
		center := board[i][j]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]
		southwest := board[i+1][j-1]
		west := board[i][j-1]
		neighbors := []Cell{east, southeast, south, southwest, west}
		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 下边行
	if i == numRow-1 && j > 0 && j < numCol-1 {

		center := board[i][j]
		northwest := board[i-1][j-1]
		north := board[i-1][j]
		northeast := board[i-1][j+1]
		east := board[i][j+1]
		west := board[i][j-1]

		neighbors := []Cell{northwest, north, northeast, east, west}
		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 左边行 j <0 是固定的
	if i > 0 && i < numRow-1 && j == 0 {
		center := board[i][j]
		north := board[i-1][j]
		northeast := board[i-1][j+1]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]

		neighbors := []Cell{north, northeast, east, southeast, south}
		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 右边行
	if i > 0 && i < numRow-1 && j == numRow-1 {
		center := board[i][j]
		northwest := board[i-1][j-1]
		north := board[i-1][j]
		south := board[i+1][j]
		southwest := board[i+1][j-1]
		west := board[i][j-1]

		neighbors := []Cell{northwest, north, south, southwest, west}
		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 右下角
	if i == numRow-1 && j == numCol-1 {
		center := board[i][j]
		northwest := board[i-1][j-1]
		north := board[i-1][j]
		west := board[i][j-1]
		neighbors := []Cell{northwest, north, west}
		board[i][j] = ValueCalCell(center, neighbors, b)

	}

	// 右上角
	if i == 0 && j == numCol-1 {
		center := board[i][j]
		south := board[i+1][j]
		southwest := board[i+1][j-1]
		west := board[i][j-1]
		neighbors := []Cell{south, southwest, west}
		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 左下角
	if i == numRow-1 && j == 0 {
		center := board[i][j]
		north := board[i-1][j]
		northeast := board[i-1][j+1]
		east := board[i][j+1]
		neighbors := []Cell{north, northeast, east}
		board[i][j] = ValueCalCell(center, neighbors, b)
	}

	// 中心neighbor
	if i > 0 && i < numRow-1 && j > 0 && j < numCol-1 {
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
	}

	return board[i][j]
}

// 计算中心位置的值
func FindMaxNbr(neighbors []Cell) Cell {
	tempMax := Cell{strategy: "", score: 0.0}
	for _, neighbor := range neighbors {
		if neighbor.score > tempMax.score {
			tempMax = neighbor
		}
	}
	tempMax.score = 0.0
	return tempMax

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
