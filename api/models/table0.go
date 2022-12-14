package models

import (
	"os"
	"io"
	"fmt"
	"log"
	"time"
	"errors"
	"crypto/sha256"
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


func createTable0(tables *tables) {
	fmt.Println("Generating pruning table for phase 0")

	cube := &Cube{}
	cube.initialize()
	parents := []Cube{
		*cube,
	}

	for depth := 1; depth < 7; depth++ {
		var children []Cube
		fmt.Println("table0: depth =", depth)
		for _, parent := range parents {
			for _, move := range GetGroupMoves(0, parent) {
				child := copyAndMove(&parent, move)
				index := edgeOriToInt(child.edgeOri)
				
				if index != 0 && tables.Table0[index] == 0 {
					tables.Table0[index] = uint8(depth)
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	for i, depth := range tables.Table0 {
		if i > 0 && depth == 0 {
			tables.Table0[i] = 7
		}
	}
}

func setTable0(tables *tables) {
	if _, err := os.Stat("pruningTables"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("pruningTables", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat("pruningTables/Table0"); os.IsNotExist(err) {
		start := time.Now()
		createTable0(tables)
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
		f, err := os.Open("pruningTables/Table0")
		if err != nil {
		  log.Fatal(err)
		}
		defer f.Close()
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		strSha256 := fmt.Sprintf("%x", h.Sum(nil))
		if strSha256 != "adc07ef872db78f680433ac7eb7b5807c431530b4380696785b938f0c17581ca" {
			log.Fatal("Table0 has wrong sha256 hash. Please delete it and launch cmd again.")
		}

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
