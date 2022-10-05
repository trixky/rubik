package server

import (
	"fmt"
	"net/http"

	"github.com/trixky/rubik/algo"
	"github.com/trixky/rubik/models"
)

func resolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result := algo.Resolve(models.Set{})

	fmt.Fprintf(w, result.ToString())
}
