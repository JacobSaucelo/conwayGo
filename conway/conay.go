package conway

import (
	"fmt"
	"math/rand"
)

var rows uint8 = 10
var cols uint8 = 10

type Cell struct {
	isAlive bool
}

type Grid [][]Cell

type GridContainer struct {
	Grid Grid
}

func NewGrid() GridContainer {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]Cell, cols)
		for j := range grid[i] {
			grid[i][j] = Cell{isAlive: rand.Intn(2) == 1}
		}
	}

	return GridContainer{
		Grid: grid,
	}
}

/*
	[][][][]
	[][][][]
	[][][][]
*/

func (g *GridContainer) DisplayGrid() {
	for i := range g.Grid {
		for j := range g.Grid[i] {
			if g.Grid[i][j].isAlive {
				fmt.Print("1")
				// fmt.Print(g.Grid[i][j])
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
