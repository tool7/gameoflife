package utils

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func InitInstructionsText(screenWidth, screenHeight int) *text.Text {
	face, err := LoadTTF("intuitive.ttf", 36)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(float64(screenWidth / 2), float64(screenHeight - 60)), atlas)

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

func InitGameStatusText(screenWidth, screenHeight int) *text.Text {
	face, err := LoadTTF("intuitive.ttf", 60)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(float64(screenWidth - 350), float64(screenHeight - 120)), atlas)

	return text
}

func DrawGameStatusText(win *pixelgl.Window, textHandle *text.Text, isGameStarted bool) {
	textHandle.Clear()
	
	if isGameStarted {
		fmt.Fprintf(textHandle, "RUNNING...")
	} else {
		fmt.Fprintf(textHandle, "STOPPED")
	}

	textHandle.Draw(win, pixel.IM)
}

func InitGameIntervalText(screenHeight int) *text.Text {
	face, err := LoadTTF("intuitive.ttf", 42)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	text := text.New(pixel.V(50, float64(screenHeight - 120)), atlas)

	return text
}

func DrawGameIntervalText(win *pixelgl.Window, textHandle *text.Text, currentGameIntervalInMs int) {
	textHandle.Clear()
	
	fmt.Fprintf(textHandle, "Sleep interval:  %d ms", currentGameIntervalInMs)

	textHandle.Draw(win, pixel.IM)
}
