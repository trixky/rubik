package models

import (
	"fmt"
)

const (
	urf = iota
	ubr = iota
	dlf = iota
	dfr = iota
	ulb = iota
	ufl = iota
	drb = iota
	dbl = iota
)

const (
	uf = iota
	ur = iota
	ub = iota
	ul = iota
	df = iota
	dr = iota
	db = iota
	dl = iota
	fr = iota
	br = iota
	bl = iota
	fl = iota
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
	c.cornerPos = [8]uint8{urf, ubr, dlf, dfr, ulb, ufl, drb, dbl}
	c.cornerOri = [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	c.edgePos = [12]uint8{uf, ur, ub, ul, df, dr, db, dl, fr, br, bl, fl}
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


func (c *Cube) rotateUp(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[urf]
		c.cornerPos[urf] = c.cornerPos[ubr]
		c.cornerPos[ubr] = c.cornerPos[ulb]
		c.cornerPos[ulb] = c.cornerPos[ufl]
		c.cornerPos[ufl] = tmpPos

		tmpOri := c.cornerOri[urf]
		c.cornerOri[urf] = c.cornerOri[ubr]
		c.cornerOri[ubr] = c.cornerOri[ulb]
		c.cornerOri[ulb] = c.cornerOri[ufl]
		c.cornerOri[ufl] = tmpOri

		tmpPos = c.edgePos[uf]
		c.edgePos[uf] = c.edgePos[ur]
		c.edgePos[ur] = c.edgePos[ub]
		c.edgePos[ub] = c.edgePos[ul]
		c.edgePos[ul] = tmpPos

		tmpOri = c.edgeOri[uf]
		c.edgeOri[uf] = c.edgeOri[ur]
		c.edgeOri[ur] = c.edgeOri[ub]
		c.edgeOri[ub] = c.edgeOri[ul]
		c.edgeOri[ul] = tmpOri
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
		c.cornerOri[dfr] = c.cornerOri[dlf]
		c.cornerOri[dlf] = c.cornerOri[dbl]
		c.cornerOri[dbl] = c.cornerOri[drb]
		c.cornerOri[drb] = tmpOri

		tmpPos = c.edgePos[df]
		c.edgePos[df] = c.edgePos[dl]
		c.edgePos[dl] = c.edgePos[db]
		c.edgePos[db] = c.edgePos[dr]
		c.edgePos[dr] = tmpPos

		tmpOri = c.edgeOri[df]
		c.edgeOri[df] = c.edgeOri[dl]
		c.edgeOri[dl] = c.edgeOri[db]
		c.edgeOri[db] = c.edgeOri[dr]
		c.edgeOri[dr] = tmpOri
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
		c.cornerOri[ulb] = (1 + c.cornerOri[dbl]) % 3
		c.cornerOri[dbl] = (2 + c.cornerOri[dlf]) % 3
		c.cornerOri[dlf] = (1 + c.cornerOri[ufl]) % 3
		c.cornerOri[ufl] = (2 + tmpOri) % 3

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
		c.cornerOri[urf] = (1 + c.cornerOri[dfr]) % 3
		c.cornerOri[dfr] = (2 + c.cornerOri[drb]) % 3
		c.cornerOri[drb] = (1 + c.cornerOri[ubr]) % 3
		c.cornerOri[ubr] = (2 + tmpOri) % 3

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

func (c *Cube) rotateFront(nbRotations int) {
	for i := 0; i < nbRotations; i++ {
		tmpPos := c.cornerPos[urf]
		c.cornerPos[urf] = c.cornerPos[ufl]
		c.cornerPos[ufl] = c.cornerPos[dlf]
		c.cornerPos[dlf] = c.cornerPos[dfr]
		c.cornerPos[dfr] = tmpPos

		tmpOri := c.cornerOri[urf]
		c.cornerOri[urf] = (2 + c.cornerOri[ufl]) % 3
		c.cornerOri[ufl] = (1 + c.cornerOri[dlf]) % 3
		c.cornerOri[dlf] = (2 + c.cornerOri[dfr]) % 3
		c.cornerOri[dfr] = (1 + tmpOri) % 3

		tmpPos = c.edgePos[uf]
		c.edgePos[uf] = c.edgePos[fl]
		c.edgePos[fl] = c.edgePos[df]
		c.edgePos[df] = c.edgePos[fr]
		c.edgePos[fr] = tmpPos

		tmpOri = c.edgeOri[uf]
		c.edgeOri[uf] = 1 - c.edgeOri[fl]
		c.edgeOri[fl] = 1 - c.edgeOri[df]
		c.edgeOri[df] = 1 - c.edgeOri[fr]
		c.edgeOri[fr] = 1 - tmpOri
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
		c.cornerOri[ubr] = (2 + c.cornerOri[drb]) % 3
		c.cornerOri[drb] = (1 + c.cornerOri[dbl]) % 3
		c.cornerOri[dbl] = (2 + c.cornerOri[ulb]) % 3
		c.cornerOri[ulb] = (1 + tmpOri) % 3

		tmpPos = c.edgePos[ub]
		c.edgePos[ub] = c.edgePos[br]
		c.edgePos[br] = c.edgePos[db]
		c.edgePos[db] = c.edgePos[bl]
		c.edgePos[bl] = tmpPos

		tmpOri = c.edgeOri[ub]
		c.edgeOri[ub] = 1 - c.edgeOri[br]
		c.edgeOri[br] = 1 - c.edgeOri[db]
		c.edgeOri[db] = 1 - c.edgeOri[bl]
		c.edgeOri[bl] = 1 - tmpOri
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

func (c Cube) Print() {
	fmt.Println("print cube")

	fmt.Println("move:", c.move)
	fmt.Println("corners")
	for _, value := range c.cornerPos {
		fmt.Print(value, " ")
	}
	fmt.Println()
	fmt.Println("corners orientation")
	for _, value := range c.cornerOri {
		fmt.Print(value, " ")
	}
	fmt.Println()
	fmt.Println("edges")
	for _, value := range c.edgePos {
		fmt.Print(value, " ")
	}
	fmt.Println()
	fmt.Println("edges orientation")
	for _, value := range c.edgeOri {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")
}
