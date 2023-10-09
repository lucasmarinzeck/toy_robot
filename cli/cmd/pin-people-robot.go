package main

import (
	"fmt"
	"os"
	"os/signal"
	"pin-people-robot/internal/domain/direction"
	"pin-people-robot/internal/domain/robot"
	"syscall"
)

const (
	Move = iota + 1
	Left
	Right
	Report
)

const (
	North = iota + 1
	East
	South
	West
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	gridWidth := 5
	gridHeight := 5

	var initialX, initialY, initialFacing int
	var face direction.Direction

	fmt.Println("--- PLACE BOT INITIAL POSITION ---")
	fmt.Println("X value:")
	fmt.Scan(&initialX)

	if initialX >= gridWidth || initialX < 0 {
		panic("Invalid X value, out of bounds")
	}

	fmt.Println("Y value:")
	fmt.Scan(&initialY)

	if initialY >= gridHeight || initialY < 0 {
		panic("Invalid Y value, out of bounds")
	}

	fmt.Printf("\nSelect your facing:\n")
	fmt.Println("(1)North")
	fmt.Println("(2)East")
	fmt.Println("(3)South")
	fmt.Println("(4)West")
	fmt.Scan(&initialFacing)

	switch initialFacing {
	case North:
		face = direction.North
	case East:
		face = direction.East
	case South:
		face = direction.South
	case West:
		face = direction.West
	}

	r := robot.New(initialX, initialY, face, gridWidth, gridHeight)

	r.DrawGrid()

	for {
		var command int

		fmt.Printf("\n")
		fmt.Println("--- SELECT A COMMAND ---")
		fmt.Println("(1) MOVE")
		fmt.Println("(2) LEFT")
		fmt.Println("(3) RIGHT")
		fmt.Println("(4) Report")
		fmt.Scan(&command)

		switch command {
		case Move:
			r.Move()
		case Left:
			r.Turn("left")
		case Right:
			r.Turn("right")
		case Report:
			r.Report()
		default:
			fmt.Printf("\nInvalid input\n")
		}

	}
}
