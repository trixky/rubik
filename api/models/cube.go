package models

import (
	"fmt"
	"strings"
)

const (
	ulb = iota // 0
	dlf = iota // 1
	drb = iota // 2
	urf = iota // 3
	ufl = iota // 4
	dbl = iota // 5
	dfr = iota // 6
	ubr = iota // 7
)

const (
	ul = iota // 0
	dl = iota
	dr = iota
	ur = iota
	bl = iota // 4
	fl = iota
	fr = iota
	br = iota
	uf = iota // 8
	df = iota
	db = iota
	ub = iota // 11
)

type Cube struct {
	move     string
	cornerPos [8]uint8
	cornerOri [8]byte
	edgePos   [12]uint8
	edgeOri   [12]byte
}

func (c *Cube) equal(other Cube) bool {
	if c.cornerPos != other.cornerPos {
		return false
	}
	if c.cornerOri != other.cornerOri {
		return false
	}
	if c.edgePos != other.edgePos {
		return false
	}
	if c.edgeOri != other.edgeOri {
		return false
	}
	return true
}


func (c *Cube) initialize() {
	c.move = ""
	c.cornerPos = [8]uint8{ulb, dlf, drb, urf, ufl, dbl, dfr, ubr}
	c.cornerOri = [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	c.edgePos = [12]uint8{ul, dl, dr, ur, bl, fl, fr, br, uf, df, db, ub}
	c.edgeOri = [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}


func (c *Cube) rotateSequence(words []string) {
	for _, s := range words {
		c.rotate(s)
	}
}

func (c *Cube) rotate(move string) {
	amount := 1
	if len(move) == 2 {
		if move[1] == '\'' {
			amount = 3
		} else if move[1] == '2' {
			amount = 2
		}
	}
	switch move[0] {
	case 'U':
		c.rotateUp(amount)
	case 'D':
		c.rotateDown(amount)
	case 'L':
		c.rotateLeft(amount)
	case 'R':
		c.rotateRight(amount)
	case 'F':
		c.rotateFront(amount)
	case 'B':
		c.rotateBack(amount)
	}
	c.move = move
}

const Green		= "\x1B[32m"
const Yellow	= "\x1B[33m"
const Reset		= "\x1B[0m"

func (cube *Cube) Print(group int) {
	fmt.Printf("|%s|\n", strings.Repeat("-", 71))
	fmt.Printf("|%25s      |%21s                  |\n", "Corner", "Edge")
	fmt.Printf("|       Index | 0 1 2 3 4 5 6 7 |  0  1  2  3  4  5  6  7  8  9 10 11   |\n")
	fmt.Printf("|%s|\n", strings.Repeat("-", 71))

	fmt.Printf("| Orientation | ")
	if group == 1 {
		fmt.Printf("%s", Yellow)
	} else if group > 1 {
		fmt.Printf("%s", Green)
	}
	for _, cOri := range cube.cornerOri {
		fmt.Printf("%d ", cOri)
	}
	fmt.Printf("%s| ", Reset)


	if group == 0 {
		fmt.Printf("%s", Yellow)
	} else if group > 0 {
		fmt.Printf("%s", Green)
	}
	for _, eOri := range cube.edgeOri {
		fmt.Printf("%2d ", eOri)
	}
	fmt.Printf("%s  |\n", Reset)


	fmt.Printf("|    Position | ")
	if group == 2 {
		fmt.Printf("%s", Yellow)
	} else if group > 2 {
		fmt.Printf("%s", Green)
	}
	for _, cPos := range cube.cornerPos {
		fmt.Printf("%d ", cPos)
	}
	fmt.Printf("%s| ", Reset)

	for i, ePos := range cube.edgePos {
		if group == 1 {
			if i == 8 {
				fmt.Printf("%s", Yellow)
			}
		} else if group == 2 {
			if i == 0 {
				fmt.Printf("%s", Yellow)
			}
			if i == 8 {
				fmt.Printf("%s", Reset)
			}
		} else if group == 3 {
			fmt.Printf("%s", Green)
		}
		fmt.Printf("%2d ", ePos)
	}
	fmt.Printf("%s  | \n", Reset)
	fmt.Printf("|%s|\n", strings.Repeat("-", 71))
}

func (c *Cube) rotateUp(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[urf]
		c.cornerPos[urf] = c.cornerPos[ubr] 
		c.cornerPos[ubr] = c.cornerPos[ulb] 
		c.cornerPos[ulb] = c.cornerPos[ufl] 
		c.cornerPos[ufl] = tmpPos			

		tmpOri := c.cornerOri[urf]
		c.cornerOri[urf] = (2 + c.cornerOri[ubr]) % 3
		c.cornerOri[ubr] = (1 + c.cornerOri[ulb]) % 3
		c.cornerOri[ulb] = (2 + c.cornerOri[ufl]) % 3
		c.cornerOri[ufl] = (1 + tmpOri) % 3

		tmpPos = c.edgePos[uf]
		c.edgePos[uf] = c.edgePos[ur]  
		c.edgePos[ur] = c.edgePos[ub] 
		c.edgePos[ub] = c.edgePos[ul] 
		c.edgePos[ul] = tmpPos		  

		tmpOri = c.edgeOri[uf]
		c.edgeOri[uf] = 1 - c.edgeOri[ur]
		c.edgeOri[ur] = 1 - c.edgeOri[ub]
		c.edgeOri[ub] = 1 - c.edgeOri[ul]
		c.edgeOri[ul] = 1 - tmpOri
	}
}

func (c *Cube) rotateDown(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[dfr]
		c.cornerPos[dfr] = c.cornerPos[dlf] 
		c.cornerPos[dlf] = c.cornerPos[dbl] 
		c.cornerPos[dbl] = c.cornerPos[drb] 
		c.cornerPos[drb] = tmpPos			

		tmpOri := c.cornerOri[dfr]
		c.cornerOri[dfr] = (2 + c.cornerOri[dlf]) % 3
		c.cornerOri[dlf] = (1 + c.cornerOri[dbl]) % 3
		c.cornerOri[dbl] = (2 + c.cornerOri[drb]) % 3
		c.cornerOri[drb] = (1 + tmpOri) % 3

		tmpPos = c.edgePos[df]
		c.edgePos[df] = c.edgePos[dl]	
		c.edgePos[dl] = c.edgePos[db]	
		c.edgePos[db] = c.edgePos[dr]	
		c.edgePos[dr] = tmpPos			 

		tmpOri = c.edgeOri[df]
		c.edgeOri[df] = 1 - c.edgeOri[dl]
		c.edgeOri[dl] = 1 - c.edgeOri[db]
		c.edgeOri[db] = 1 - c.edgeOri[dr]
		c.edgeOri[dr] = 1 - tmpOri
	}
}

func (c *Cube) rotateFront(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[urf]
		c.cornerPos[urf] = c.cornerPos[ufl] 
		c.cornerPos[ufl] = c.cornerPos[dlf] 
		c.cornerPos[dlf] = c.cornerPos[dfr] 
		c.cornerPos[dfr] = tmpPos			

		tmpOri := c.cornerOri[urf]
		c.cornerOri[urf] = (1 + c.cornerOri[ufl]) % 3
		c.cornerOri[ufl] = (2 + c.cornerOri[dlf]) % 3
		c.cornerOri[dlf] = (1 + c.cornerOri[dfr]) % 3
		c.cornerOri[dfr] = (2 + tmpOri) % 3

		tmpPos = c.edgePos[uf]
		c.edgePos[uf] = c.edgePos[fl]
		c.edgePos[fl] = c.edgePos[df]
		c.edgePos[df] = c.edgePos[fr] 
		c.edgePos[fr] = tmpPos		  

		tmpOri = c.edgeOri[uf]
		c.edgeOri[uf] = c.edgeOri[fl]
		c.edgeOri[fl] = c.edgeOri[df]
		c.edgeOri[df] = c.edgeOri[fr]
		c.edgeOri[fr] = tmpOri
	}
}

func (c *Cube) rotateBack(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[ubr]
		c.cornerPos[ubr] = c.cornerPos[drb] 
		c.cornerPos[drb] = c.cornerPos[dbl] 
		c.cornerPos[dbl] = c.cornerPos[ulb] 
		c.cornerPos[ulb] = tmpPos			

		tmpOri := c.cornerOri[ubr]
		c.cornerOri[ubr] = (1 + c.cornerOri[drb]) % 3
		c.cornerOri[drb] = (2 + c.cornerOri[dbl]) % 3
		c.cornerOri[dbl] = (1 + c.cornerOri[ulb]) % 3
		c.cornerOri[ulb] = (2 + tmpOri) % 3
		
		
		tmpPos = c.edgePos[ub]
		c.edgePos[ub] = c.edgePos[br] 
		c.edgePos[br] = c.edgePos[db] 
		c.edgePos[db] = c.edgePos[bl]
		c.edgePos[bl] = tmpPos

		tmpOri = c.edgeOri[ub]
		c.edgeOri[ub] = c.edgeOri[br]
		c.edgeOri[br] = c.edgeOri[db]
		c.edgeOri[db] = c.edgeOri[bl]
		c.edgeOri[bl] = tmpOri
	}
}

func (c *Cube) rotateLeft(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[ulb]
		c.cornerPos[ulb] = c.cornerPos[dbl] 
		c.cornerPos[dbl] = c.cornerPos[dlf] 
		c.cornerPos[dlf] = c.cornerPos[ufl] 
		c.cornerPos[ufl] = tmpPos			

		tmpOri := c.cornerOri[ulb]
		c.cornerOri[ulb] = c.cornerOri[dbl]
		c.cornerOri[dbl] = c.cornerOri[dlf]
		c.cornerOri[dlf] = c.cornerOri[ufl]
		c.cornerOri[ufl] = tmpOri
		
		
		tmpPos = c.edgePos[ul]
		c.edgePos[ul] = c.edgePos[bl]
		c.edgePos[bl] = c.edgePos[dl]
		c.edgePos[dl] = c.edgePos[fl]
		c.edgePos[fl] = tmpPos

		tmpOri = c.edgeOri[ul]
		c.edgeOri[ul] = c.edgeOri[bl]
		c.edgeOri[bl] = c.edgeOri[dl]
		c.edgeOri[dl] = c.edgeOri[fl]
		c.edgeOri[fl] = tmpOri
	}
}

func (c *Cube) rotateRight(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[urf]
		c.cornerPos[urf] = c.cornerPos[dfr] 
		c.cornerPos[dfr] = c.cornerPos[drb] 
		c.cornerPos[drb] = c.cornerPos[ubr] 
		c.cornerPos[ubr] = tmpPos			

		tmpOri := c.cornerOri[urf]
		c.cornerOri[urf] = c.cornerOri[dfr]
		c.cornerOri[dfr] = c.cornerOri[drb]
		c.cornerOri[drb] = c.cornerOri[ubr]
		c.cornerOri[ubr] = tmpOri
		
		
		tmpPos = c.edgePos[ur]
		c.edgePos[ur] = c.edgePos[fr] 
		c.edgePos[fr] = c.edgePos[dr] 
		c.edgePos[dr] = c.edgePos[br] 
		c.edgePos[br] = tmpPos		  

		tmpOri = c.edgeOri[ur]
		c.edgeOri[ur] = c.edgeOri[fr]
		c.edgeOri[fr] = c.edgeOri[dr]
		c.edgeOri[dr] = c.edgeOri[br]
		c.edgeOri[br] = tmpOri
	}
}

func makeCubeFromSequence(sequence []string) *Cube {
	cube := &Cube{}
	cube.initialize()
	cube.rotateSequence(sequence)
	return cube
}

func copyAndMove(cube *Cube, move string) *Cube {
	copy := &Cube{
		move: cube.move,
		cornerPos: cube.cornerPos,
		cornerOri: cube.cornerOri,
		edgePos: cube.edgePos,
		edgeOri: cube.edgeOri,
	}
	copy.rotate(move)
	return copy
}
