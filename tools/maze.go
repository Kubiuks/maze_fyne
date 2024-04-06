package tools

const (
	up = 0
	down
	left
	right
)

// possible directions
var directions = []int{up, down, left, right}

// Maze structure
type Maze struct {
	Width, Height int
	Grid          [][]int
	StartPoint    *Point
	EndPoint      *Point
	Solved        bool
}

// Point on the maze
type Point struct {
	X, Y int
}

// Are the two points the same
func (point_1 *Point) Equal(point_2 *Point) bool {
	return point_1.X == point_2.X && point_1.Y == point_2.Y
}

func CreateMaze(w, h int) Maze {
	maze := Maze{Width: w, Height: h}
	maze.Grid = make([][]int, h)
	for i := 0; i < w; i++ {
		maze.Grid[i] = make([]int, w)
		for j := 0; j < h; j++ {
			maze.Grid[i][j] = 1
		}
	}

	return maze
}
