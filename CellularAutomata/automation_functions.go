package main

// neighborhood "moore" string
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

	// 一种思路是从0 到 nRows 进行计算，然后用if加以判断，得到新的棋盘上的值
	//
	//这里update 边角的值
	//然后用currBoard进行接受

	//currBoard = UpdateCornerCell(currBoard, numRows, numCols, b)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			//注意这里返回的是key-val形式的{C 0}
			newBoard[r][c] = UpdateCell(currBoard, r, c, numRows, numCols, b)
		}
	}

	// return newBoard
	return newBoard
}

//UpdateCell
func UpdateCell(board GameBoard, r, c, numRow, numCol int, b float64) Cell {
	// 返回四周最大的值
	// 这里的r，c代表的
	updatedBoard := FindMaxNbrs(board, r, c, numRow, numCol, b)

	return updatedBoard
}

// 找到四周的nbrs
// i, j, r, c 的问题可能存在

func FindMaxNbrs(board GameBoard, r, c, numRow, numCol int, b float64) Cell {

	// r,c 是中心的位置

	// 计算其值
	// 根据neighbor的情况来获得board[r][c]的值
	board[r][c] = ObtainNeighbors(board, r, c, numRow, numCol, b)
	return board[r][c]
}

// i（横）,j（纵） 代表的是neighbor的位置 [0,0] [numCol,numRow]的情况
// 可能会出现outOfbounds的情况
// 关于边和行的问题，他需要考虑是否重叠的问题

func ObtainNeighbors(board GameBoard, i, j, numRow, numCol int, b float64) Cell {
	// 左上角
	if i == 0 && j == 0 {
		// 这里的center还是右这样的问题
		center := board[i][j]
		east := board[i][j+1]
		southeast := board[i+1][j+1]
		south := board[i+1][j]
		neighbors := []Cell{east, southeast, south}
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

func ValueCalCell(center Cell, neighbors []Cell, b float64) Cell {
	var totalVal float64 = 0

	for _, neighbor := range neighbors {
		// strategy is cooperation

		centerState := center.strategy

		if centerState == "C" {
			// centerState 为 C
			if neighbor.strategy == centerState {
				totalVal++
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
