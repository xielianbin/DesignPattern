package main

import (
	"factory/abstract"
	"factory/specific"
	"fmt"
)

func main() {
	ak47, _ := specific.GetGun("ak47")
	musket, _ := specific.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g abstract.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
