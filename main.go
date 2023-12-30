package main

import (
	"fmt"

	"com.jacobsaucelo.conway/conway"
)

func main() {
	fmt.Println(conway.NewGrid())

	grid := conway.NewGrid()

	grid.DisplayGrid()
}
