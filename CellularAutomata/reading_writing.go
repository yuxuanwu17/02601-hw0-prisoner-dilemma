package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"os"
)

// NEED function to read in rule in current form of Moore and then produce all possible rules.

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

func ImagesToGIF(imglist []image.Image, filename string) {

	// get ready to write images to files
	w, err := os.Create(filename + ".gif")

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
