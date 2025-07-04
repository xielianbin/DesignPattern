package main

import (
	"bridge/internal"
	"fmt"
)

func main() {

	hpPrinter := &internal.Hp{}
	epsonPrinter := &internal.Epson{}

	macComputer := &internal.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &internal.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()

}
