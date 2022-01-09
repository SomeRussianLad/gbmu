package main

import (
	"flag"
	"gbmu/emulation"
)

const (
	GBMU_OUTPUT_FYNE = iota
	GBMU_OUTPUT_TERM
)

var (
	help bool
)

func init() {
	flag.BoolVar(&help, "help", false, "")
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	emulator := emulation.NewEmulator()
	emulator.Run()
}
