package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func initInstructionsText() *text.Text {
	face, err := loadTTF("intuitive.ttf", 36)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(screenWidth / 2, screenHeight - 60), atlas)

	instructions := []string{
		"Press SPACE to toggle simulation, C to clear the grid and +/- to change sleep interval\n",
		"Press 1 through 4 to insert different patterns:",
		"1) RPentomino   2) Diehard   3) Acorn   4) GosperGliderGun",
	}

	for _, line := range instructions {
		text.Dot.X -= text.BoundsOf(line).W() / 2
		fmt.Fprintln(text, line)
	}

	return text
}

func initGameStatusText() *text.Text {
	face, err := loadTTF("intuitive.ttf", 60)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(screenWidth - 350, screenHeight - 120), atlas)

	return text
}

func drawGameStatusText(win *pixelgl.Window, textHandle *text.Text) {
	textHandle.Clear()
	
	if isGameStarted {
		fmt.Fprintf(textHandle, "RUNNING...")
	} else {
		fmt.Fprintf(textHandle, "STOPPED")
	}

	textHandle.Draw(win, pixel.IM)
}

func initGameIntervalText() *text.Text {
	face, err := loadTTF("intuitive.ttf", 42)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(50, screenHeight - 120), atlas)

	return text
}

func drawGameIntervalText(win *pixelgl.Window, textHandle *text.Text) {
	textHandle.Clear()
	
	fmt.Fprintf(textHandle, "Sleep interval:  %d ms", currentGameIntervalInMs)

	textHandle.Draw(win, pixel.IM)
}
