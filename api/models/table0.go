package models

import (
	"os"
	"fmt"
	"log"
	"time"
	"errors"
)

// group 0 -> group 1 : edge orientation
func edgeOriToInt(edgeOri [12]byte) int {
	decimal := 0
	for _, val := range edgeOri[:len(edgeOri) - 1] {
		decimal *= 2
		decimal += int(val)
	}
	return decimal
}


func createTable0() [2048]uint8 {
	fmt.Println("Generating pruning table for G0")

	table0 := [2048]uint8{}
	cube := &Cube{}
	cube.initialize()
	parents := []Cube{
		*cube,
	}

	depth  := 0
	for depth < 6 {
		var children []Cube
		depth++
		fmt.Println("table0: depth =", depth)
		for _, parent := range parents {
			for _, move := range getGroupMoves(0, parent) {
				child := copyAndMove(&parent, move)
				index := edgeOriToInt(child.edgeOri)
				
				if index != 0 && table0[index] == 0 {
					table0[index] = uint8(depth)
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	for i, depth := range table0 {
		if i > 0 && depth == 0 {
			table0[i] = 7
		}
	}
	fmt.Printf("\n")
	return table0
}

func setTable0(tables *tables) {
	fmt.Println("Getting table for phase 0 (G0 -> G1)")

	if _, err := os.Stat("pruningTables"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("pruningTables", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat("pruningTables/Table0"); os.IsNotExist(err) {
		fmt.Println("Creating table for phase 0")
		
		start := time.Now()
		tables.Table0 = createTable0()
		fmt.Println("Took ", time.Since(start))

		file, err := os.Create("pruningTables/Table0")
		if err != nil {
			fmt.Println("Error: file creation: Table0")
			os.Exit(1)
		}
		defer file.Close()

		for i := 0; i < len(tables.Table0); i++ {
			_, err := file.WriteString(fmt.Sprintf("%d", tables.Table0[i]))
			if err != nil {
				fmt.Println("Error: file write: Table0")
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("Reading and setting phase 0 map")
		content, err := os.ReadFile("pruningTables/Table0")
		if err != nil {
			fmt.Println("failed to read pruning table file")
			os.Exit(1)
		}
		for i, depth := range content {
			tables.Table0[i] = depth - '0'
		}
	}
}
