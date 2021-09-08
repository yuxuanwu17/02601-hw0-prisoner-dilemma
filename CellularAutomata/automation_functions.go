package main

// neighborhood "moore" string
func PlaySpatialGames(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1])
	}

	return boards
}

func UpdateBoard(currBoard GameBoard) GameBoard {
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	newBoard := InitializeBoard(numRows, numCols)

	// 一种思路是从0 到 nRows 进行计算，然后用if加以判断，得到新的棋盘上的值
	//

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			//注意这里返回的是key-val形式的{C 0}
			newBoard[r][c] = UpdateCell(currBoard, r, c)
		}
	}

	// return newBoard
	return newBoard
}

//UpdateCell
func UpdateCell(board GameBoard, r, c int) GameBoard {
	// 返回四周最大的值
	updatedBoard := FindMaxNbrs()

	return updatedBoard
}

// 找到四周的nbrs
func FindMaxNbrs(board GameBoard, r, c int) string {

	// 计算其值
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i != r || j != c) &&
				InField(board, i, j) {
				if board[i][j] == 1 {
					count++
				}
			}
		}

	}
	return ""
}

// 计算中心位置的值
func ValueCalc(board) {

}

func RuleMatch(board GameBoard, r, c int) {
	// r-1 < 0

	// c-1 < 0

}

//// InField takes a GameBoard board as well as row and col indices (i,j) and returns truer if board[i][j] is in the board and false otherwise
func InField(board GameBoard, i, j int) int {
	if i < 0 || j < 0 {
		return false
	}
	if i >= CountRows(board) || j >= CountCols(board) {
		return false
	}

	return 1

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

//关于如何update四周的问题
