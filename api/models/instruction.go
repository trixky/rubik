package models

import "fmt"

// F R U B L D
const (
	MOVE_F byte = 'F'
	MOVE_R byte = 'R'
	MOVE_U byte = 'U'
	MOVE_B byte = 'B'
	MOVE_L byte = 'L'
	MOVE_D byte = 'D'
)

const (
	MODIFIER_NOTHING byte = 'n'
	MODIFIER_REVERSE byte = '\''
	MODIFIER_DOUBLE  byte = '2'
)

type Instruction struct {
	Move     byte
	Modifier byte
}

func (i *Instruction) ToString() string {
	if i.Modifier != MODIFIER_NOTHING {
		return string(i.Move) + string(i.Modifier)
	}

	return string(i.Move)
}

func (i *Instruction) Print() {
	fmt.Print(i.ToString())
}
