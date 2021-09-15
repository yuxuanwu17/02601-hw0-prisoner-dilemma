package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
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

/*
type Cell struct {
	strategy string  //represents "C" or "D" corresponding to the type of prisoner in the cell
	score    float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}
*/

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
		currentArray := make([]Cell, 0)

		for i := range currentLine {
			val := currentLine[i : i+1]
			cell := Cell{strategy: val, score: 0}
			currentArray = append(currentArray, cell)
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

func ImageToPNG(finalImalge image.Image, filename string) {
	outputFile, err := os.Create(filename + ".png")
	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, finalImalge)
	if err != nil {
		fmt.Println(err)
	}

}

//ImagesToGIF() takes a slice of images and uses them to generate an animated GIF
// with the name "filename.out.gif" where filename is an input parameter.
func ImagesToGIF(imglist []image.Image, filename string) {

	// get ready to write images to files
	w, err := os.Create(filename + ".out.gif")

	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}

	defer w.Close()
	var g gif.GIF
	g.Delay = make([]int, len(imglist))
	g.Image = make([]*image.Paletted, len(imglist))
	g.LoopCount = 10

	for i := range imglist {
		g.Image[i] = ImageToPaletted(imglist[i])
		g.Delay[i] = 1
	}

	gif.EncodeAll(w, &g)
}

// ImageToPaletted converts an image to an image.Paletted with 256 colors.
// It is used by a subroutine by process() to generate an animated GIF.
func ImageToPalettedVersion1(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, palette.WebSafe)
		draw.Draw(pm, pm.Bounds(), img, image.Point{}, draw.Src)
	}
	return pm
}

var mapOfColorIndices map[color.Color]uint8

func init() {
	mapOfColorIndices = make(map[color.Color]uint8)
}

func ImageToPaletted(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, palette.WebSafe)
		var prevC color.Color = nil
		var idx uint8
		var ok bool
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				c := img.At(x, y)
				if c != prevC {
					if idx, ok = mapOfColorIndices[c]; !ok {
						idx = uint8(pm.Palette.Index(c))
						mapOfColorIndices[c] = idx
					}
					prevC = c
				}
				i := pm.PixOffset(x, y)
				pm.Pix[i] = idx
			}
		}
	}
	return pm
}
