package models

import (
	"fmt"
	"time"
	"strings"
)

func isInPath(succ *Cube, path []Cube) bool {
	for _, cube := range path {
		if succ.equal(cube) == true {
			return true
		}
	}
	return false
}

func heuristic(node *Cube, group int, tables *tables) int {
	if group == 0 {
		return int(tables.Table0[edgeOriToInt(node.edgeOri)])
	} else if group == 1 {
		return int(tables.Table1[tables.T1EdgePosIndex[edgePosToInt(node.edgePos)]][cornerOriToInt(node.cornerOri)])
	} else if group == 2 {
		return int(tables.Table2[cPosToInt(node.cornerPos)][tables.T2EdgePosIndex[edgePosToIntT2(node.edgePos)]])
	} else {
		ePos := edgePosConverter(node.edgePos)
		return int(tables.Table3[tables.T3cornerPosIndex[cPosToInt(node.cornerPos)]][ePos[0]][ePos[1]][ePos[2]])
	}
}

func search(path []Cube, g int, bound int, group int, tables *tables) (int, string) {
	node := path[len(path) - 1]
	h := heuristic(&node, group, tables)
	f := g + h

	if f > bound {
		return f, ""
	}
	if h == 0 {
		var solvedGroup string
		for _, cube := range path {
			solvedGroup += cube.move + " "
		}
		return 255, solvedGroup
	}

	min := 255
	moves := GetGroupMoves(group, node)
	for _, move := range moves {
		succ := copyAndMove(&node, move)
		if isInPath(succ, path) == false {
			path = append(path, *succ)
			cost, solvedGroup := search(path, g + heuristic(succ, group, tables), bound, group, tables)
			if cost == 255 {
				return 255, solvedGroup
			}
			if cost < min {
				min = cost
			}
			path = path[:len(path) - 1]
		}
	}
	return min, ""
}

func idaStar(node *Cube, group int, tables *tables) []string {
	path := []Cube{*node}
	bound := heuristic(node, group, tables)
	for {
		cost, solvedGroup := search(path, 0, bound, group, tables)
		if cost == 255 {
			return strings.Fields(solvedGroup)
		}
		bound = cost
	}
}

func SolveSequence(server bool, verbose bool, sequence []string) []string {
	tables := &tables{}
	tables.setTables()

	node := makeCubeFromSequence(sequence)

	if verbose {
		fmt.Println("This is the cube after doing your sequence :")
		node.Print(-1)
		fmt.Println("It it part of group", whichGroup(node))
		fmt.Println("Time to solve :)")
		fmt.Println()
	}

	var totalSequenceSolve []string
	for group := 0; group < 4; group++ {
		node.move = ""
		
		start := time.Now()
		groupSequence := idaStar(node, group, tables)
		
		node.rotateSequence(groupSequence)
		
		totalSequenceSolve = append(totalSequenceSolve, groupSequence...)
		
		if server && group != 3 {
			totalSequenceSolve = append(totalSequenceSolve, ",")
		}

		if verbose {
			fmt.Println("Going from group", group, "to", group + 1)
			node.Print(group)
			fmt.Println("solve sequence : ", groupSequence)
			fmt.Println("length : ", len(groupSequence))
			fmt.Println("time : ", time.Since(start))
			fmt.Println()
		}
	}
	if isSolved(node) == false {
		fmt.Println("NOT SOLVED ???")
	} 
	return totalSequenceSolve
}
