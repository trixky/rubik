package models

import (
)

type tables struct {
	Table0 [2048]uint8
	Table1 [495][2187]uint8
	T1EdgePosIndex [2048]uint16
	Table2 [40320][70]uint8
	T2EdgePosIndex [255]uint8
	Table3 [96][24][24][24]uint8
	T3cornerPosIndex [40320]uint8
}

func (tables *tables) setTables() {
	setTable0(tables)
	setTable1(tables)
	setTable2(tables)
	setTable3(tables)
}
