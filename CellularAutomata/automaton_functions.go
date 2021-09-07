package main

type GameBoard [][]int

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
	// first, create new board corresponding to the next generation.
	// let's have all cells have state 0 to begin.
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	newBoard := InitializeBoard(numRows, numCols)

	//now, update values of newBoard
	//range through all cells of currBoard and update each one into newBoard.
	for r := 1; r < numRows-1; r++ {
		// r will range over rows of board
		// current row is currBoard[r]
		// range over values in currBoard[r]
		for c := 1; c < numCols-1; c++ {
			//curr value is currBoard[r][c]
			newBoard[r][c] = UpdateCell(currBoard, r, c)
		}
	}

	// return newBoard
	return newBoard
}

//UpdateCell takes a gameboard along with row and col indices and
//it returns the state of the board at these row/col indices is in the next generation
func UpdateCell(board GameBoard, r, c int) int {
	numCorpNbrs := CountCorpCenterNbrs(board, r, c)
	if board[r][c] == 1 {
		// corporation

		// 比较
		// board[r][c] ==0 的策略，算出来一个值
		// 这里写是update的策略

	} else {
		// defect

	}

}
return -1 // we know we didn't find a match
}


func CountCorpCenterNbrs(board GameBoard, r, c int) int {
	count := 0

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
}

// InField takes a GameBoard board as well as row and col indices (i,j) and returns truer if board[i][j] is in the board and false otherwise
func InField(board GameBoard, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= CountRows(board) || j >= CountCols(board) {
		return false
	}

	return true

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
		board[r] = make([]int, numCols)
	}

	return board
}
