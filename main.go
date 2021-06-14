package main

import (
	"os"
	"os/signal"
)

func main() {
	//	flags handling block

	gameboy := NewGameboy()

	go gameboy.Launch()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	signal.Notify(exit, os.Kill)

	if <-exit != nil {
		exitGracefully()
	}
}
