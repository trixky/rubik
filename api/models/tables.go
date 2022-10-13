package models

import (
	"fmt"
)

type tables struct {
	Table0 [2048]uint8
	Table1 [495][2187]uint8
	T1EdgePosIndex [2048]uint16
}

func (tables *tables) setTables() {
	fmt.Println("Getting tables")
	setTable0(tables)
	setTable1(tables)
	// setTable2(tables)
	// setTable3(tables)
}
