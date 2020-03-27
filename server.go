package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		panic("Forgot something maybe?")
	}
	TLCallbacks, typeError := strconv.Atoi(os.Args[2])
	if typeError != nil {
		panic("Cannot understand supplied value of top level callbacks.")
	}
	url := os.Args[1] // Add test cases to validate URL structure
	c := make(chan string)

	fmt.Println("[STANDBY] T minus 3 seconds...")
	time.Sleep(time.Second * 3)

	// Single way pinging for now...
	for i := 0; i < TLCallbacks; i++ {
		go strike(url, c)
	}

	for i := 0; i < TLCallbacks; i++ {
		<-c
	}
}
