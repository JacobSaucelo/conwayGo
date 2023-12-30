package main

import (
	"fmt"

	"com.jacobsaucelo.conway/conway"
)

func main() {

	grid := conway.NewGrid()

	// grid.DisplayGrid()
	fmt.Println(grid.NeighborCount(3, 3))

}
