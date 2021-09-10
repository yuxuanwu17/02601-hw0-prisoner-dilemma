package main

func PlaySpatialGames(initialBoard GameBoard, numGens int, b float64) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard
	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1], b)
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
		}
	}

	newStrategyBoard := InitializeBoard(numRows, numCols)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			newStrategyBoard[r][c] = StrateyReplaceByNbrs(newBoard, r, c, numRows, numCols, b)
		}
	}

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
