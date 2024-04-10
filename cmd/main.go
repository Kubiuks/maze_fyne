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

	maze := tools.CreateMaze(width, height, scale)

	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("Maze")

	// lines for layout
	rightLine := canvas.NewLine(color.Black)
	rightLine.StrokeWidth = 2
	bottomLine := canvas.NewLine(color.Black)
	bottomLine.StrokeWidth = 2

	lines := make([]fyne.CanvasObject, width*height*2+2)
	lines[0] = rightLine
	lines[1] = bottomLine

	for i := 0; i < width*height*2; i++ {
		line := canvas.NewLine(color.Black)
		line.StrokeWidth = 2
		lines[i+2] = line
	}

	// agents
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	agentObj := canvas.NewRectangle(green)
	agent := tools.NewAgent(agentObj, tools.Coordinates{X: 0, Y: 0}, &maze)

	allObjs := append(lines, agentObj)
	content := container.New(NewMazeLayout(lines, &maze, agentObj, scale), allObjs...)

	myWindow.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		var win bool
		switch k.Name {
		case fyne.KeyA:
			// left
			win = agent.MoveAgent(-1, 0)
		case fyne.KeyW:
			// up
			win = agent.MoveAgent(0, -1)
		case fyne.KeyD:
			// right
			win = agent.MoveAgent(1, 0)
		case fyne.KeyS:
			// down
			win = agent.MoveAgent(0, 1)
		}
		if win {
			fmt.Println("YOU WON")
		}
	})

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(float32(width*scale+30), float32(height*scale+30)))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
