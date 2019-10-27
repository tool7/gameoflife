package main

import (
	"fmt"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	screenWidth = 2400
	screenHeight = 1800
	cellSize = 10
	cellMargin = 0
	cellsHorizontalCount = 240
	cellsVerticalCount = 160
)

var (
	isGameStarted bool = false
	currentGameIntervalInMs int = 70
)

type cell struct {
	isAlive bool
	x       float64
	y       float64
}

type grid struct {
	cells [cellsHorizontalCount][cellsVerticalCount]cell
}

func getCellIndiciesByMousePosition(mousePosition pixel.Vec) (xIndex int, yIndex int) {
	xIndex = int(math.Floor(mousePosition.X / cellSize))
	yIndex = int(math.Floor(mousePosition.Y / cellSize))

	if xIndex < 0 { xIndex = 0 }
	if yIndex < 0 { yIndex = 0 }
	return
}

func getAliveNeighbourCellsCount(grid *grid, x int, y int) int {
	if x < 0 || x >= cellsHorizontalCount || y < 0 || y >= cellsVerticalCount {
		return -1
	}

	count := 0

	if x + 1 < cellsHorizontalCount && grid.cells[x + 1][y].isAlive {
		count++
	}
	if x - 1 > 0 && grid.cells[x - 1][y].isAlive {
		count++
	}
	if y + 1 < cellsVerticalCount && grid.cells[x][y + 1].isAlive {
		count++
	}
	if y - 1 > 0 && grid.cells[x][y - 1].isAlive {
		count++
	}
	if x + 1 < cellsHorizontalCount && y + 1 < cellsVerticalCount && grid.cells[x + 1][y + 1].isAlive {
		count++
	}
	if x + 1 < cellsHorizontalCount && y - 1 > 0 && grid.cells[x + 1][y - 1].isAlive {
		count++
	}
	if x - 1 > 0 && y + 1 < cellsVerticalCount && grid.cells[x - 1][y + 1].isAlive {
		count++
	}
	if x - 1 > 0 && y - 1 > 0 && grid.cells[x - 1][y - 1].isAlive {
		count++
	}

	return count
}

func clearGrid(grid *grid) {
	for i := 0; i < cellsHorizontalCount; i++ {
		for j := 0; j < cellsVerticalCount; j++ {
			grid.cells[i][j].isAlive = false
		}
	}
}

func handleMouseEvents(win *pixelgl.Window, grid *grid) {
	if isGameStarted { return }

	mousePosition := win.MousePosition()
	cellXIndex, cellYIndex := getCellIndiciesByMousePosition(mousePosition)
	if cellXIndex >= cellsHorizontalCount || cellYIndex >= cellsVerticalCount {
		return
	}

	cell := &grid.cells[cellXIndex][cellYIndex]

	if win.Pressed(pixelgl.MouseButtonLeft) {
		cell.isAlive = true
	}
	if win.Pressed(pixelgl.MouseButtonRight) {
		cell.isAlive = false
	}
}

func handleKeyboardEvents(win *pixelgl.Window, grid *grid, gameIntervalChannel chan int) {
	if win.JustPressed(pixelgl.KeySpace) {
		isGameStarted = !isGameStarted
	}

	if isGameStarted { return }

	if win.JustPressed(pixelgl.KeyC) {
		clearGrid(grid)
	}

	mousePosition := win.MousePosition()

	if win.JustPressed(pixelgl.Key1) {
		addPatternToGrid(rPentomino, mousePosition, grid)
	}
	if win.JustPressed(pixelgl.Key2) {
		addPatternToGrid(diehard, mousePosition, grid)
	}
	if win.JustPressed(pixelgl.Key3) {
		addPatternToGrid(acorn, mousePosition, grid)
	}
	if win.JustPressed(pixelgl.Key4) {
		addPatternToGrid(gosperGliderGun, mousePosition, grid)
	}

	if win.JustPressed(pixelgl.KeyKPAdd) {
		gameInterval := currentGameIntervalInMs + 10
		if (gameInterval < 250) {
			currentGameIntervalInMs = gameInterval
			gameIntervalChannel <- currentGameIntervalInMs
		}
	}
	if (win.JustPressed(pixelgl.KeyKPSubtract)) {
		gameInterval := currentGameIntervalInMs - 10
		if (gameInterval > 0) {
			currentGameIntervalInMs = gameInterval
			gameIntervalChannel <- currentGameIntervalInMs
		}
	}
}

func addPatternToGrid(patternType patternType, position pixel.Vec, grid *grid) {
	patternOffsets := getPatternOffsets(patternType)
	cellXIndex, cellYIndex := getCellIndiciesByMousePosition(position)

	for _, offset := range patternOffsets {
		offsetedX := cellXIndex + offset.x
		offsetedY := cellYIndex + offset.y
		
		if offsetedX < 0 || offsetedX >= cellsHorizontalCount || offsetedY < 0 || offsetedY >= cellsVerticalCount {
			return
		}

		grid.cells[offsetedX][offsetedY].isAlive = true
	}
}

func initGrid() grid {
	grid := grid{}

	for i := 0; i < cellsHorizontalCount; i++ {
		for j := 0; j < cellsVerticalCount; j++ {
			cellPositionX := float64(i * cellSize) + cellMargin
			cellPositionY := float64(j * cellSize) + cellMargin

			cell := cell{false, cellPositionX, cellPositionY}
			grid.cells[i][j] = cell
		}
	}

	return grid
}

func drawGrid(win *pixelgl.Window, imd *imdraw.IMDraw, grid *grid) {
	imd.Clear()

	for i := 0; i < cellsHorizontalCount; i++ {
		for j := 0; j < cellsVerticalCount; j++ {
			cell := grid.cells[i][j]

			if cell.isAlive {
				imd.Color = colornames.Black
			} else {
				imd.Color = colornames.Skyblue
			}

			imd.Push(
				pixel.V(cell.x, cell.y),
				pixel.V(cell.x + cellSize - cellMargin, cell.y + cellSize - cellMargin))
			imd.Rectangle(0)
		}
	}

	imd.Draw(win)
}

func updateGridState(grid *grid) {
	// Cloning grid to temp grid variable
	tempGrid := *grid

	for i := 0; i < cellsHorizontalCount; i++ {
		for j := 0; j < cellsVerticalCount; j++ {
			aliveNeighboursCount := getAliveNeighbourCellsCount(grid, i, j)

			if grid.cells[i][j].isAlive {
				if aliveNeighboursCount < 2 || aliveNeighboursCount > 3 {
					tempGrid.cells[i][j].isAlive = false
				}
			} else {
				if aliveNeighboursCount == 3 {
					tempGrid.cells[i][j].isAlive = true
				}
			}
		}
	}

	for i := 0; i < cellsHorizontalCount; i++ {
		for j := 0; j < cellsVerticalCount; j++ {
			grid.cells[i][j].isAlive = tempGrid.cells[i][j].isAlive
		}
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	grid := initGrid()
	instructionsText := initInstructionsText()
	gameStatusText := initGameStatusText()
	gameIntervalText := initGameIntervalText()
	gameIntervalChannel := make(chan int)
	quitTickerChannel := make(chan bool)
	
    go func() {
		ticker := time.NewTicker(time.Duration(currentGameIntervalInMs) * time.Millisecond)

        for {
			select {
			case interval := <-gameIntervalChannel:

				fmt.Println(interval)

				ticker.Stop()
				ticker = time.NewTicker(time.Duration(interval) * time.Millisecond)
			default:
			}

            select {
			case <-ticker.C:
				if !isGameStarted { continue }
				updateGridState(&grid)
			case <-quitTickerChannel:
				ticker.Stop()
                return
			}
        }
    }()

	for !win.Closed() {
		win.Clear(colornames.Black)
		
		handleMouseEvents(win, &grid)
		handleKeyboardEvents(win, &grid, gameIntervalChannel)

		instructionsText.Draw(win, pixel.IM)
		drawGameIntervalText(win, gameIntervalText)
		drawGameStatusText(win, gameStatusText)
		drawGrid(win, imd, &grid)

		win.Update()
	}
}

func main() {
	fmt.Println("Game of Life running...")

	pixelgl.Run(run)
}
