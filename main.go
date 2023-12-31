package main

import "fmt"

func main() {

	// grid := conway.NewGrid()

	// grid.DisplayGrid()
	// grid.Update()
	// simpleTestCase()
}

func simpleTestCase() {
	alive := true
	dead := false
	for i := 0; i < 10; i++ {
		alive = i == 2 || i == 3
		dead = i == 3

		if alive {
			fmt.Printf("Dies[%d]: %v\n", i, alive)
		}

		if dead {
			fmt.Printf("Birth[%d]: %v\n", i, dead)
		}
	}

}
