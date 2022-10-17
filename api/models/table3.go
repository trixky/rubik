package models

import (
	"os"
	"fmt"
	"log"
	"time"
	"errors"
)


func getSlice(sliceNum int, edgePos [12]uint8) uint8 {
	var slicePos uint8 = uint8(sliceNum * 4)
	if edgePos[0 + slicePos] == 0 + slicePos {
		if edgePos[1 + slicePos] == 1 + slicePos {
			if edgePos[2 + slicePos] == 2 + slicePos {
				return 0 // 0123
			} else {
				return 1 // 0132
			}
		} else if edgePos[1 + slicePos] == 2 + slicePos {
			if edgePos[2 + slicePos] == 1 + slicePos {
				return 2 // 0213
			} else {
				return 3 // 0231
			}
		} else {
			if edgePos[2 + slicePos] == 1 + slicePos {
				return 4 // 0312
			} else {
				return 5 // 0321
			}
		}
	} else if edgePos[0 + slicePos] == 1 + slicePos {
		if edgePos[1 + slicePos] == 0 + slicePos {
			if edgePos[2 + slicePos] == 2 + slicePos {
				return 6 // 1023
			} else {
				return 7 // 1032
			}
		} else if edgePos[1 + slicePos] == 2 + slicePos {
			if edgePos[2 + slicePos] == 0 + slicePos {
				return 8 // 1203
			} else {
				return 9 // 1230
			}
		} else {
			if edgePos[2 + slicePos] == 0 + slicePos {
				return 10 // 1302
			} else {
				return 11 // 1320
			}
		}
	} else if edgePos[0 + slicePos] == 2 + slicePos {
		if edgePos[1 + slicePos] == 0 + slicePos {
			if edgePos[2 + slicePos] == 1 + slicePos {
				return 12 // 2013
			} else {
				return 13 // 2031
			}
		} else if edgePos[1 + slicePos] == 1 + slicePos {
			if edgePos[2 + slicePos] == 0 + slicePos {
				return 14 // 2103
			} else {
				return 15 // 2130
			}
		} else {
			if edgePos[2 + slicePos] == 0 + slicePos {
				return 16 // 2301
			} else {
				return 17 // 2310
			}
		}
	} else {
		if edgePos[1 + slicePos] == 0 + slicePos {
			if edgePos[2 + slicePos] == 1 + slicePos {
				return 18 // 3012
			} else {
				return 19 // 3021
			}
		} else if edgePos[1 + slicePos] == 1 + slicePos {
			if edgePos[2 + slicePos] == 0 + slicePos {
				return 20 // 3102
			} else {
				return 21 // 3120
			}
		} else {
			if edgePos[2 + slicePos] == 0 + slicePos {
				return 22 // 3201
			} else {
				return 23 // 3210
			}
		}
	}
}

func edgePosConverter(edgePos [12]uint8) [3]uint8 {
	var sliceIndex [3]uint8
	for slice := 0; slice < 3; slice++ {
		sliceIndex[slice] = getSlice(slice, edgePos)
	}
	return sliceIndex
}

func setT3CornerPosIndex(tables *tables) {
	cube := &Cube{}
	cube.initialize()

	ninetySixCubes := []Cube{
		*cube,
	}
	parents := []Cube{
		*cube,
	}

	var indexT3CornerPos uint8 = 1
	for depth := 0; depth < 4; depth++ {
		var children []Cube
		for _, parent := range parents {
			for _, move := range GetGroupMoves(2, parent) {
				child := copyAndMove(&parent, move)
				if (cornersInOrbit(child) == true && cornerPosInList(child, ninetySixCubes) == false) {
					ninetySixCubes = append(ninetySixCubes, *child)
					tables.T3cornerPosIndex[cPosToInt(child.cornerPos)] = indexT3CornerPos
					indexT3CornerPos++
				}
				children = append(children, *child)
			}
		}
		parents = children
	}
}

func createTable3(tables *tables) {
	fmt.Println("Generating pruning table for phase 3")

	cube := &Cube{}
	cube.initialize()
	
	parents := []Cube{
		*cube,
	}

	for depth := 1; depth < 16; depth++ {
		var children []Cube
		fmt.Println("table3: depth =", depth, "done")
		for _, parent := range parents {
			for _, move := range GetGroupMoves(3, parent) {
				child := copyAndMove(&parent, move)
				indexCornerPos := tables.T3cornerPosIndex[cPosToInt(child.cornerPos)]
				indexEdgePos := edgePosConverter(child.edgePos)
				if tables.Table3[indexCornerPos][indexEdgePos[0]][indexEdgePos[1]][indexEdgePos[2]] == 0 {
					tables.Table3[indexCornerPos][indexEdgePos[0]][indexEdgePos[1]][indexEdgePos[2]] = uint8(depth)
					children = append(children, *child)
				}
			}
		}
		parents = children
	}
}


func setTable3(tables *tables) {
	setT3CornerPosIndex(tables)

	if _, err := os.Stat("pruningTables"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("pruningTables", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat("pruningTables/Table3"); os.IsNotExist(err) {
		start := time.Now()
		createTable3(tables)
		fmt.Println("Took ", time.Since(start))

		file, err := os.Create("pruningTables/Table3")
		if err != nil {
			fmt.Println("Error: file creation: Table3")
			os.Exit(1)
		}
		defer file.Close()

		fmt.Println("writing to file")
		for i := 0; i < 96; i++ {
			for j := 0; j < 24; j++ {
				for k := 0; k < 24; k++ {
					for l := 0; l < 24; l++ {
						_, err := file.WriteString(fmt.Sprintf("%x", tables.Table3[i][j][k][l]))
						if err != nil {
							fmt.Println("Error: file write: Table3")
							os.Exit(1)
						}
					}
				}
			}
		}
	} else {
		content, err := os.ReadFile("pruningTables/Table3")
		if err != nil {
			fmt.Println("failed to read pruning table file")
			os.Exit(1)
		}

		iCPos := 0
		iEPos1 := 0
		iEPos2 := 0
		iEPos3 := 0
		for _, depth := range content {
			tables.Table3[iCPos][iEPos1][iEPos2][iEPos3] = hexToInt(depth)
			iEPos3++
			if iEPos3 >= 24 {
				iEPos3 = 0
				iEPos2++
				if iEPos2 >= 24 {
					iEPos2 = 0
					iEPos1++
					if iEPos1 >= 24 {
						iEPos1 = 0
						iCPos++
					}
				}
			}
		}
	}
}
