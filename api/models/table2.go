package models

// import (
// 	"os"
// 	"fmt"
// 	"log"
// 	"time"
// 	"errors"
// 	"strconv"
// )


// // group 1 -> group 2 : edge placement
// func edgePosToInt(edgePos [12]uint8) int {
// 	decimal := 0
// 	for i := 0; i < 11; i++ {
// 		decimal *= 2
// 		if edgePos[i] > uint8(7) {
// 			decimal += 1
// 		}
// 	}
// 	return decimal
// }


// // group 1 -> group 2 : corner orientation
// func cornerOriToInt(cornerOri [8]byte) int {
// 	decimal := 0
// 	for _, val := range cornerOri[:len(cornerOri) - 1] {
// 		decimal *= 3
// 		decimal += int(val)
// 	}
// 	return decimal
// }


// func createTable2(tables *tables) {
// 	fmt.Println("Generating pruning table for G1")

// 	cube := &Cube{}
// 	cube.initialize()
// 	parents := []Cube{
// 		*cube,
// 	}

// 	depth  := 0
// 	for depth < 8 {
// 		var children []Cube
// 		depth++
// 		for _, parent := range parents {
// 			for _, move := range getGroupMoves(1, parent) {
// 				child := copyAndMove(&parent, move)
// 				indexEdgePos := tables.T2EdgePosIndex[edgePosToInt(child.edgePos)]
// 				indexCornerOri := cornerOriToInt(child.cornerOri)
// 				if (indexEdgePos > 0 && indexCornerOri > 0 && tables.Table2[indexEdgePos][indexCornerOri] == 0) {
// 					tables.Table2[indexEdgePos][indexCornerOri] = uint8(depth)
// 				}
// 				children = append(children, *child)
// 			}
// 		}
// 		parents = children
// 		fmt.Println("table2: depth =", depth, "done")
// 	}
// 	fmt.Println("setting nines:")
// 	for i := 0; i < 495; i++ {
// 		for j := 0; j < 2187; i++ {
// 			if (i > 0 && j > 0 && tables.Table2[i][j] == 0) {
// 				tables.Table2[i][j] = 9
// 			}
// 		}
// 	}
// 	fmt.Println("setting nines ok")
// 	fmt.Println()
// }

// func convIntToSumOf1(nb uint16) uint8 {
// 	var numOf1 uint8
// 	binaryRepr := strconv.FormatInt(int64(nb), 2)
// 	for _, bit := range binaryRepr {
// 		if bit == '1' {
// 			numOf1++
// 		}
// 	}
// 	return numOf1
// }

// func setT2EdgePosConv(tables *tables) {
// 	var validIndiceOfEdgePos uint16
// 	var i uint16
// 	for i = 0; i < 2048; i++ {
// 		numOf1 := convIntToSumOf1(i)
// 		if numOf1 == 3 || numOf1 == 4 { // if in condition of edge position
// 			tables.T2EdgePosIndex[i] = validIndiceOfEdgePos
// 			validIndiceOfEdgePos += 1
// 		}
// 	}
// 	fmt.Println("nb valid indices edge pos = ", validIndiceOfEdgePos)
// }

// func setTable2(tables *tables) {
// 	setT2EdgePosConv(tables)
// 	fmt.Println("Getting table for phase 1 (G1 -> G2)")

// 	if _, err := os.Stat("pruningTables"); errors.Is(err, os.ErrNotExist) {
// 		err := os.Mkdir("pruningTables", os.ModePerm)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	if _, err := os.Stat("pruningTables/Table2"); os.IsNotExist(err) {
// 		fmt.Println("Creating table for phase 1")

// 		start := time.Now()
// 		createTable2(tables)
// 		fmt.Println("Took ", time.Since(start))

// 		file, err := os.Create("pruningTables/Table2")
// 		if err != nil {
// 			fmt.Println("Error: file creation: Table2")
// 			os.Exit(1)
// 		}
// 		defer file.Close()

// 		fmt.Println("writing to file")
// 		for i := 0; i < 495; i++ {
// 			fmt.Print(i, " ")
// 			for j := 0; j < 2187; j++ {
// 				_, err := file.WriteString(fmt.Sprintf("%d", tables.Table2[i][j]))
// 				if err != nil {
// 					fmt.Println("Error: file write: Table2")
// 					os.Exit(1)
// 				}
// 			}
// 		}
// 	} else {
// 		fmt.Println("Reading and setting phase 1 map")
// 		content, err := os.ReadFile("pruningTables/Table2")
// 		if err != nil {
// 			fmt.Println("failed to read pruning table file")
// 			os.Exit(1)
// 		}

// 		iEPos := 0
// 		iCOri := 0
// 		for _, depth := range content {
// 			tables.Table2[iEPos][iCOri] = depth - '0'
// 			iCOri++
// 			if iCOri >= 2187 {
// 				iCOri = 0
// 				iEPos++
// 			}
// 		}
// 	}
// }
