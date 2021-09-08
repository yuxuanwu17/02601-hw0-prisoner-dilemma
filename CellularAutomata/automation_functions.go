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

	for r := 1; r < numRows-1; r++ {
		for c := 1; c < numCols-1; c++ {
			//注意这里返回的是key-val形式的{C 0}
			newBoard[r][c] = UpdateCell(currBoard, r, c, b)
		}
	}

	// return newBoard
	return newBoard
}

//UpdateCell
func UpdateCell(board GameBoard, r, c int, b float64) Cell {
	// 返回四周最大的值
	updatedBoard := FindMaxNbrs(board, r, c, b)

	return updatedBoard
}

// 找到四周的nbrs
func FindMaxNbrs(board GameBoard, r, c int, b float64) Cell {

	// 计算其值
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			// 获得neighbor的值
			center := board[r][c]
			northwest := board[r-1][c-1]
			north := board[r-1][c]
			northeast := board[r-1][c+1]
			east := board[r][c+1]
			southeast := board[r+1][c+1]
			south := board[r+1][c]
			southwest := board[r+1][c-1]
			west := board[r][c-1]

			neighbors := []Cell{northwest, north, northeast, east, southeast, south, southwest, west}
			board[r][c] = ValueCalCell(center, neighbors, b)
		}
	}
	return board[r][c]
}

//if i != j {
//				// 获得board中的GameBoard的值
//
//				// 如国是合作C
//				if board[i][j].strategy=="C"{
//					//if neighbors
//				}else{
//
//				}
//}

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

////// InField takes a GameBoard board as well as row and col indices (i,j) and returns truer if board[i][j] is in the board and false otherwise
//func InField(board GameBoard, i, j int) int {
//	if i < 0 || j < 0 {
//		return false
//	}
//	if i >= CountRows(board) || j >= CountCols(board) {
//		return false
//	}
//
//	return 1
//
//}

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

//关于如何update四周的问题
