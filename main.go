package main

func main() {
	bot := botnet{}
	bot.init()
	log.Fatal(http.ListenAndServe(":"+bot.serverPort, nil))
}
