package main

import "net/http"

func strike(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- resp.Status + " : " + url
	}
	c <- resp.Status + " : " + url
}
