package robot

import (
	"fmt"
	"pin-people-robot/internal/domain/direction"
)

type Robot struct {
	X, Y                  int
	F                     direction.Direction
	icon                  string
	gridWidth, gridHeight int
}

func New(x, y int, initialFacing direction.Direction, gridWidth, gridHeight int) *Robot {
	icon := currentIcon(initialFacing)
	return &Robot{X: x, Y: y, F: initialFacing, icon: icon, gridWidth: gridWidth, gridHeight: gridHeight}
}

func currentIcon(f direction.Direction) string {
	switch f {
	case direction.West:
		return "<"
	case direction.East:
		return ">"
	case direction.South:
		return "v"
	default:
		return "^"
	}
}

func (r *Robot) Turn(turn string) direction.Direction {
	fmt.Printf("\nTurning face to %s \n\n", turn)
	currentFacing := r.F
	newFacing := currentFacing

	switch turn {
	case "left":
		switch currentFacing {
		case direction.South:
			newFacing = direction.East
		case direction.North:
			newFacing = direction.West
		case direction.West:
			newFacing = direction.South
		case direction.East:
			newFacing = direction.North

		}
	case "right":
		switch currentFacing {
		case direction.South:
			newFacing = direction.West
		case direction.North:
			newFacing = direction.East
		case direction.West:
			newFacing = direction.North
		case direction.East:
			newFacing = direction.South
		}
	}

	r.F = newFacing
	r.icon = currentIcon(r.F)
	return newFacing
}

func (r *Robot) Move() {
	switch r.F {
	case direction.North:
		if r.Y == 0 {
			fmt.Printf("\033[31m INVALID MOVE: out of grid bounds \033[0m\n")
		} else {
			r.Y = r.Y - 1
		}
	case direction.East:
		if r.X == r.gridWidth-1 {
			fmt.Printf("\033[31m INVALID MOVE: out of grid bounds \033[0m\n")
		} else {
			r.X = r.X + 1
		}
	case direction.West:
		if r.X == 0 {
			fmt.Printf("\033[31m INVALID MOVE: out of grid bounds \033[0m\n")
		} else {
			r.X = r.X - 1
		}
	case direction.South:
		if r.Y == r.gridHeight-1 {
			fmt.Printf("\n\033[31m INVALID MOVE: out of grid bounds \033[0m\n")
		} else {
			r.Y = r.Y + 1
		}
	}
}

func (r *Robot) DrawGrid() {
	fmt.Printf("\n")
	for i := 0; i < r.gridHeight; i++ {
		// Draws top line
		for j := 0; j < r.gridWidth; j++ {
			fmt.Printf("+-----")
		}
		fmt.Println("+")

		// Draws inner content cells
		for j := 0; j < r.gridWidth; j++ {
			// Draws the arrow that indicates position and facing
			if i == r.Y && j == r.X {
				fmt.Printf("| \033[31m %s  \033[0m", r.icon)
			} else {
				// Draws coordinates
				fmt.Printf("| %d,%d ", j, i)
			}
		}
		fmt.Println("|")
	}

	// Draws bottom line
	for j := 0; j < r.gridWidth; j++ {
		fmt.Printf("+-----")
	}
	fmt.Printf("+\n\n")
}

func (r *Robot) Report() {
	fmt.Printf("\nReporting...\n")
	fmt.Printf("Current position: (%d, %d). Facing %s\n\n", r.X, r.Y, r.F)
	r.DrawGrid()
}
