package main

import (
	"log"

	"fyne.io/fyne/v2"
	"github.com/Kubiuks/maze_fyne/tools"
)

type mazeLayout struct {
	walls []fyne.CanvasObject
	maze  *tools.Maze
	agent fyne.CanvasObject
	scale int
}

func NewMazeLayout(walls []fyne.CanvasObject, maze *tools.Maze, agent fyne.CanvasObject, scale int) fyne.Layout {
	return &mazeLayout{walls: walls, maze: maze, agent: agent, scale: scale}
}

func (m *mazeLayout) Layout(objets []fyne.CanvasObject, size fyne.Size) {
	log.Printf("%v %v\n", size.Width, size.Height)
	// right line
	m.walls[0].Move(fyne.NewPos(float32(m.maze.Width*m.scale+10), 10))
	m.walls[0].Resize(fyne.NewSize(0, float32(m.maze.Height*m.scale)))
	// bottom line
	m.walls[1].Move(fyne.NewPos(10, float32(m.maze.Height*m.scale+10)))
	m.walls[1].Resize(fyne.NewSize(float32(m.maze.Width*m.scale), 0))

	i := 2
	for _, row := range m.maze.Grid {
		for _, square := range row {
			if square.Walls[0] {
				// left line of square
				m.walls[i].Move(fyne.NewPos(float32(square.Location.X*m.scale+10), float32(square.Location.Y*m.scale+10)))
				m.walls[i].Resize(fyne.NewSize(0, float32(m.scale)))
				i++
			}
			if square.Walls[1] {
				// top line of square
				m.walls[i].Move(fyne.NewPos(float32(square.Location.X*m.scale+10), float32(square.Location.Y*m.scale+10)))
				m.walls[i].Resize(fyne.NewSize(float32(m.scale), 0))
				i++
			}
		}
	}

	m.agent.Resize(fyne.NewSize(float32(m.scale-6), float32(m.scale-6)))
	m.agent.Move(fyne.NewPos(13, 13))
}

func (m *mazeLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(100, 100)
}
