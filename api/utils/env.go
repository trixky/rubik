package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	ENV_PORT     = "API_PORT"
	default_port = 3020
)

// ReadPort read the port environment variable
func ReadPort() int {
	const (
		port_max = 65535
		port_min = 1000
	)

	// Read the port environment variable
	env_port := os.Getenv(ENV_PORT)

	if len(env_port) == 0 {
		// If port environment variable is not set
		// Return the default port
		return default_port
	}

	// Convert string port environment variable to integer
	port, err := strconv.Atoi(env_port)

	if err != nil {
		// If conversion fail
		log.Fatal(ENV_PORT + "environment variable need to be a number")
	}

	if port < port_min || port > port_max {
		// If port is too big/small
		log.Fatal(fmt.Errorf("port need to be included between %d and %d (%d)", port_min, port_max, port))
	}

	return port
}
