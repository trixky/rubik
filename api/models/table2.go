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


func edgePosToIntT2(edgePos [12]uint8) int {
	decimal := 0
	for i := 0; i < 8; i++ {
		decimal *= 2
		if edgePos[i] > uint8(3) {
			decimal += 1
		}
	}
	return decimal
}


func cPosToInt(cornerPos [8]uint8) int {
	decimal := 0
	for i := 0; i < len(cornerPos); i++ {
		decimal  *= len(cornerPos) - i
		for j := i + 1; j < len(cornerPos); j++ {
			if cornerPos[i] > cornerPos[j] {
				decimal += 1
			}
		}
	}
	return decimal
}


func cornersInOrbit(c *Cube) bool {
	for i := 0; i < 4; i++ {
		if c.cornerPos[i] > 3 {
			return false
		}
	}
	return true
}


func cornerPosInList(c *Cube, list []Cube) bool {
	for _, cube := range list {
		if cPosToInt(c.cornerPos) == cPosToInt(cube.cornerPos) {
			return true
		}
	}
	return false
}

func init96Cubes() []Cube {
	cube := &Cube{}
	cube.initialize()

	ninetySixCubes := []Cube{
		*cube,
	}
	parents := []Cube{
		*cube,
	}
	for depth := 0; depth < 4; depth++ {
		var children []Cube
		fmt.Println("depth", depth)
		for _, parent := range parents {
			for _, move := range GetGroupMoves(3, parent) {
				child := copyAndMove(&parent, move)
				if (cornersInOrbit(child) == true && cornerPosInList(child, ninetySixCubes) == false) {
					// child.Print()
					// fmt.Println()
					ninetySixCubes = append(ninetySixCubes, *child)
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	return ninetySixCubes
}

func createTable2(tables *tables) {
	fmt.Println("Generating pruning table for phase 2")

	cube := &Cube{}
	cube.initialize()
	parents := init96Cubes()

	fmt.Println(GetGroupMoves(2, parents[0]))
	for depth := 1; depth < 15; depth++ {
		var children []Cube
		for _, parent := range parents {
			for _, move := range GetGroupMoves(2, parent) {
				child := copyAndMove(&parent, move)
				indexCornerPos := cPosToInt(child.cornerPos)
				indexEdgePos := tables.T2EdgePosIndex[edgePosToIntT2(child.edgePos)]
				if ((indexCornerPos > 0 || indexEdgePos > 0) && tables.Table2[indexCornerPos][indexEdgePos] == 0) {
					tables.Table2[indexCornerPos][indexEdgePos] = uint8(depth)
					children = append(children, *child)
				}
			}
		}
		parents = children
		fmt.Println("table2: depth =", depth, "done")
	}
}

func hexToInt(char byte) uint8 {
	// fmt.Println("hexToInt ", char)
	if char >= 'a' {
		return char - 'a' + 10
	}
	return char - '0'
}

func setT2EdgePosConv(tables *tables) {
	var validIndiceOfEdgePos uint8
	var i uint16
	for i = 0; i < 255; i++ {
		numOf1 := convIntToSumOf1(i)
		if numOf1 == 4 { // if in condition of edge position
			tables.T2EdgePosIndex[i] = validIndiceOfEdgePos
			validIndiceOfEdgePos += 1
		}
	}
}

func setTable2(tables *tables) {
	setT2EdgePosConv(tables)

	if _, err := os.Stat("pruningTables"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("pruningTables", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat("pruningTables/Table2"); os.IsNotExist(err) {
		start := time.Now()
		createTable2(tables)
		fmt.Println("Took ", time.Since(start))

		file, err := os.Create("pruningTables/Table2")
		if err != nil {
			fmt.Println("Error: file creation: Table2")
			os.Exit(1)
		}
		defer file.Close()

		fmt.Println("writing to file")
		for i := 0; i < 40320; i++ {
			for j := 0; j < 70; j++ {
				_, err := file.WriteString(fmt.Sprintf("%x", tables.Table2[i][j]))
				if err != nil {
					fmt.Println("Error: file write: Table2")
					os.Exit(1)
				}
			}
		}
	} else {
		f, err := os.Open("pruningTables/Table2")
		if err != nil {
		  log.Fatal(err)
		}
		defer f.Close()
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		strSha256 := fmt.Sprintf("%x", h.Sum(nil))
		if strSha256 != "72b204274fed87660b5647c121325f0f97e285d7cc0e12ab6a0906a1a2a9b1c7" {
			log.Fatal("Table2 has wrong sha256 hash. Please delete it and launch cmd again.")
		}

		content, err := os.ReadFile("pruningTables/Table2")
		if err != nil {
			fmt.Println("failed to read pruning table file")
			os.Exit(1)
		}

		iCPos := 0
		iEPos := 0
		for _, depth := range content {
			tables.Table2[iCPos][iEPos] = hexToInt(depth)
			iEPos++
			if iEPos >= 70 {
				iEPos = 0
				iCPos++
			}
		}
	}
}
