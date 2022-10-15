package models

import (
)

/*
** isSolved checks if the cube given is solved
** argument:
** @cube *Cube: the cube
** returns a boolean: true if solved, false otherwise
**
** in this file because explicitely checks groups' rules in the given cube
*/
func isSolved(cube *Cube) bool { // WIP
	for i := range cube.cornerPos {
		if cube.cornerPos[i] != uint8(i) {
			return false
		}
	}
	for i := range cube.cornerOri {
		if cube.cornerOri[i] != 0 {
			return false
		}
	}
	for i := range cube.edgePos {
		if cube.edgePos[i] != uint8(i) {
			return false
		}
	}
	for i := range cube.edgeOri {
		if cube.edgeOri[i] != 0 {
			return false
		}
	}
	return true
}


/*
** whichGroup returns group number of given cube
** argument:
** @cube *Cube: the cube
** returns the cube's group number
*/
func whichGroup(cube *Cube) int {
	/*
	** phase 1:
	** orient edges
	*/
	for i := range cube.edgeOri { // all edges
		if cube.edgeOri[i] != byte(0) { // are oriented
			return 0
		}
	}

	/*
	** phase 2:
	** orient corners
	** set middle edges in LR faces
	*/
	for i := range cube.cornerOri { // all corners
		if cube.cornerOri[i] != byte(0) { // are oriented
			return 1
		}
	}
	for i := 8; i < 12; i++ { // edges 8 to 11 are middle one (fr, fl, br, bl)
		if cube.edgePos[i] < uint8(8) { // check that they are in the middle
			return 1
		}
	}

	/*
	** phase 3:
	** set corners in correct tetrads
	** set other edges in correct faces
	** has even parity
	*/
	for i := 0; i < 4; i++ { // corner correct tetrad
		if cube.cornerPos[i] > uint8(3) {
			return 2
		}
	}
	for i := 4; i < 8; i++ { // corner correct tetrad
		if cube.cornerPos[i] < uint8(4) {
			return 2
		}
	}
	for i := 0; i < 4; i++ { // edge correct face
		if cube.edgePos[i] > uint8(3) {
			return 2
		}
	}
	for i := 4; i < 8; i++ { // edge correct face
		if cube.edgePos[i] < uint8(4) || cube.edgePos[i] > uint8(7) {
			return 2
		}
	}
	var parity int
	for i := range cube.cornerPos {
		if cube.cornerPos[i] != uint8(i) {
			parity++
		}
	}
	if parity % 2 == 1 { // if parity is uneven, still phase 2
		return 2
	}

	/*
	** phase 4
	** solve the rubik
	*/
	if isSolved(cube) == false {
		return 3
	}

	return 4
}


/*
** getGroupMoves returns coset moves to get to the next group
** argument:
** @group uint8: the current group number
** @c Cube: cube from which we get two last moves not to do them
** returns a array of string: the coset moves to advance to next group
*/
func getGroupMoves(group int, c Cube) []string { // not sure if need all moves or only elementary
	moves := []string{}
	if group == 0 {
		moves = []string{
			"U",
			"U'",
			"U2",
			"D",
			"D'",
			"D2",
			"L",
			"L'",
			"L2",
			"R",
			"R'",
			"R2",
			"F",
			"F'",
			"F2",
			"B",
			"B'",
			"B2",
		}
	} else if group == 1 {
		moves = []string{
			"U",
			"U'",
			"U2",
			"D",
			"D'",
			"D2",
			"L",
			"L'",
			"L2",
			"R",
			"R'",
			"R2",
			"F2",
			"B2",
		}
	} else if group == 2 {
		moves = []string{
			"U",
			"U'",
			"U2",
			"D",
			"D'",
			"D2",
			"L2",
			"R2",
			"F2",
			"B2",
		}
	} else if group == 3 {
		moves = []string{
			"U2",
			"D2",
			"L2",
			"R2",
			"F2",
			"B2",
		}
	}
	if c.move != "" {

		var fewer_moves []string
		for _, move := range moves {
			if c.move[0] != move[0] {
				fewer_moves = append(fewer_moves, move)
			}
		}
		moves = fewer_moves
	}
	return moves
}
