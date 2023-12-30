package conway

import "math/rand"

var rows uint8 = 10
var cols uint8 = 10

type Cell struct {
	isAlive bool
}

type Grid [][]Cell

func NewGrid() Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]Cell, cols)
		for j := range grid[i] {
			grid[i][j] = Cell{isAlive: rand.Intn(2) == 1}
		}
	}

	return grid
}
