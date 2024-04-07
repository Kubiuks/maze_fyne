package tools

const (
	up = 0
	down
	left
	right
)

type Square struct {
	Agent    bool
	Empty    bool
	Location Coordinates
	Walls    [4]bool
}

// possible directions
var directions = []int{up, down, left, right}

// Maze structure
type Maze struct {
	Width, Height int
	Grid          [][]Square
	StartPoint    *Coordinates
	EndPoint      *Coordinates
	Solved        bool
}

// Point on the maze
type Coordinates struct {
	X, Y int
}

func CreateMaze(w, h int) Maze {
	maze := Maze{Width: w, Height: h}
	maze.Grid = make([][]Square, h)
	for i := 0; i < w; i++ {
		maze.Grid[i] = make([]Square, w)
		for j := 0; j < h; j++ {
			maze.Grid[i][j].Walls = [4]bool{true, true, true, true}
			maze.Grid[i][j].Location.X = i
			maze.Grid[i][j].Location.Y = j
		}
	}

	return maze
}
