package main

import (
	"os"
	"fmt"
	"strings"
)


func valid_shuffle(words []string) bool {
	var valid_chars = "UDFRBL"
	var valid_sec_chars = "'2"

	for _, s := range words {
		if (len(s) > 2) {
			return false
		}
		if (strings.Contains(valid_chars, string(s[0])) == false) {
			return false
		}
		if (len(s) == 2) {
			if (strings.Contains(valid_sec_chars, string(s[1])) == false) {
				return false
			}
		}
	}
	return true
}

func create_patterns() {
	os.Mkdir("patterns", os.ModePerm)

}

func main() {
    args := os.Args[1:]
	var words []string

	if (len(args) == 1) {
		words = strings.Fields(args[0])
	}
	if (valid_shuffle(words) == false) {
		fmt.Println("Argument is wrong")
		fmt.Println("  Valid characters are 'UDFRBL' and can be followed by either ' (reverse) and 2 (twice).")
		fmt.Println("  Each move seperated by whitespace(s).")
		fmt.Println("  Good Example : \"F R U2 B' R' L D2\"")
	}
	if _, err := os.Stat("/patterns"); os.IsNotExist(err) {
		create_patterns()
	}

	c := NewCube()
	c.startPosition(words)
	// s := NewSolver(c, )
	c.print()
	// print words
	// for _, s := range words {
	// 	fmt.Println(s)
	// }
}
