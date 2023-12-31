package main

import (
	"time"

	"com.jacobsaucelo.conway/conway"
)

func main() {

	grid := conway.NewGrid()

	for gen := 0; gen < 100; gen++ {
		conway.ClearScreen()
		grid.DisplayGrid()
		grid = grid.Update()
		time.Sleep(10 * time.Millisecond)

	}

	// simpleTestCase()
}

// func simpleTestCase() {
// 	alive := true
// 	dead := false
// 	for i := 0; i < 10; i++ {
// 		alive = i == 2 || i == 3
// 		dead = i == 3

// 		if alive {
// 			fmt.Printf("Dies[%d]: %v\n", i, alive)
// 		}

// 		if dead {
// 			fmt.Printf("Birth[%d]: %v\n", i, dead)
// 		}
// 	}

// }
