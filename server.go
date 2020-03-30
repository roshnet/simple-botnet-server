package main

import (
	"log"
	"net/http"
)

type botnet struct {
	targetHost    string
	serverPort    string
	weakEndpoints []string
	maxRoutines   int32
	httpVerb      string
}


// A mandatory method to call after declaring a `botnet`.
// Initialize the botnet with default values.
func (b botnet) init() {
	// The target server (along with port if not 80 or 443)
	b.targetHost = "http://localhost:8000"

	// Set the port where this script is to be mounted (string)
	b.serverPort = "8080"

	// A byte-slice of most vulnerable endpoints
	b.weakEndpoints = []string{
		"/",
	}

	// The default HTTP method to use
	b.httpVerb = "GET"
	// TBD: Use a map of endpoints-methods.

	// Set the maximum value of goroutines to execute.
	// More threads => more load, but higher CPU use.
	// But, definitely worth increasing if byte amplification is followed.
	b.maxThreads = 10000
}


// `attack()` receives no arguments, but utilises all neccessary members of
// the `botnet` struct
func (b botnet) attack() {
	if b.httpVerb == "GET" {
		_, err := http.Get(b.targetHost)
		// Checkpoint. Do more stuff later.
	}
}
