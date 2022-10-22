package models

import (
	"fmt"
	"time"
	"sync"
	"strings"
	"context"
)

var SolvedGroup string

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
		indexEdgePos := tables.T1EdgePosIndex[edgePosToInt(node.edgePos)]
		indexCornerOri := cornerOriToInt(node.cornerOri)
		return int(tables.Table1[indexEdgePos][indexCornerOri])
	} else if group == 2 {
		return int(tables.Table2[cPosToInt(node.cornerPos)][tables.T2EdgePosIndex[edgePosToIntT2(node.edgePos)]])
	} else {
		ePos := edgePosConverter(node.edgePos)
		return int(tables.Table3[tables.T3cornerPosIndex[cPosToInt(node.cornerPos)]][ePos[0]][ePos[1]][ePos[2]])
	}
}

func search(ctx context.Context, path []Cube, g int, bound int, group int, tables *tables) int {
	if ctx.Err()!=nil {
		return -1
	}
	node := path[len(path) - 1]
	h := heuristic(&node, group, tables)
	f := g + h

	if f > bound {

		return f
	}
	if h == 0 {
		if SolvedGroup == "" {
			for _, cube := range path {
				SolvedGroup += cube.move + " "
			}
		}
		return -1
	}

	min := 255
	moves := GetGroupMoves(group, node)
	for _, move := range moves {
		succ := copyAndMove(&node, move)
		if isInPath(succ, path) == false {
			path = append(path, *succ)
			cost := search(ctx, path, g + heuristic(succ, group, tables), bound, group, tables)
			if cost == -1 {
				return -1
			}
			if cost < min {
				min = cost
			}
			path = path[:len(path) - 1]
		}
	}
	return min
}

func idaStar(node *Cube, group int, tables *tables) {
	h := heuristic(node, group, tables)
	if h == 0 {
		return
	}

	var wg sync.WaitGroup
	ctx, cancel:= context.WithCancel(context.Background())

	for _, move := range GetGroupMoves(group, *node) {
		// succ2 := copyAndMove(node, move)
		wg.Add(1)
		// fmt.Println("bound", heuristic(succ2, group, tables))
		go func(move string) {
			defer wg.Done()
			defer cancel() // cancel context once this goroutine ends

			succ := copyAndMove(node, move)
			bound := heuristic(succ, group, tables)
			path := []Cube{*succ}
			for {
				cost := search(ctx, path, 0, bound, group, tables)
				// fmt.Println("bound", bound)
				if cost == -1 {
					// fmt.Println("bound found", bound)
					return
				}
				bound = cost
			}
		 }(move)
	}
	wg.Wait()
	return
}

func SolveSequence(server bool, verbose bool, sequence []string) []string {
	tables := &tables{}
	tables.setTables()

	node := MakeCubeFromSequence(sequence)

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
		SolvedGroup = ""
		
		start := time.Now()
		idaStar(node, group, tables)
		groupSequence := strings.Fields(SolvedGroup)
		SolvedGroup = ""

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
		node.Print(-1)
		fmt.Println("og sequence", sequence)
		fmt.Println("solve sequence", totalSequenceSolve)
	}
	return totalSequenceSolve
}
