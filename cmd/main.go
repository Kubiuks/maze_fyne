package main

import (
	"fmt"
	"image/color"

	"github.com/Kubiuks/maze_fyne/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	width, height := 30, 30
	scale := 20

	maze := tools.CreateMaze(width, height)

	fmt.Printf("Maze of dimensions: %d, %d\n", maze.Width, maze.Height)

	myApp := app.New()

	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("Maze")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	rightLine := canvas.NewLine(color.Black)
	bottomLine := canvas.NewLine(color.Black)
	agent := canvas.NewRectangle(green)
	bottomLine.StrokeWidth = 2
	rightLine.StrokeWidth = 2

	lines := []fyne.CanvasObject{rightLine, bottomLine}

	for i := 0; i < width*height*2; i++ {
		line := canvas.NewLine(color.Black)
		line.StrokeWidth = 2
		lines = append(lines, line)
	}

	allObjs := append(lines, agent)
	content := container.New(NewMazeLayout(lines, &maze, agent, scale), allObjs...)

	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(float32(width*scale+30), float32(height*scale+30)))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
