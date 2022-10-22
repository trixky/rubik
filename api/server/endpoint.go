package server

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"io/ioutil"

	"github.com/trixky/rubik/models"
)

func resolve(w http.ResponseWriter, r *http.Request) {
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
	start := time.Now()
	result := models.SolveSequence(true, false, sequence)
	duration := time.Since(start)
	strToFront := fmt.Sprint(duration.Milliseconds()) + " "
	for _, s := range result {
		strToFront += s + " "
	}
	
	fmt.Fprintf(w, strToFront)
}
