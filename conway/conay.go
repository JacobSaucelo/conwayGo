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
				fmt.Print("▉▉")
			} else {
				fmt.Print("░░")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

/*
	0 0 0
	0 0 0
	0 0 0
*/

func (g *GridContainer) NeighborCount(row int, col int) int {
	count := 0

	// row = 0
	// col = 0

	// i = -1
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < int(rows) && j >= 0 && j < int(cols) && !(i == row && j == col) {
				fmt.Print(g.Grid[i][j])
				fmt.Printf("(%d,%d)", i, j)
				// working dont get center
				if g.Grid[i][j].isAlive {
					count++
				}
			}
		}
		fmt.Println()
	}

	// fmt.Println("you picked", g.Grid[row][col], "val:", row, " : ", col)
	return count
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
