package main

import (
	"fmt"
	"maze/tools"
)

func main() {
	width, height := 30, 30

	maze := tools.CreateMaze(width, height)

	fmt.Printf("Maze of dimensions: %d, %d\n", maze.Width, maze.Height)
}
