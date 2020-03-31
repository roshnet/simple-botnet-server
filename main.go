package main

import "fmt"

func main() {
	bot := &botnet{}
	bot.init()
	bot.displayConfig()
	bot.attack()
	fmt.Println("FINISH")
}
