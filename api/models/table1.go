package models

import (
	"os"
	"io"
	"fmt"
	"log"
	"time"
	"errors"
	"strconv"
	"crypto/sha256"
)


/*
** edge pos of the middle ones to int
** last one not needed because determined by the 11 other pos
*/
func edgePosToInt(edgePos [12]uint8) int {
	decimal := 0
	for i := 0; i < 11; i++ {
		decimal *= 2
		if edgePos[i] > uint8(7) {
			decimal += 1
		}
	}
	return decimal
}


/*
** corner orientation to int
** 8th one is determined by all the other so no need to keep it hence len - 1
*/
func cornerOriToInt(cornerOri [8]byte) int {
	decimal := 0
	for _, val := range cornerOri[:len(cornerOri) - 1] {
		decimal *= 3
		decimal += int(val)
	}
	return decimal
}


func createTable1(tables *tables) {
	fmt.Println("Generating pruning table for phase 1")

	cube := &Cube{}
	cube.initialize()
	parents := []Cube{
		*cube,
	}

	for depth := 1; depth < 9; depth++ {
		var children []Cube

		fmt.Println("table1: depth =", depth)
		for _, parent := range parents {
			for _, move := range GetGroupMoves(1, parent) {
				child := copyAndMove(&parent, move)

				indexEdgePos := tables.T1EdgePosIndex[edgePosToInt(child.edgePos)]
				indexCornerOri := cornerOriToInt(child.cornerOri)
				if ((indexEdgePos > 0 || indexCornerOri > 0) && tables.Table1[indexEdgePos][indexCornerOri] == 0) {
					tables.Table1[indexEdgePos][indexCornerOri] = uint8(depth)
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
	for i := 0; i < 495; i++ {
		for j := 0; j < 2187; j++ {
			if ((i > 0 || j > 0) && tables.Table1[i][j] == 0) {
				tables.Table1[i][j] = 9
			}
		}
	}
}

/*
** int to count of activated bits in its string
*/
func convIntToSumOf1(nb uint16) uint8 {
	var numOf1 uint8
	binaryRepr := strconv.FormatInt(int64(nb), 2)
	for _, bit := range binaryRepr {
		if bit == '1' {
			numOf1++
		}
	}
	return numOf1
}

/*
** Since edgePosToInt doesn't return all numbers between 0 and 2083
** we need an array used as a map to both:
** - have a clear double table with correct dimensions
** - use 4 times less space for this step
*/
func setT1EdgePosConv(tables *tables) {
	var actualIndexedgePos uint16
	var i uint16
	for i = 0; i < 2048; i++ {
		numOf1 := convIntToSumOf1(i)
		if numOf1 == 3 || numOf1 == 4 { // if in condition of edge position
			tables.T1EdgePosIndex[i] = actualIndexedgePos
			actualIndexedgePos += 1
		}
	}
}


func setTable1(tables *tables) {
	setT1EdgePosConv(tables)
	
	if _, err := os.Stat("pruningTables"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("pruningTables", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat("pruningTables/Table1"); os.IsNotExist(err) {
		start := time.Now()
		createTable1(tables)
		fmt.Println("Took ", time.Since(start))

		file, err := os.Create("pruningTables/Table1")
		if err != nil {
			fmt.Println("Error: file creation: Table1")
			os.Exit(1)
		}
		defer file.Close()

		for i := 0; i < 495; i++ {
			for j := 0; j < 2187; j++ {
				_, err := file.WriteString(fmt.Sprintf("%d", tables.Table1[i][j]))
				if err != nil {
					fmt.Println("Error: file write: Table1")
					os.Exit(1)
				}
			}
		}
	} else {
		f, err := os.Open("pruningTables/Table1")
		if err != nil {
		  log.Fatal(err)
		}
		defer f.Close()
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		strSha256 := fmt.Sprintf("%x", h.Sum(nil))
		if strSha256 != "99f73043a1006dd2d09fd2d793e439447666c1a03a213f6ce06eaa8adaf00841" {
			log.Fatal("Table0 has wrong sha256 hash. Please delete it and launch cmd again.")
		}

		content, err := os.ReadFile("pruningTables/Table1")
		if err != nil {
			fmt.Println("failed to read pruning table file")
			os.Exit(1)
		}

		iEPos := 0
		iCOri := 0
		for _, depth := range content {
			tables.Table1[iEPos][iCOri] = depth - '0'
			iCOri++
			if iCOri >= 2187 {
				iCOri = 0
				iEPos++
			}
		}
	}
}
