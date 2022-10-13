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

func heuristic(node Cube, group int, tables *tables) int {
	if group == 0 {
		return int(tables.Table0[edgeOriToInt(node.edgeOri)])
	} else if group == 1 {
		return int(tables.Table1[cornerOriToInt(node.cornerOri)][edgePosToInt(node.edgeOri)])
	}
	return int(tables.Table1[cornerOriToInt(node.cornerOri)][edgePosToInt(node.edgeOri)])

	// } else if group == 2 {
		
	// } else if group == 3 {
		
	// }
}

func search(path []Cube, g int, bound int, group int, tables *tables) (int, string) {
	node := path[len(path) - 1]
	h := heuristic(node, group, tables)
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
	moves := getGroupMoves(group, node)
	for _, move := range moves {
		fmt.Println(move)
		succ := copyAndMove(&node, move)
		if isInPath(succ, path) == false {
			path = append(path, *succ)
			cost, solvedGroup := search(path, g + heuristic(*succ, group, tables), bound, group, tables)
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
	bound := heuristic(*node, group, tables)
	for {
		cost, solvedGroup := search(path, 0, bound, group, tables)
		if cost == 255 {
			return strings.Fields(solvedGroup)
		}
		bound = cost
	}
}

func SolveSequence(verbose bool, sequence []string) []string {
	tables := &tables{}
	tables.setTables()
	node := makeCubeFromSequence(sequence)
	fmt.Println("group of beg node : ", whichGroup(node))
	var totalSequenceSolve []string
	for group := whichGroup(node); group < 2; group++ {
		node.move = ""
		start := time.Now()
		
		groupSequence := idaStar(node, group, tables)
		
		node.rotateSequence(groupSequence)
		
		fmt.Println("\ngroup : ", group)
		fmt.Println("solve sequence : ", groupSequence)
		fmt.Println("time : ", time.Since(start))
		fmt.Println()

		totalSequenceSolve = append(totalSequenceSolve, groupSequence...)
	}
	return totalSequenceSolve
}
