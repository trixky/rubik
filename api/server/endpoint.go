package server

import (
	"io/ioutil"
	"fmt"
	"strings"
	"net/http"
	"log"

	"github.com/trixky/rubik/models"
)

func resolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}
	sequence := strings.Fields(string(body))
	result := models.SolveSequence(false, sequence)

	var str string
	for _, s := range result {
		str += s
	}
	fmt.Fprintf(w, str)
}
