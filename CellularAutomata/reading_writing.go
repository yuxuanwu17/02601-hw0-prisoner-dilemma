package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// NEED function to read in rule in current form of Moore and then produce all possible rules.

//ReadRulesFromFile takes a file and reads the rule strings provided in this file.
//It stores the result in a list of strings.
func ReadRulesFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ruleStrings := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		ruleStrings = append(ruleStrings, currentLine)
	}

	return ruleStrings
}

//WriteStringsToFile takes a collection of strings and a filename and
//writes these strings to the given file, with each string on one line.
func WriteStringsToFile(patterns []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for _, pattern := range patterns {
		fmt.Fprintln(file, pattern)
	}
}

//

//func Convert2Int(filename string) GameBoard {
//
//}

//ReadBoardFromFile takes a filename as a string and reads in the data provided
//in this file, returning a game board.
func ReadBoardFromFile(filename string) GameBoard {
	board := make(GameBoard, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		// skip the first line of the scanner
		if count == 0 {
			count++
			continue
		}
		currentLine := scanner.Text()
		currentArray := make([]int, 0)

		for i := range currentLine {
			// converting the string format to the binary 1 or 0 files
			var val = currentLine[i : i+1]
			var valInt int
			if val == "C" {
				// cooperation
				valInt = 1
			} else {
				//defect
				valInt = 0
			}
			//
			//if err2 != nil {
			//	panic("Error: Issue converting string to int from file.")
			//}
			currentArray = append(currentArray, valInt)
		}

		board = append(board, currentArray)
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return board
}

//WriteBoardToFile takes a gameboard and a filename as a string and writes the
//gameboard to the specified output file.
func WriteBoardToFile(board GameBoard, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for r := range board {
		for c := range board[r] {
			fmt.Fprint(file, board[r][c])
		}
		fmt.Fprintln(file)
	}
}
