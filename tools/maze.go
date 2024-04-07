package tools

import (
	"math/rand"
)

type Square struct {
	Agent    bool
	Visited  bool
	Location Coordinates
	Walls    [4]bool
}

// possible directions
var directions = [4]bool{false, false, false, false} //{left, up, right, down}

// Maze structure
type Maze struct {
	Width, Height int
	Grid          [][]Square
	StartPoint    Coordinates
	EndPoint      Coordinates
	Solved        bool
}

// Point on the maze
type Coordinates struct {
	X, Y int
}

func CreateMaze(w, h int) Maze {
	maze := Maze{Width: w, Height: h}
	maze.Grid = make([][]Square, h)
	maze.StartPoint = Coordinates{0, 0}
	maze.EndPoint = Coordinates{h - 1, w - 1}
	maze.Solved = false
	for i := 0; i < w; i++ {
		maze.Grid[i] = make([]Square, w)
		for j := 0; j < h; j++ {
			maze.Grid[i][j].Location = Coordinates{i, j}
			maze.Grid[i][j].Walls = [4]bool{true, true, true, true}
		}
	}
	maze.generate()
	return maze
}

func (m *Maze) generate() {

	stack := NewStack()

	stack.Push(m.StartPoint)
	nVisitedSquares := 1

	for nVisitedSquares < m.Height*m.Width {

		// 1. Create set of unvisted Neighbours
		co := stack.Peek().(Coordinates)
		m.Grid[co.X][co.Y].Visited = true
		var unvisitedNeighbours []int

		// Left neighbour
		if co.X > 0 && !m.Grid[co.X-1][co.Y].Visited {
			unvisitedNeighbours = append(unvisitedNeighbours, 0)
		}

		// Top neighbour
		if co.Y > 0 && !m.Grid[co.X][co.Y-1].Visited {
			unvisitedNeighbours = append(unvisitedNeighbours, 1)
		}

		// Right neighbour
		if co.X < m.Width-1 && !m.Grid[co.X+1][co.Y].Visited {
			unvisitedNeighbours = append(unvisitedNeighbours, 2)
		}

		// Down neighbour
		if co.Y < m.Height-1 && !m.Grid[co.X][co.Y+1].Visited {
			unvisitedNeighbours = append(unvisitedNeighbours, 3)
		}

		if len(unvisitedNeighbours) == 0 {
			// no neighbours, need to backtrack
			stack.Pop()
		} else {
			// choose random neigbour
			nextCellDir := unvisitedNeighbours[rand.Intn(len(unvisitedNeighbours))]

			// delete Wall between curr and next and push next cell to stack
			switch nextCellDir {
			case 0:
				m.Grid[co.X][co.Y].Walls[0] = false
				m.Grid[co.X-1][co.Y].Walls[2] = false
				stack.Push(Coordinates{co.X - 1, co.Y})
			case 1:
				m.Grid[co.X][co.Y].Walls[1] = false
				m.Grid[co.X][co.Y-1].Walls[3] = false
				stack.Push(Coordinates{co.X, co.Y - 1})
			case 2:
				m.Grid[co.X][co.Y].Walls[2] = false
				m.Grid[co.X+1][co.Y].Walls[0] = false
				stack.Push(Coordinates{co.X + 1, co.Y})
			case 3:
				m.Grid[co.X][co.Y].Walls[3] = false
				m.Grid[co.X][co.Y+1].Walls[1] = false
				stack.Push(Coordinates{co.X, co.Y + 1})
			}
			nVisitedSquares++
		}

	}

}
