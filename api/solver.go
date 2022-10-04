package main

import (
	// "fmt"
	// "strings"
)

type Solver struct {
	initCube Cube
}

func NewSolver(c Cube) Solver {
	s := Solver {c}
	return s
}