package input

import (
	"github.com/trixky/rubik/models"
	"github.com/trixky/rubik/parser"
)

func Get() (set models.Set, api_mode bool) {
	instructions := readArg()

	if isApiMode(instructions) {
		api_mode = true
	} else {
		set = parser.ProcessString(instructions)
	}

	return
}
