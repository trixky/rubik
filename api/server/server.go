package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/trixky/rubik/utils"
)

func Start() {
	port := utils.ReadPort()

	mux := http.NewServeMux()

	// register the resolve handler to the mux
	mux.HandleFunc("/resolve", resolve)

	// Wrapp mux in the CORS middleware
	wrapped_mux := AllowCORS(mux)
	addr := ":" + strconv.Itoa(port)

	fmt.Println("start api on port :" + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(addr, wrapped_mux))
}
