package models

import (
	"fmt"
	"strings"
)

type Set struct {
	Instructions []Instruction
}

func (s *Set) ToString() string {
	instructions := make([]string, len(s.Instructions))

	for index, instruction := range s.Instructions {
		instructions[index] = instruction.ToString()
	}

	return strings.Join(instructions, " ")
}

func (s *Set) Print() {
	fmt.Println(s.ToString())
}
