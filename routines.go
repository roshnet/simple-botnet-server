package main

import "net/http"

func strike(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- resp.Status + " : " + url
	}
	c <- resp.Status + " : " + url
}

func multiStrike(url string, N uint64, c chan string) {
	for i := N - 1; i >= 0; i-- {
		strike(url, c)
	}
}
