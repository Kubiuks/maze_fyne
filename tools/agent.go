package tools

import (
	"fyne.io/fyne/v2"
)

// possible directions
// var directions = [4]bool{false, false, false, false} //{left, up, right, down}

type Agent struct {
	Agent    fyne.CanvasObject
	Position Coordinates
	maze     *Maze
}

func NewAgent(agent fyne.CanvasObject, startingPosition Coordinates, maze *Maze) *Agent {
	return &Agent{Agent: agent, Position: startingPosition, maze: maze}
}

func (a *Agent) MoveAgent(dx, dy int) bool {
	if a.allowedMove(dx, dy) {
		a.Agent.Move(fyne.NewPos(
			float32(int(a.Agent.Position().X)+dx*a.maze.drawScale),
			float32(int(a.Agent.Position().Y)+dy*a.maze.drawScale)))
		a.Agent.Refresh()
		a.Position.X = a.Position.X + dx
		a.Position.Y = a.Position.Y + dy
	}
	if a.Position == a.maze.EndPoint {
		return true
	}
	return false
}

func (a *Agent) allowedMove(dx, dy int) bool {
	if a.maze.inBounds(a.Position, dx, dy) && !a.maze.wallHit(a.Position, dx, dy) {
		return true
	}
	return false
}
